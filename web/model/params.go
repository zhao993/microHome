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

type ParamAuth struct {
	RealName string `json:"real_name"`
	IDCard   string `json:"id_card"`
}

type ParamUsers struct {
	ID        int32  `json:"user_id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	RealName  string `json:"real_name"`
	IDCard    string `json:"id_card"`
	AvatarUrl string `json:"avatar_url"`
}

type ParamHouseStu struct {
	Acreage   string   `json:"acreage"`
	Address   string   `json:"address"`
	AreaId    string   `json:"area_id"`
	Beds      string   `json:"beds"`
	Capacity  string   `json:"capacity"`
	Deposit   string   `json:"deposit"`
	Facility  []string `json:"facility"`
	MaxDays   string   `json:"max_days"`
	MinDays   string   `json:"min_days"`
	Price     string   `json:"price"`
	RoomCount string   `json:"room_count"`
	Title     string   `json:"title"`
	Unit      string   `json:"unit"`
}

type OrderStu struct {
	EndDate   string `json:"end_date"`
	HouseId   string `json:"house_id"`
	StartDate string `json:"start_date"`
}

type StatusStu struct {
	Action string `json:"action"`
	Reason string `json:"reason"`
}
