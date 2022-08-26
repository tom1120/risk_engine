package util

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"regexp"
	"strconv"
	"time"
)

var intPattern = regexp.MustCompile(`^\d+$`)
var floatPattern = regexp.MustCompile(`^(\d+)(\.+)(\d+)$`)

var shortTimePattern = regexp.MustCompile(`^(\d){4}-(\d){2}-(\d){2}$`)
var longTimePattern = regexp.MustCompile(`^(\d){4}-(\d){2}-(\d){2} (\d){2}:(\d){2}:(\d){2}$`)

func GetType(val interface{}) (string, error) {
	switch val.(type) {
	case string:
		if IsInt(val) {
			return "int", nil
		}
		if IsFloat(val) {
			return "float", nil
		}
		if IsBool(val) {
			return "bool", nil
		}
		if IsDate(val) {
			return "date", nil
		}
		return "string", nil
	case int:
		return "int", nil
	case int32:
		return "int", nil
	case int64:
		return "int", nil
	case float32:
		if IsFloat32Int(val.(float32)) {
			return "int", nil
		}
		return "float", nil
	case float64:
		if IsFloat64Int(val.(float64)) {
			return "int", nil
		}
		return "float", nil
	case bool:
		return "bool", nil
	case time.Time:
		return "date", nil
	case []interface{}:
		return "array", nil
	case map[string]interface{}:
		return "map", nil
	}
	return "default", errcode.ErrorFeatureTypeUnknow
}

//date type jundge
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

func IsFloat64Int(val float64) bool {
	return val == float64(int(val))
}

func IsFloat32Int(val float32) bool {
	return val == float32(int(val))
}

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

func IsBool(val interface{}) bool {
	switch val.(type) {
	case bool:
		return true
	case string:
		if val.(string) == "true" || val.(string) == "false" {
			return true
		}
		return false
	default:
		return false
	}
}

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

//data type covert
func StringToDate(val string) (time.Time, error) {
	if shortTimePattern.MatchString(val) {
		return time.Parse(configs.DATE_FORMAT, val)
	}
	if longTimePattern.MatchString(val) {
		return time.Parse(configs.DATE_FORMAT_DETAIL, val)
	}
	return time.Time{}, errcode.ErrorTypeConvert
}

func ToString(val interface{}) (ret string, err error) {
	switch val.(type) {
	case string:
		ret = val.(string)
	case int8, int16, int, int32, int64:
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

func ToInt(val interface{}) (ret int, err error) {
	if v, err := ToInt64(val); err == nil {
		ret = int(v)
	} else {
		err = errcode.ErrorTypeConvert
	}
	return
}
