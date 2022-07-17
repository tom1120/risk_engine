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
