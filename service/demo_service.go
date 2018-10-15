package service

import (
	"dapphouse/entity"
	"fmt"
	"dapphouse/dao"
)

// QueryDemoList
func QueryDemoList(req *entity.QueryDemoReq) (*entity.QueryDemoResp, error) {
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
func QueryDemoById(id string) (*entity.Demo, error)  {
	return dao.QueryDemoById(id)
}

// DeleteDemoById
func DeleteDemoById(id string) error {
	return dao.DeleteDemoById(id)
}

// AddDemo
func AddDemo(demo *entity.Demo) error {
	return dao.AddDemo(demo)
}

// UpdateDemo
func UpdateDemo(demo *entity.Demo) error {
	return dao.UpdateDemo(demo)
}


