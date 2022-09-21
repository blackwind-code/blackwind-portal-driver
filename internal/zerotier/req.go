package zerotier

type ZT_Member_Approve struct {
	Authorized bool `json:"authorized"`
}

type ZT_Member_Change_Tag struct {
	Tags [][]int `json:"tags"`
}

type Device_Create struct {
	ZTAddress string `json:"zt_address"`
}

type Device_Update struct {
	ZTAddress  string `json:"zt_address"`
	DeviceType int    `json:"device_type"`
}

type Device_Delete struct {
	ZTAddress string `json:"zt_address"`
}
