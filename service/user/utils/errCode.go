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
