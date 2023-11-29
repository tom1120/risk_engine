// Copyright (c) 2023
//
// @author 贺鹏Kavin
// 微信公众号:技术岁月
// https://github.com/skyhackvip/risk_engine
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
package util

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var intPattern = regexp.MustCompile(`^\d+$`)
var floatPattern = regexp.MustCompile(`^(\d+)(\.+)(\d+)$`)

var shortTimePattern = regexp.MustCompile(`^(\d){4}-(\d){2}-(\d){2}$`)
var longTimePattern = regexp.MustCompile(`^(\d){4}-(\d){2}-(\d){2} (\d){2}:(\d){2}:(\d){2}$`)

var TRUE = "true"
var FALSE = "false"

//infer type according to value
func GetType(val interface{}) (string, error) {
	switch val.(type) {
	case string:
		if IsInt(val) {
			return configs.INT, nil
		}
		if IsFloat(val) {
			return configs.FLOAT, nil
		}
		if IsBool(val) {
			return configs.BOOL, nil
		}
		if IsDate(val) {
			return configs.DATE, nil
		}
		return configs.STRING, nil
	case int:
		return configs.INT, nil
	case int32:
		return configs.INT, nil
	case int64:
		return configs.INT, nil
	case float32:
		if IsFloat32Int(val.(float32)) {
			return configs.INT, nil
		}
		return configs.FLOAT, nil
	case float64:
		if IsFloat64Int(val.(float64)) {
			return configs.INT, nil
		}
		return configs.FLOAT, nil
	case bool:
		return configs.BOOL, nil
	case time.Time:
		return configs.DATE, nil
	case []interface{}: //slice
		return configs.ARRAY, nil
	case map[string]interface{}: //only support map key is string
		return configs.MAP, nil
	}
	//array type like [3]int
	if reflect.TypeOf(val).Kind() == reflect.Array {
		return configs.ARRAY, nil
	}
	return configs.DEFAULT, errcode.ErrorFeatureTypeUnknow
}

//match type  int match float
func MatchType(typeA, typeB string) bool {
	if typeA == configs.INT {
		typeA = configs.FLOAT
	}
	if typeB == configs.INT {
		typeB = configs.FLOAT
	}
	return typeA == typeB
}

//is number type
func isNum(val interface{}) bool {
	return IsInt(val) || IsFloat(val)
}

//is int type
func IsInt(val interface{}) bool {
	switch val.(type) {
	case int8, int16, int, int32, int64:
		return true
	case string:
		return intPattern.MatchString(val.(string))
	case float32:
		return IsFloat32Int(val.(float32))
	case float64:
		return IsFloat64Int(val.(float64))
	default:
		return false
	}
}

//1.0 = 1
func IsFloat64Int(val float64) bool {
	return val == float64(int(val))
}

//1.0 = 1
func IsFloat32Int(val float32) bool {
	return val == float32(int(val))
}

//is float type
func IsFloat(val interface{}) bool {
	switch val.(type) {
	case float32:
		return true
	case float64:
		return true
	case string:
		return floatPattern.MatchString(val.(string))
	default:
		return false
	}
}

//is bool type
func IsBool(val interface{}) bool {
	switch val.(type) {
	case bool:
		return true
	case string:
		if strings.ToLower(val.(string)) == TRUE || strings.ToLower(val.(string)) == FALSE { //true, True, TRUE
			return true
		}
		return false
	default:
		return false
	}
}

//is date type
func IsDate(val interface{}) bool {
	switch val.(type) {
	case time.Time:
		return true
	case string:
		return (shortTimePattern.MatchString(val.(string)) || longTimePattern.MatchString(val.(string)))
	default:
		return false
	}
}

//other type convert to bool
func ToBool(val interface{}) (bool, error) {
	if !IsBool(val) {
		return false, errcode.ErrorTypeConvert
	}
	switch val.(type) {
	case bool:
		return val.(bool), nil
	case string:
		if strings.ToLower(val.(string)) == TRUE {
			return true, nil
		}
		if strings.ToLower(val.(string)) == FALSE {
			return false, nil
		}
	}
	return false, errcode.ErrorTypeConvert
}

//other type convert to data
func ToDate(val interface{}) (time.Time, error) {
	if !IsDate(val) {
		return time.Time{}, errcode.ErrorTypeConvert
	}
	switch val.(type) {
	case time.Time:
		return val.(time.Time), nil
	case string:
		if shortTimePattern.MatchString(val.(string)) {
			return time.Parse(configs.DATE_FORMAT, val.(string))
		}
		if longTimePattern.MatchString(val.(string)) {
			return time.Parse(configs.DATE_FORMAT_DETAIL, val.(string))
		}
	}
	return time.Time{}, errcode.ErrorTypeConvert
}

//other type convert to string
func ToString(val interface{}) (ret string, err error) {
	switch val.(type) {
	case string:
		ret = val.(string)
	case int8, int16, int, int32, int64, uint8, uint16, uint32, uint64:
		if v, err := ToInt64(val); err != nil {
			err = errcode.ErrorTypeConvert
		} else {
			ret = strconv.FormatInt(v, 10)
		}
	case float64:
		ret = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case bool:
		ret = strconv.FormatBool(val.(bool))
	default:
		err = errcode.ErrorTypeConvert
	}
	return
}

//other type convert to int64
func ToInt64(val interface{}) (ret int64, err error) {
	switch val.(type) {
	case int8:
		ret = int64(val.(int8))
	case int16:
		ret = int64(val.(int16))
	case int:
		ret = int64(val.(int))
	case int32:
		ret = int64(val.(int32))
	case int64:
		ret = val.(int64)
	case float64:
		ret = int64(val.(float64))
	case string:
		if v, err := strconv.Atoi(val.(string)); err != nil {
			err = errcode.ErrorTypeConvert
		} else {
			ret = int64(v)
		}
	default:
		err = errcode.ErrorTypeConvert
	}
	return
}

//other type convert to int
func ToInt(val interface{}) (ret int, err error) {
	if v, err := ToInt64(val); err == nil {
		ret = int(v)
	} else {
		err = errcode.ErrorTypeConvert
	}
	return
}

//other type convert to float64
func ToFloat64(val interface{}) (ret float64, err error) {
	switch val.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64:
		if retInt, e := ToInt(val); e != nil {
			err = e
			return
		} else {
			ret = float64(retInt)
		}
	case float32:
		ret = float64(val.(float32))
	case float64:
		ret = val.(float64)
	case string:
		if ret, err = strconv.ParseFloat(val.(string), 64); err != nil {
			err = errcode.ErrorTypeConvert
			return
		}
	//case reflect.Value:
	//	ret, err = RVToFloat64(val.(reflect.Value))
	//	return
	default:
		err = errcode.ErrorTypeConvert
	}
	return
}

//reflect value convert to float64
func RVToFloat64(val reflect.Value) (ret float64, err error) {
	var num interface{}
	switch val.Kind() {
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Uint:
		num = val.Uint()

	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Int:
		num = val.Int()

	case reflect.Float64:
		fallthrough
	case reflect.Float32:
		num = val.Float()

	case reflect.String:
		num = val.String()
	default:
		err = errcode.ErrorTypeConvert
	}
	if err != nil {
		return
	}
	return ToFloat64(num)
}
