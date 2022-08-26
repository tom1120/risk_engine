package operator

import (
	"errors"
	"github.com/Knetic/govaluate"
	"log"
)

func Evaluate(exprStr string, params map[string]interface{}) (bool, error) {
	expr, err := govaluate.NewEvaluableExpression(exprStr)
	log.Println("base evaluate:", expr, params)
	if err != nil {
		return false, err
	}
	eval, err := expr.Evaluate(params)
	if err != nil {
		return false, err
	}
	if result, ok := eval.(bool); ok {
		return result, nil
	}
	return false, errors.New("convert error")
}

func InArray(arr []interface{}, val interface{}) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func AInB(a []interface{}, b []interface{}) bool {
	if len(b) == 0 {
		return false
	}
	if len(a) == 0 {
		return true
	}
	if len(a) > len(b) {
		return false
	}
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
