package handler

import (
	"github.com/gin-gonic/gin"

	"dapphouse/util"
	"github.com/lexkong/log"
	"dapphouse/common/errno"
	"dapphouse/service"
	"dapphouse/entity"
	"github.com/wlcy/tron/explorer/web/handler"
)

func QueryDemoList(c *gin.Context) {
	req := &entity.QueryDemoReq{}
	req.Start = util.ConvertStringToInt64(c.Query("start"), 0)
	req.Limit = util.ConvertStringToInt64Limit(c.Query("limit"), 40, 100)
	req.Username = c.Query("username")
	req.Password = c.Query("password")
	log.Infof("QueryDemoList req:%#v", req)

	queryResp, err := service.QueryDemoList(req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, nil, queryResp)
}

func QueryDemoById(c *gin.Context) {
	id := c.Param("id")
	log.Infof("QueryDemoById req:%s", id)

	queryResp, err := service.QueryDemoById(id)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, queryResp)
}

func AddDemo(c *gin.Context) {
	var demo entity.Demo
	if err := c.Bind(&demo); err != nil {
		log.Infof("AddDemo bind err:%#v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	log.Infof("AddDemo bind req:%#v", demo)
	if demo.Username == ""  || demo.Password == "" {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := service.AddDemo(&demo); err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}

func UpdateDemo(c *gin.Context) {
	var demo entity.Demo
	if err := c.Bind(&demo); err != nil {
		log.Infof("UpdateDemo bind err:%v", err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	log.Infof("UpdateDemo bind req:%v", demo)
	if demo.Username == ""  || demo.Password == "" {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := service.UpdateDemo(&demo); err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}


func DeleteDemoById(c *gin.Context) {
	id := c.Param("id")
	log.Infof("DeleteDemoById req:%s", id)
	if err := service.DeleteDemoById(id); err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}


