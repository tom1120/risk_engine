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
package global

import (
	"github.com/skyhackvip/risk_engine/internal/udf"
	"sync"
)

var udfMap map[string]udf.Udf = make(map[string]udf.Udf)
var mu sync.RWMutex

func RegisterUdf(key string, fn udf.Udf) {
	mu.Lock()
	defer mu.Unlock()
	udfMap[key] = fn //override if exists
}

func GetUdf(key string) udf.Udf {
	mu.RLock()
	defer mu.RUnlock()
	return udfMap[key]
}
