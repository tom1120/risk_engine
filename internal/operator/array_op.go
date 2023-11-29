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

// jundge val in arr
func InArray(arr []interface{}, val interface{}) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if ok, err := Compare("EQ", v, val); err == nil && ok {
			return true
		}
	}
	return false
}

// jundge array A in Array B
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
