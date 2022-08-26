package operator

import (
	"errors"
	"fmt"
	"github.com/skyhackvip/risk_engine/configs"
)

//compare expression:left [><=] right
func Compare(operator string, left interface{}, right interface{}) (bool, error) {
	var params = make(map[string]interface{})
	params["left"] = left
	params["right"] = right

	if _, ok := configs.OperatorMap[operator]; !ok {
		return false, errors.New("not support operator")
	}
	expr := fmt.Sprintf("left %s right", configs.OperatorMap[operator])

	return Evaluate(expr, params)
}

func CompareArray(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)]
	tmp := make(map[interface{}]struct{}, len(b))
	for _, v := range b {
		tmp[v] = struct{}{}
	}
	for _, v := range a {
		if _, ok := tmp[v]; !ok {
			return false
		}
	}
	return true
}
