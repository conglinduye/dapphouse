package util

import (
	"strconv"
	"time"
	"errors"
)

//日期时间的格式
const DATEFORMAT = "2006-01-02"
const DATETIMEFORMAT = "2006-01-02 15:04:05"

//ConvertStringToTime 将时间字符串转化为时间对象
func ConvertStringToTime(strTimeString string) (time.Time, error) {
	var retTime time.Time
	if len(strTimeString) <= 0 {
		return retTime, errors.New("时间字符串无效：字符串为空值")
	}

	var err error
	retTime, err = time.Parse(DATETIMEFORMAT, strTimeString)
	if err != nil {
		return retTime, errors.New("时间字符串无效：" + strTimeString)
	}
	return retTime, nil
}

//ConvertStringToBoolean 将字符串转化为Boolean类型
func ConvertStringToBoolean(booleanString string, defaultValue bool) bool {
	retBool := defaultValue

	if len(booleanString) > 0 {
		b, err := strconv.ParseBool(booleanString)
		if err == nil {
			retBool = b
		}
	}
	return retBool
}

//ConvertStringToFloat 将字符串转化为Float类型
func ConvertStringToFloat(val string, defaultValue float64) float64 {
	ret := defaultValue

	if len(val) > 0 {
		b, err := strconv.ParseFloat(val, 64)
		if err == nil {
			ret = b
		}
	}
	return ret
}

//ConvertStringToInt 将字符串转化为int类型
func ConvertStringToInt(intString string, defaultValue int) int {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.Atoi(intString)
		if err == nil {
			retInt = i64
		}
	}
	return retInt
}

//ConvertStringToInt 将字符串转化为uint类型
func ConvertStringToUint(intString string, defaultValue uint) uint {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.ParseUint(intString, 10, 64)
		if err == nil {
			retInt = uint(i64)
		}
	}
	return retInt
}

//ConvertStringToInt32 将字符串转化为int32类型
func ConvertStringToInt32(intString string, defaultValue int32) int32 {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.ParseInt(intString, 10, 32)
		if err == nil {
			retInt = int32(i64)
		}
	}
	return retInt
}

//ConvertStringToInt64 将字符串转化为int64类型
func ConvertStringToInt64(intString string, defaultValue int64) int64 {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.ParseInt(intString, 10, 64)
		if err == nil {
			retInt = i64
		}
	}
	return retInt
}

func ConvertStringToInt64Limit(intString string, defaultValue, maxValue int64) int64 {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.ParseInt(intString, 10, 64)
		if err == nil {
			if i64 > maxValue {
				retInt = maxValue
			} else {
				retInt = i64
			}
		}
	}
	return retInt
}

//ConvertDBValueToString 返回字符串结果
func ConvertDBValueToString(colValue string) string {
	return colValue
}

//ConvertDBValueToInt64 返加int64结果
func ConvertDBValueToInt64(colValue string) int64 {
	retInt64 := int64(0)

	if len(colValue) > 0 {
		i64, err := strconv.ParseInt(colValue, 10, 64)
		if err == nil {
			retInt64 = i64
		}
	}
	return retInt64
}

//ConvertDBValueToUint64 返回uint64结果
func ConvertDBValueToUint64(colValue string) uint64 {
	retInt64 := uint64(0)

	if len(colValue) > 0 {
		i64, err := strconv.ParseUint(colValue, 10, 64)
		if err == nil {
			retInt64 = i64
		}
	}
	return retInt64
}

//ConvertDBValueToInt 返回int结果
func ConvertDBValueToInt(colValue string) int {
	retInt := int(0)

	if len(colValue) > 0 {
		i64, err := strconv.Atoi(colValue)
		if err == nil {
			retInt = i64
		}
	}
	return retInt
}

//ConvertDBValueToUint 返回int结果
func ConvertDBValueToUint(colValue string) uint {
	retInt := uint(0)

	if len(colValue) > 0 {
		ui, err := strconv.ParseUint(colValue, 10, 0)
		if err == nil {
			retInt = uint(ui)
		}
	}
	return retInt
}

//ConvertDBValueToFloat32 返回float32结果
func ConvertDBValueToFloat32(colValue string) float32 {
	retInt := float32(0)

	if len(colValue) > 0 {
		ui, err := strconv.ParseFloat(colValue, 32)
		if err == nil {
			retInt = float32(ui)
		}
	}
	return retInt
}

//ConvertDBValueToFloat64 返回float64结果
func ConvertDBValueToFloat64(colValue string) float64 {
	retInt := float64(0)

	if len(colValue) > 0 {
		ui, err := strconv.ParseFloat(colValue, 64)
		if err == nil {
			retInt = float64(ui)
		}
	}
	return retInt
}

//ConvertDBValueToBool 返回boolean结果
func ConvertDBValueToBool(colValue string) bool {
	retBool := false

	if len(colValue) > 0 {
		b, err := strconv.ParseBool(colValue)
		if err == nil {
			retBool = b
		}
	}
	return retBool
}

//ConvertDBValueToTime 返回时间结果
func ConvertDBValueToTime(colValue string) time.Time {
	var retTime time.Time

	if len(colValue) > 0 {
		t, err := time.ParseInLocation(DATETIMEFORMAT, colValue, time.Local) // 使用系统当前时区
		if err == nil {
			retTime = t
		}
	}
	return retTime
}


