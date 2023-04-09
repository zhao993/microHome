package handler

import (
	"context"
	"fmt"
	"math/rand"
	"register/dao/mysql"
	"register/dao/redis"
	pb "register/proto"
	"register/utils"
	"time"
)

type Register struct{}

// SendSms 发送验证码
func (e *Register) SendSms(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	//校验图片验证码是否正确
	if redis.CheckImgCode(req.Uuid, req.ImgCode) == false {
		rsp.ErrCode = string(utils.RecodeImgErr)
		return nil
	}
	//发送短信验证码
	rand.Seed(time.Now().UnixNano()) //播种随机数种子
	//随机生成6位数
	smsCode := fmt.Sprintf("%06d", rand.Int31n(1000000))
	err := utils.GetSms(req.Phone, smsCode)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeSmsErr)
		return nil
	}
	err = redis.SaveSmsCode(req.Phone, smsCode)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeSmsErr)
		return nil
	}
	rsp.ErrCode = string(utils.RecodeOk)
	return nil
}

// Register 注册用户
func (e *Register) Register(ctx context.Context, req *pb.RegReq, rsp *pb.CallResponse) error {
	//检验短信验证码是否正确
	if redis.CheckSmsCode(req.Mobile, req.SmsCode) == false {
		rsp.ErrCode = string(utils.RecodeSmsErr)
		return nil
	}
	//注册用户
	err := mysql.RegisterUser(req.Mobile, req.Password)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeUserOnErr)
		return nil
	}
	rsp.ErrCode = string(utils.RecodeOk)
	return nil
}

// Login 登录设置session
func (e *Register) Login(ctx context.Context, req *pb.LoginReq, rsp *pb.Response) error {
	userName, exist := mysql.FindUser(req.Mobile, req.Password)
	if exist {
		rsp.Name = userName
		rsp.ErrCode = string(utils.RecodeOk)
	} else {
		rsp.ErrCode = string(utils.RecodeUserErr)
	}
	return nil
}
