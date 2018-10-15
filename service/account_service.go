package service

import (
	"dapphouse/util"
	"dapphouse/entity"
)

func UserAccountRegister(uaRegister *entity.UserAccountRegister) {
	salt := util.GetRandomString(8)
	pwdCrypt := util.MD5(salt + uaRegister.Password)
	userAccount := &entity.UserAccount{}
	userAccount.Email = uaRegister.Email
	userAccount.Salt = salt
	userAccount.PwdCrypt = pwdCrypt
	userAccount.Status = 0


}
