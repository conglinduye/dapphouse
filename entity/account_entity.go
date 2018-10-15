package entity

// UserAccountRegister
type UserAccountRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserAccountActivate
type UserAccountActivate struct {
}

// UserAccountReq
type UserAccountLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserAccount
type UserAccount struct {
	Id         int64  `json:"id"`
	Email      string `json:"email"`
	Salt       string `json:"salt"`
	PwdCrypt   string `json:"pwdCrypt"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
