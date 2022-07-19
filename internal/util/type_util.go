package util

import (
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"regexp"
)

func GetFeatureType(val interface{}) (string, error) {
	switch val.(type) {
	case string:
		if IsInt(val.(string)) {
			return "int", nil
		}
		if IsFloat(val.(string)) {
			return "float", nil
		}
		if IsBool(val.(string)) {
			return "bool", nil
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
	}
	return "default", errcode.FeatureTypeErrorUnknow
}

var intPattern = regexp.MustCompile(`^\d+$`)
var floatPattern = regexp.MustCompile(`^(\d+)(\.+)(\d+)$`)

func IsInt(val string) bool {
	return intPattern.MatchString(val)
}

func IsFloat64Int(val float64) bool {
	return val == float64(int(val))
}

func IsFloat32Int(val float32) bool {
	return val == float32(int(val))
}

func IsFloat(val string) bool {
	return floatPattern.MatchString(val)
}

func IsBool(val string) bool {
	if val == "true" || val == "false" {
		return true
	}
	return false
}
