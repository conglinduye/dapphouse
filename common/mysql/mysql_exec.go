package mysql

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lexkong/log"
)

var DBHost = "host"
var DBPort = "port"
var DBSchema = "schema"
var DBName = "user"
var DBPass = "password"


//数据库的连接配置
type DBParam struct {
	Mode         string
	ConnSQL      string
	MaxOpenconns int
	MaxIdleConns int
}

//连接DB的实例对象
var dbInstance *TronDB

// Initialize 初始化
func Initialize(host, port, schema, username, password string) {
	DBHost = strings.TrimSpace(host)
	DBPort = strings.TrimSpace(port)
	DBSchema = strings.TrimSpace(schema)
	DBName = strings.TrimSpace(username)
	DBPass = strings.TrimSpace(password)
}

//GetDatabase Get一个连接的数据库对象
func GetDatabase() (*TronDB, error) {
	return retrieveDatabase()
}

//RefreshDatabase 刷新DB的连接
func retrieveDatabase() (*TronDB, error) {
	defer CatchError()
	if nil == dbInstance {
		//连接数据库的参数
		para := GetMysqlConnectionInfo()
		//打开这个DB对象
		dbPtr, err := OpenDB(para.Mode, para.ConnSQL)
		if err != nil {
			return nil, err
		}
		if dbPtr == nil {
			return nil, errors.New("db not connected")
		}
		//设置连接池信息
		dbPtr.SetConnsParam(para.MaxOpenconns, para.MaxIdleConns)
		dbInstance = dbPtr
	}
	//测试一下是否是连接成功的
	if err := dbInstance.Ping(); err != nil {
		//dbInstance.Close()
		dbInstance = nil
		return nil, err
	}
	return dbInstance, nil
}

//GetMysqlConnectionInfo 获取连接mysql的相关信息
func GetMysqlConnectionInfo() DBParam {
	dbConfig := DBParam{
		Mode:         string("mysql"),
		ConnSQL:      fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", DBName, DBPass, DBHost, DBPort, DBSchema),
		MaxOpenconns: 10,
		MaxIdleConns: 10,
	}
	return dbConfig
}



//QueryTableData 查询数据库数据
func QueryTableData(strSQL string) (*TronDBRows, error) {
	//获取数据库对象
	var dbPtr *TronDB
	var err error
	if dbPtr, err = GetDatabase(); err != nil {
		return nil, errors.New("get database error")
	}

	//查询数据集
	rows, err := dbPtr.Select(strSQL)
	if err != nil {
		return rows, err
	}
	return rows, err
}


//QueryTableDataCount 返回某个表的记录个数
func QueryTableDataCount(tableName string) (int64, error) {
	rowCount := int64(0) //返回的数据库表行数

	//判断输入参数
	if len(tableName) < 0 {
		return 0, errors.New("parameter invalid")
	}

	//获取数据库对象
	var dbPtr *TronDB
	var err error
	if dbPtr, err = GetDatabase(); err != nil {
		return 0, errors.New("db not connected")
	}

	strSQL := "select count(*) as rowcounts from " + tableName
	var data *TronDBRows
	if data, err = dbPtr.Select(strSQL); err != nil {
		return 0, errors.New("query table data count error")
	}

	if data.NextT() {
		strValue := data.GetField("rowcounts")
		if count, err := strconv.ParseInt(strValue, 10, 64); err != nil {
			return 0, errors.New("query table data count error")
		} else {
			rowCount = count //set the count
		}
	}

	return rowCount, nil
}

//QueryTableDataCount 返回某个表的记录个数
func QuerySQLViewCount(strSQLView string) (int64, error) {
	rowCount := int64(0) //返回的数据库表行数

	//判断输入参数
	if len(strSQLView) < 0 {
		return 0, errors.New("parameter invalid")
	}

	//获取数据库对象
	var dbPtr *TronDB
	var err error
	if dbPtr, err = GetDatabase(); err != nil {
		return 0, errors.New("db not connected")
	}

	strSQL := fmt.Sprintf(`
		select count(*) as rowcounts from (
			%s
		) newtableName`, strSQLView)
	var data *TronDBRows
	if data, err = dbPtr.Select(strSQL); err != nil {
		return 0, errors.New("query sql view count error")
	}

	if data.NextT() {
		strValue := data.GetField("rowcounts")
		if count, err := strconv.ParseInt(strValue, 10, 64); err != nil {
			return 0, errors.New("query sql view count error")
		} else {
			rowCount = count //set the count
		}
	}

	return rowCount, nil
}

//ExecuteSQLCommand 执行insert update, delete操作,依次返回 插入消息的主键，影响的条数，错误对象
func ExecuteSQLCommand(strSQL string) (int64, int64, error) {
	var key int64
	var rows int64
	var err error
	var dbPtr *TronDB

	//获取数据库对象
	if dbPtr, err = GetDatabase(); err != nil {
		return 0, 0, errors.New("db not connected")
	}

	//执行语句
	if key, rows, err = dbPtr.Execute(strSQL); err != nil {
		return key, rows, errors.New("execute sql command error")
	}

	return key, rows, err
}


// CatchError 捕获异常
func CatchError() {
	if err := recover(); err != nil {
		log.Error("catch error", errors.New("recover error"))
	}
}
