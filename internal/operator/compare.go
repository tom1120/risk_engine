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
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/log"
	"github.com/skyhackvip/risk_engine/internal/util"
)

//compare operator expression
//left [operator] right
//operator: [eq,neq,gt,lt,ge,le]
func Compare(operator string, left interface{}, right interface{}) (bool, error) {
	log.Infof("compare operator: %v", operator, left, right)
	if _, ok := configs.CompareOperators[operator]; !ok {
		log.Errorf("not compare operator: %v", operator)
		return false, errcode.ErrorNotSupportOperator
	}

	switch operator {
	case configs.EQ:
		return equal(left, right)
	case configs.NEQ:
		rs, err := equal(left, right)
		return !rs, err

	//only number can compare(gt,lt,ge,le)
	case configs.GT:
		return numCompare(left, right, operator)
	case configs.LT:
		return numCompare(left, right, operator)
	case configs.GE:
		return numCompare(left, right, operator)
	case configs.LE:
		return numCompare(left, right, operator)
	}
	return false, errcode.ErrorNotSupportOperator
}

//jundge left == right
func equal(left, right interface{}) (bool, error) {
	leftType, err := util.GetType(left)
	if err != nil {
		log.Errorf("left type unknow: %v", left, err)
		return false, err //unknow type
	}
	rightType, err := util.GetType(right)
	if err != nil {
		log.Errorf("right type unknow: %v", right, err)
		return false, err
	}
	if !util.MatchType(leftType, rightType) {
		return false, nil
	}
	if leftType == configs.ARRAY {
		return arrayEqual(left.([]interface{}), right.([]interface{})), nil
	}
	if leftType == configs.MAP {
		return false, errcode.ErrorNotSupportOperator
	}
	if leftType == configs.STRING {
		return left.(string) == right.(string), nil
	}
	if leftType == configs.BOOL {
		leftBool, err := util.ToBool(left)
		if err != nil {
			return false, err
		}
		rightBool, err := util.ToBool(right)
		if err != nil {
			return false, err
		}
		return leftBool == rightBool, nil
	}
	if leftType == configs.INT || leftType == configs.FLOAT {
		leftNum, err := util.ToFloat64(left)
		if err != nil {
			return false, err
		}
		rightNum, err := util.ToFloat64(right)
		if err != nil {
			return false, err
		}
		return numCompare(leftNum, rightNum, configs.EQ)
	}
	if leftType == configs.DATE {
		leftDate, err := util.ToDate(left)
		if err != nil {
			return false, err
		}
		rightDate, err := util.ToDate(right)
		if err != nil {
			return false, err
		}
		return leftDate.Equal(rightDate), nil
	}
	if leftType == configs.DEFAULT {
		return left == right, nil
	}
	return false, errcode.ErrorNotSupportOperator
}

//a == b true
//a != b false
func arrayEqual(a, b []interface{}) bool {
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

//compare number (lt, gt, le, ge, eq, neq)
//only number can compare
func numCompare(left, right interface{}, op string) (bool, error) {
	leftNum, err := util.ToFloat64(left)
	if err != nil {
		return false, errcode.ErrorNotANumber
	}
	rightNum, err := util.ToFloat64(right)
	if err != nil {
		return false, errcode.ErrorNotANumber
	}
	switch op {
	case configs.EQ:
		return leftNum == rightNum, nil
	case configs.NEQ:
		return leftNum != rightNum, nil
	case configs.GT:
		return leftNum > rightNum, nil
	case configs.LT:
		return leftNum < rightNum, nil
	case configs.GE:
		return leftNum >= rightNum, nil
	case configs.LE:
		return leftNum <= rightNum, nil
	default:
		return false, errcode.ErrorNotSupportOperator
	}
}
