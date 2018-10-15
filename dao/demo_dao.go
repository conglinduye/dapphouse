package dao

import (
	"dapphouse/entity"
	"github.com/lexkong/log"
	"dapphouse/common/mysql"
	"dapphouse/util"
	"fmt"
)

// QueryDemoList
func QueryDemoList(strSQL, filterSQL, sortSQL, pageSQL string) (*entity.QueryDemoResp, error) {
	strFullSQL := strSQL + " " + filterSQL + " " + sortSQL + " " + pageSQL
	log.Info(strFullSQL)
	dataPtr, err := mysql.QueryTableData(strFullSQL)
	if err != nil {
		log.Errorf(err,"QueryDemoList error")
		return nil, err
	}

	demoResp := &entity.QueryDemoResp{}
	demoList := make([]*entity.Demo, 0)

	for dataPtr.NextT() {
		demo := &entity.Demo{}
		demo.Id = util.ConvertDBValueToInt64(dataPtr.GetField("id"))
		demo.Username = dataPtr.GetField("username")
		demo.Password = dataPtr.GetField("password")
		demo.CreateTime = dataPtr.GetField("create_time")
		demo.UpdateTime = dataPtr.GetField("update_time")
		demoList = append(demoList, demo)

	}

	var total = int64(len(demoList))
	total, err = mysql.QuerySQLViewCount(strSQL + " " + filterSQL)
	if err != nil {
		log.Errorf(err,"QuerySQLViewCount error sql:%s", strSQL)
	}

	demoResp.Total = total
	demoResp.Data = demoList

	return demoResp, nil
}

func QueryDemoById(id string) (*entity.Demo, error) {
	strSQL := fmt.Sprintf(`
		select id, username, password, create_time, update_time from demo where id = %v`, id)
	dataPtr, err := mysql.QueryTableData(strSQL)
	if err != nil {
		log.Errorf(err,"QueryDemoById error")
		return nil, err
	}

	demo := &entity.Demo{}
	for dataPtr.NextT() {
		demo.Id = util.ConvertDBValueToInt64(dataPtr.GetField("id"))
		demo.Username = dataPtr.GetField("username")
		demo.Password = dataPtr.GetField("password")
		demo.CreateTime = dataPtr.GetField("create_time")
		demo.UpdateTime = dataPtr.GetField("update_time")

	}
	return demo, nil
}

// DeleteDemoById
func DeleteDemoById(id string) error {
	strSQL := fmt.Sprintf(`
		delete from demo where id = %v`, id)
	_, _, err := mysql.ExecuteSQLCommand(strSQL)
	if err != nil {
		log.Errorf(err, "DeleteDemoById error, sql:%s", strSQL)
		return err
	}
	log.Info("DeleteDemoById success")
	return nil
}

// AddDemo
func AddDemo(demo *entity.Demo) error {
	strSQL := fmt.Sprintf(`
		insert into demo 
		(username, password)
		values('%v', '%v')`,
		demo.Username, demo.Password)
	_, _, err := mysql.ExecuteSQLCommand(strSQL)
	if err != nil {
		log.Errorf(err, "AddDemo error:%s", strSQL)
		return err
	}
	log.Info("AddDemo success")
	return nil
}

// InsertDemo
func UpdateDemo(demo *entity.Demo) error {
	strSQL := fmt.Sprintf(`
		update demo set username = '%v', password='%v' where id = %v`,
		demo.Username, demo.Password, demo.Id)
	_, _, err := mysql.ExecuteSQLCommand(strSQL)
	if err != nil {
		log.Errorf(err, "UpdateDemo error:%s", strSQL)
		return err
	}
	log.Info("UpdateDemo success")
	return nil
}
