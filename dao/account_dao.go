package dao

import (
	"dapphouse/entity"
	"fmt"
	"dapphouse/common/mysql"
	"github.com/lexkong/log"
	"dapphouse/util"
)

func QueryUserAccount(userAccount *entity.UserAccount) (*entity.UserAccount, error) {
	strSQL := fmt.Sprintf(`
		select id, email, salt, pwd_crypt, status, create_time, update_time 
		from user_account where email = '%v' and pwd_crypt = '%v'`, userAccount.Email, userAccount.PwdCrypt)
	dataPtr, err := mysql.QueryTableData(strSQL)
	if err != nil {
		log.Errorf(err,"QueryUserAccount error sql:%v", strSQL)
		return nil, err
	}

	newUserAccount := &entity.UserAccount{}
	for dataPtr.NextT() {
		newUserAccount.Id = util.ConvertDBValueToInt64(dataPtr.GetField("id"))
		newUserAccount.Email = dataPtr.GetField("email")
		newUserAccount.Salt = dataPtr.GetField("salt")
		newUserAccount.PwdCrypt = dataPtr.GetField("pwd_crypt")
		newUserAccount.Status = util.ConvertDBValueToInt(dataPtr.GetField("status"))
		newUserAccount.CreateTime = dataPtr.GetField("create_time")
		newUserAccount.UpdateTime = dataPtr.GetField("update_time")

	}
	return newUserAccount, nil
}

// AddUserAccount
func AddUserAccount(userAccount *entity.UserAccount) error {
	strSQL := fmt.Sprintf(`
		insert into user_account 
		(email, salt, pwd_crypt, status)
		values('%v', '%v', '%v', %v)`,
		userAccount.Email, userAccount.Salt, userAccount.PwdCrypt, userAccount.Status)
	_, _, err := mysql.ExecuteSQLCommand(strSQL)
	if err != nil {
		log.Errorf(err, "AddUserAccount error sql:%s", strSQL)
		return err
	}
	log.Info("AddUserAccount success")
	return nil
}

// UpdateUserAccount
func UpdateUserAccount(userAccount *entity.UserAccount) error {
	strSQL := fmt.Sprintf(`
		update user_account set salt = '%v', pwd_crypt = '%v' where email = '%v''`,
		userAccount.Salt, userAccount.PwdCrypt, userAccount.Email)
	_, _, err := mysql.ExecuteSQLCommand(strSQL)
	if err != nil {
		log.Errorf(err, "UpdateUserAccount error sql:%s", strSQL)
		return err
	}
	log.Info("UpdateUserAccount success")
	return nil
}

// UpdateUserAccountStatus
func UpdateUserAccountStatus(userAccount *entity.UserAccount) error {
	strSQL := fmt.Sprintf(`
		update user_account set status = %v where email = '%v''`,
		userAccount.Status, userAccount.Email)
	_, _, err := mysql.ExecuteSQLCommand(strSQL)
	if err != nil {
		log.Errorf(err, "UpdateUserAccountStatus error sql:%s", strSQL)
		return err
	}
	log.Info("UpdateUserAccountStatus success")
	return nil
}



