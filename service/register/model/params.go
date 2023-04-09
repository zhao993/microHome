package model

type ParamUser struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	SmsCode  string `json:"sms_code"`
}

type ParamLogin struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type ParamUsers struct {
	ID        int32  `json:"user_id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	RealName  string `json:"real_name"`
	IDCard    string `json:"id_card"`
	AvatarUrl string `json:"avatar_url"`
}
