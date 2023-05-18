package cmd

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	db "myWeb/DataBase/DB"
	"myWeb/kitex_gen/userInfo"
	"myWeb/pkg/errno"
	redisMiddleware "myWeb/pkg/redis"
	"strings"
)

type UserInfoInstance struct {
	ctx context.Context
}

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// NewUserInfoInstance 获取实例
func NewUserInfoInstance(ctx context.Context) *UserInfoInstance {
	return &UserInfoInstance{ctx: ctx}
}
func (info *UserInfoInstance) GetUserInfo(req *userInfo.GetUserRequest) (*userInfo.User, error) {
	id := req.Id
	user := new(userInfo.User)
	u, err := db.GetUserInfo(info.ctx, id)
	if err != nil {
		return nil, err
	}
	user.Flag = u.Flag
	user.Password = u.Password
	user.Id = u.Id
	user.Flag = u.Flag
	user.Age = u.Age
	user.City = u.City
	user.Province = u.Province
	user.Department = u.Department
	user.Profession = u.Profession
	user.Name = u.Name
	return user, nil

}

func (info *UserInfoInstance) ResetPassword(req *userInfo.ResetPasswordRequest, argon2Params *Argon2Params) (bool, error) {
	code, err := redisMiddleware.GetEmailCode("email", req.Email)
	if err != nil {
		return false, errors.New(err.Error() + "未发送邮件")
	}
	if code != req.Credential {
		return false, errors.New("验证码错误")
	}
	//修改密码为默认
	user, err := db.QueryUserByEmail(info.ctx, req.Email)
	if err != nil {
		return false, err
	}
	if user.Username == "" || user.Flag == false {
		return false, errno.ErrUserNotFound
	}
	//对密码进行加密
	password, err := generateFromPassword(user.Username+"123456", argon2Params)
	if err != nil {
		return false, err
	}
	if err != db.ModifyPassword(context.Background(), user.ID, password) {
		return false, err
	}
	return true, nil
}

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

func (info *UserInfoInstance) UpdateUserInfo(req *userInfo.UpdateUserRequest) error {
	u := req
	user := new(db.UserInfo)
	user.Flag = u.Flag
	user.Password = ""
	user.Age = u.Age
	user.City = u.City
	user.Province = u.Province
	user.Department = u.Department
	user.Profession = u.Profession
	user.Name = u.Name
	user.Flag = true
	return db.UpdateUserInfo(info.ctx, req.Id, user)
}

func (info *UserInfoInstance) ChangeUserPassword(req *userInfo.ChangePasswordRequest, argon2Params *Argon2Params) error {
	u, err := db.GetUserInfoByUserId(info.ctx, req.Id)
	if err != nil {
		return err
	}
	passWordMatch, err := comparePasswordAndHash(req.OldPassword, u.Password)
	if err != nil {
		return err
	}
	if !passWordMatch {
		return errno.ErrPasswordIncorrect
	}
	n, err := generateFromPassword(req.ConfirmNewPassword, argon2Params)
	if err != nil {
		return err
	}
	err = db.UpdateUserPassword(info.ctx, req.Id, n)
	if err != nil {
		return err
	}
	return nil

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
