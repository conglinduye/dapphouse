package service

import (
	"dapphouse/util"
	"dapphouse/entity"
	"dapphouse/dao"
)

func AddUserAccount(uaRegister *entity.UserAccountRegister) error {
	salt := util.GetRandomString(8)
	pwdCrypt := util.MD5(salt + uaRegister.Password)
	newUserAccount := &entity.UserAccount{}
	newUserAccount.Email = uaRegister.Email
	newUserAccount.Salt = salt
	newUserAccount.PwdCrypt = pwdCrypt
	newUserAccount.Status = 0
	return  dao.AddUserAccount(newUserAccount)
}

// QueryUserAccountByEmail
func QueryUserAccountByEmail(email string) (*entity.UserAccount, error) {
	return dao.QueryUserAccountByEmail(email)
}
