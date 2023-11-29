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
package udf

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Log(Sum(100))
	t.Log(Sum(100.5))
	t.Log(Sum([]interface{}{3.3, 100, 5.0, 8}))
	t.Log(Sum([]interface{}{3.3, 100, 5.0, 8, "88"}))
	t.Log(Sum([]interface{}{3.3, 100, 5.0, 8, true}))
	t.Log(Sum([]interface{}{3.3, 100, 5.0, 8, "8ab"}))
	t.Log(Sum([3]int{3, 4, 5}))
	t.Log(Sum([3]float32{3.0, 4.1, 5}))
	t.Log(Sum([3]string{"3.0", "4.1", "5"}))
	t.Log(Sum([3]string{"3.a", "b.1", "5"}))
}
