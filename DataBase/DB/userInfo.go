package db

import (
	"context"
)

type UserInfo struct {
	Id         int64  `gorm:"column:id"`         // 用户ID
	Name       string `gorm:"column:name"`       // 用户名
	Password   string `gorm:"column:password"`   // 密码
	Age        int32  `gorm:"column:age"`        // 年龄
	Profession string `gorm:"column:profession"` // 职业
	Department string `gorm:"column:department"` // 院系
	Province   string `gorm:"column:province"`   // 省份
	City       string `gorm:"column:city"`       // 城市
	Flag       bool   `gorm:"column:flag"`       // 标志位
}

func (UserInfo) TableName() string {
	return "user_info"
}
func init() {
	USERDB.AutoMigrate(&UserInfo{})
}
func GetUserInfo(ctx context.Context, id int64) (*UserInfo, error) {
	user := new(UserInfo)
	if err := USERDB.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func ModifyPassword(ctx context.Context, id int64, password string) error {
	return USERDB.WithContext(ctx).Model(&User{}).Where("id=?", id).Update("password", password).Error
}

func UpdateUserInfo(ctx context.Context, id int64, user *UserInfo) error {
	return USERDB.WithContext(ctx).Model(new(UserInfo)).Where("id=?", id).Updates(user).Error
}

func UpdateUserPassword(ctx context.Context, id int64, password string) error {
	return USERDB.WithContext(ctx).Model(&User{}).Where("id=?", id).Update("password", password).Error
}
