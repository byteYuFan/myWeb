package command

import (
	"context"
	"crypto/subtle"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"log"
	db "myWeb/DataBase/DB"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/pkg/errno"
	redisMiddleware "myWeb/pkg/redis"
	"myWeb/pkg/ttviper"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type CheckUserService struct {
	ctx context.Context
}
type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}
type EmailConfig struct {
	SmtpServer     string
	SmtpPort       string
	SenderEmail    string
	SenderPassword string
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser 检查用户是否存在
func (s *CheckUserService) CheckUser(req *usersrv.UsernamePasswordLoginRequest) (int64, error) {
	username := req.Username
	user, err := db.QueryUserByName(s.ctx, username)
	if err != nil {
		return 0, err
	}
	if user.Username == "" || user.Flag == false {
		return 0, errno.ErrUserNotFound
	}
	passWordMatch, err := comparePasswordAndHash(req.Password, user.Password)
	if err != nil {
		return 0, err
	}
	if !passWordMatch {
		return 0, errno.ErrPasswordIncorrect
	}
	return int64(user.ID), nil
}
func (s *CheckUserService) CheckUserEmail(req *usersrv.EmailLoginRequest) (int64, string, error) {
	email := req.Email
	user, err := db.QueryUserByEmail(s.ctx, email)
	if err != nil {
		return 0, "", err
	}
	if user.Username == "" || user.Flag == false {
		return 0, "", errno.ErrUserNotFound
	}
	code, err := redisMiddleware.GetEmailCode("email", email)
	if err != nil {
		return 0, "", err
	}
	if code != req.Credential {
		return 0, "", errors.New("验证码错误")
	}
	return int64(user.ID), user.Username, nil
}

// comparePasswordAndHash compares the password and hash of the given password.
func comparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	argon2Params, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the input password using the same parameters.
	inputHash := argon2.IDKey([]byte(password), salt, argon2Params.Iterations, argon2Params.Memory, argon2Params.Parallelism, argon2Params.KeyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, inputHash) == 1 {
		return true, nil
	}
	return false, nil
}

// decodeHash decode the hash of the password from the database.
//
// returns an error if the password is not valid.
func decodeHash(encodedHash string) (argon2Params *Argon2Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errno.ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errno.ErrIncompatibleVersion
	}

	argon2Params = &Argon2Params{}
	if _, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &argon2Params.Memory, &argon2Params.Iterations, &argon2Params.Parallelism); err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	argon2Params.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	argon2Params.KeyLength = uint32(len(hash))

	return argon2Params, salt, hash, nil
}

type SendEmailService struct {
	ctx context.Context
}

func NewSendEmailService(ctx context.Context) *SendEmailService {
	return &SendEmailService{
		ctx: ctx,
	}
}
func (s *SendEmailService) Send(to string) (string, error) {
	var emailCfg EmailConfig
	cfg := ttviper.ConfigInit("userConfig.yml")
	emailCfg.SmtpServer = cfg.Viper.GetString("LoginServer.Email.SMTP_SERVER")
	emailCfg.SmtpPort = cfg.Viper.GetString("LoginServer.Email.SMTP_PORT")
	emailCfg.SenderEmail = cfg.Viper.GetString("LoginServer.Email.SENDER_EMAIL")
	emailCfg.SenderPassword = cfg.Viper.GetString("LoginServer.Email.SENDER_PASSWORD")
	fmt.Println("--------", emailCfg)
	return sendCode(to, emailCfg)
}
func sendCode(to string, config EmailConfig) (string, error) {
	// 生成 6 位数的验证码
	code := generateCode()
	// 发送邮件
	subject := "您的验证码"
	body := "您的验证码为：" + code + "，有效时间为 5 分钟，请尽快完成操作。"
	err := sendEmail(config, config.SenderEmail, to, subject, body)
	if err != nil {
		log.Println("发送邮件失败", err)
		return "", err
	} else {
		//将验证码存入redis缓存，并设置过期时间
		err := redisMiddleware.StoreCode("email", to, code, time.Minute*5)
		if err != nil {
			return "", err
		}
		log.Println("发送成功")
		return code, nil
	}
}

// 生成 6 位数的验证码
func generateCode() string {
	return strconv.Itoa(int(time.Now().UnixNano() / 1000000 % 1000000))
}

// 发送邮件
func sendEmail(config EmailConfig, from, to, subject, body string) error {

	// 邮件头
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"

	// 邮件体
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 连接 SMTP 服务器
	println(config.SenderPassword)
	auth := smtp.PlainAuth("", from, config.SenderPassword, config.SmtpServer)
	tlsConfig := &tls.Config{
		ServerName: config.SmtpServer,
	}
	addr := fmt.Sprintf("%s:%s", config.SmtpServer, config.SmtpPort)
	fmt.Println(addr)
	client, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("连接 SMTP 服务器失败：%w", err)
	}
	defer client.Quit()

	// 开始认证
	if err := client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("启用 TLS 加密失败：%w", err)
	}
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("认证失败：%w", err)
	}

	// 设置发件人、收件人、邮件体
	if err := client.Mail(from); err != nil {
		return fmt.Errorf("设置发件人失败：%w", err)
	}
	if err := client.Rcpt(to); err != nil {
		return fmt.Errorf("设置收件人失败：%w", err)
	}
	wc, err := client.Data()
	if err != nil {

		return fmt.Errorf("获取写入流失败：%w", err)
	}
	defer wc.Close()

	// 写入邮件内容
	if _, err := fmt.Fprintf(wc, message); err != nil {
		return fmt.Errorf("写入邮件内容失败：%w", err)
	}

	return nil
}
