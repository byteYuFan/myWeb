package command

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"log"
	db "myWeb/DataBase/DB"
	"myWeb/kitex_gen/user/usersrv"
	"myWeb/pkg/errno"
	redisMiddleware "myWeb/pkg/redis"
)

type CreateUserService struct {
	ctx context.Context
}
type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// NewCreateUserService 创建一个实例
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (c *CreateUserService) CreateUser(req *usersrv.RegisterRequest, argon2Params *Argon2Params) error {
	user, err := db.QueryUserByName(c.ctx, req.Username)
	if err != nil {
		return err
	}
	log.Println(user)
	if user.Flag == true {
		return errno.ErrUserAlreadyExist
	}
	//对密码进行加密
	password, err := generateFromPassword(req.Password, argon2Params)
	if err != nil {
		return err
	}
	user.Username = req.Username
	user.Password = password
	user.Email = req.Email
	codeV, err := redisMiddleware.GetEmailCode("email", user.Email)
	if err != nil {
		return err
	}
	if req.Code != codeV {
		return errno.NewErrNo(100601, "验证码错误")
	}
	return db.CreateUser(c.ctx, user)
}

// generateFromPassword 对密码进行加密
func generateFromPassword(password string, argon2Params *Argon2Params) (encodedHash string, err error) {
	salt, err := generateRandomBytes(argon2Params.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, argon2Params.Iterations, argon2Params.Memory, argon2Params.Parallelism, argon2Params.KeyLength)

	// Base64 encode the salt and hashed password.
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, argon2Params.Memory, argon2Params.Iterations, argon2Params.Parallelism, base64Salt, base64Hash)

	return encodedHash, nil
}

// generateRandomBytes returns a random bytes.
func generateRandomBytes(saltLength uint32) ([]byte, error) {
	buf := make([]byte, saltLength)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
