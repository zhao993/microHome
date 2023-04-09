package utils

//定义错误状态码

type MyCode string

const (
	RecodeOk         MyCode = "0"
	RecodeDbErr      MyCode = "4001"
	RecodeNoData     MyCode = "4002"
	RecodeDataExist  MyCode = "4003"
	RecodeDataErr    MyCode = "4004"
	RecodeSessionErr MyCode = "4101"
	RecodeLoginErr   MyCode = "4102"
	RecodeParamErr   MyCode = "4103"
	RecodeUserOnErr  MyCode = "4104"
	RecodeRoleErr    MyCode = "4105"
	RecodePwdErr     MyCode = "4106"
	RecodeUserErr    MyCode = "4107"
	RecodeSmsErr     MyCode = "4108"
	RecodeImgErr     MyCode = "4109"
	RecodeMobileErr  MyCode = "4110"

	RecodeReqErr MyCode = "4201"
	RecodeIpErr  MyCode = "4202"

	RecodeThirdErr MyCode = "4301"
	RecodeIoErr    MyCode = "4302"

	RecodeServerErr MyCode = "4500"
	RecodeUnKnowErr MyCode = "4501"
)

var recodeText = map[MyCode]string{
	RecodeOk:        "成功",
	RecodeDbErr:     "数据库操作错误",
	RecodeNoData:    "无数据",
	RecodeDataExist: "数据已存在",
	RecodeDataErr:   "数据错误",

	RecodeSessionErr: "用户未登录",
	RecodeLoginErr:   "用户登录失败",
	RecodeParamErr:   "参数错误",
	RecodeUserOnErr:  "用户已注册",
	RecodeRoleErr:    "用户身份错误",
	RecodePwdErr:     "密码错误",
	RecodeUserErr:    "用户不存在或密码错误",
	RecodeSmsErr:     "短信验证码错误",
	RecodeImgErr:     "图片验证码错误",
	RecodeMobileErr:  "手机号格式错误",

	RecodeReqErr: "非法请求",
	RecodeIpErr:  "IP受限",

	RecodeThirdErr: "第三方系统错误",
	RecodeIoErr:    "文件读写错误",

	RecodeServerErr: "服务器内部错误",
	RecodeUnKnowErr: "未知错误",
}

func (c MyCode) RecodeText() string {
	msg, ok := recodeText[c]
	if ok {
		return msg
	}
	return recodeText[RecodeUnKnowErr]
}
