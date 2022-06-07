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
