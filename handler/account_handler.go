package handler

import (
	"github.com/gin-gonic/gin"
	"dapphouse/entity"
	"dapphouse/common/errno"
	"github.com/lexkong/log"
	"dapphouse/service"
	"dapphouse/util"
)

func RegisterUserAccount(c *gin.Context) {
	var uaRegister entity.UserAccountRegister
	if err := c.Bind(&uaRegister); err != nil {
		log.Errorf(err, "RegisterUserAccount bind error")
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	if uaRegister.Email == ""  || uaRegister.Password == "" {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	userAccount, err := service.QueryUserAccountByEmail(uaRegister.Email)
	if err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}
	if userAccount.Email != "" {
		SendResponse(c, errno.ErrUserAccountExisted, nil)
		return
	}

	if err := service.AddUserAccount(&uaRegister); err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}

	SendResponse(c, nil, nil)
 }

 func LoginUserAccount(c *gin.Context) {
	 var uaLogin entity.UserAccountLogin
	 if err := c.Bind(&uaLogin); err != nil {
		 log.Errorf(err, "LoginUserAccount bind error")
		 SendResponse(c, errno.ErrBind, nil)
		 return
	 }

	 userAccount, err := service.QueryUserAccountByEmail(uaLogin.Email)
	 if err != nil {
		 SendResponse(c, errno.InternalServerError, nil)
		 return
	 }
	 if userAccount.Email == "" {
		 SendResponse(c, errno.ErrUserAccountPassword, nil)
		 return
	 }

	 salt := userAccount.Salt
	 pwdCrypt := util.MD5(salt + uaLogin.Password)
	 if pwdCrypt != userAccount.PwdCrypt {
		 SendResponse(c, errno.ErrUserAccountPassword, nil)
		 return
	 }

	 SendResponse(c, nil, nil)
 }