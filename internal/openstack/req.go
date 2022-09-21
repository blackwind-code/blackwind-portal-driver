package openstack

type OS_User_Create struct {
	Email       string `json:"email"`
	PassworHash string `json:"password_hash"`
}

type OS_User_Update struct {
	OldEmail     string `json:"old_email"`
	NewEmail     string `json:"new_email"`
	PasswordHash string `json:"password_hash"`
}

type OS_User_Delete struct {
	Email string `json:"email"`
}
