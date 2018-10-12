package service

import (
	"dapphouse/entity/demo"
	"fmt"
	"dapphouse/dao/demo"
)

// QueryDemoList
func QueryDemoList(req *demo.QueryReq) (*demo.QueryResp, error) {
	var filterSQL, sortSQL, pageSQL string
	strSQL := fmt.Sprintf(`select id, username, password, create_time, update_time from demo where 1=1 `)

	if req.Username != "" {
		filterSQL = fmt.Sprintf(" and username='%v'", req.Username)
	}
	if req.Password != "" {
		filterSQL = fmt.Sprintf(" and asset_name='%v'", req.Password)
	}

	sortSQL = fmt.Sprintf("order by id desc")
	pageSQL = fmt.Sprintf("limit %v, %v", req.Start, req.Limit)

	return dao.QueryDemoList(strSQL, filterSQL, sortSQL, pageSQL)
}

// QueryDemoById
func QueryDemoById(id string) (*demo.Demo, error)  {
	return dao.QueryDemoById(id)
}

// DeleteDemoById
func DeleteDemoById(id string) error {
	return dao.DeleteDemoById(id)
}

// AddDemo
func AddDemo(demo *demo.Demo) error {
	return dao.AddDemo(demo)
}

// UpdateDemo
func UpdateDemo(demo *demo.Demo) error {
	return dao.UpdateDemo(demo)
}


