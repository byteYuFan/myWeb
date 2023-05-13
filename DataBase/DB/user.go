package db

import (
	"context"
	"time"
)

type User struct {
	ID       int       `gorm:"primaryKey"`
	Username string    `gorm:"uniqueIndex:idx_username"`
	Email    string    `gorm:"uniqueIndex:idx_email"`
	Password string    `gorm:"not null"`
	CreateAt time.Time `gorm:"autoCreateTime"`
	Flag     bool      `gorm:"default:true;index:idx_flag"`
}

// TableName 返回User表在数据库中对应的表名称
func (User) TableName() string {
	return "user_register_info"
}

// GetUserInfoByUserName 根据用户名查找用户信息
func GetUserInfoByUserName(ctx context.Context, username string) (*User, error) {
	res := new(User)
	if err := USERDB.WithContext(ctx).Where("username=? AND flag = ?", username, true).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserInfoByUserId 根据用户id查找用户信息
func GetUserInfoByUserId(ctx context.Context, id int64) (*User, error) {
	res := new(User)
	if err := USERDB.WithContext(ctx).Where("id = ? AND flag =?", id, true).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserInfoByUserEmail 根据用户邮箱查找用户信息
func GetUserInfoByUserEmail(ctx context.Context, email string) (*User, error) {
	res := new(User)
	if err := USERDB.WithContext(ctx).Where("email=? AND flag=?", email, true).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	if err := USERDB.WithContext(ctx).Where("id in ? AND flag = ?", userIDs, true).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser 创建用户
func CreateUser(ctx context.Context, user *User) error {
	return USERDB.WithContext(ctx).Create(user).Error
}

// QueryUserByName 根据用户名查找用户信息
func QueryUserByName(ctx context.Context, username string) (*User, error) {
	res := new(User)
	if err := USERDB.WithContext(ctx).Where("username = ?", username).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
