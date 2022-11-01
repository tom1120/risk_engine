package operator

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/skyhackvip/risk_engine/internal/log"
)

//using govalute to execute expression
func Evaluate(exprStr string, params map[string]interface{}) (bool, error) {
	expr, err := govaluate.NewEvaluableExpression(exprStr)
	log.Infof("base evaluate: %v", expr, params)
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
