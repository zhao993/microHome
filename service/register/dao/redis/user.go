package redis

import (
	"github.com/gomodule/redigo/redis"
)

// CheckImgCode 校验图片验证码
func CheckImgCode(uuid, imgCode string) bool {
	//从连接池中获取一条连接
	conn := RedisPool.Get()
	defer conn.Close()
	//获取redis验证码
	img, err := redis.String(conn.Do("get", uuid))
	if err != nil {
		return false
	}
	return img == imgCode
}

// SaveSmsCode 存储短信验证码
func SaveSmsCode(phone, smsCode string) (err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	//存储短信验证码
	_, err = conn.Do("setex", phone+"_code", 60*5, smsCode)
	return
}

// CheckSmsCode 校验短信验证码
func CheckSmsCode(phone, smsCode string) bool {
	conn := RedisPool.Get()
	defer conn.Close()
	//获取redis验证码
	sms, err := redis.String(conn.Do("get", phone+"_code"))
	if err != nil {
		return false
	}
	if sms == smsCode {
		return true
	}
	return false
}
