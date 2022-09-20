package zerotier

type ZT_Member_Approve struct {
	Authorized bool `json:"authorized"`
}

type ZT_Member_Change_Tag struct {
	Tags [][]int `json:"tags"`
}
