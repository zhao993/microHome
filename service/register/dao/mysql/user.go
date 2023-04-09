package mysql

import (
	"crypto/md5"
	"errors"
	"fmt"
	"register/model"
)

// RegisterUser 注册用户信息
func RegisterUser(mobile, pwd string) error {
	//检查手机号是否已注册
	var user model.User
	var count int64
	DB.Where("mobile =?", mobile).First(user).Count(&count)
	if count != 0 {
		return errors.New("手机号已注册")
	}
	user.Mobile = mobile
	user.Name = mobile
	user.PasswordHash = passwordEncrypt(pwd)
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// FindUser 查找用户
func FindUser(mobile, password string) (string, bool) {
	var user model.User
	var count int64
	DB.Where("mobile = ? AND password_hash = ?", mobile, passwordEncrypt(password)).Find(&user).Count(&count)
	if count == 0 {
		return "", false
	}
	return user.Name, true
}

// PasswordEncrypt 密码加密
func passwordEncrypt(password string) string {
	//用md5加密
	data := []byte(password) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
