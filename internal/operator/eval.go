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
