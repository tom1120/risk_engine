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
package util

import (
	"time"
)

func TimeSince(from time.Time) int64 {
	return int64(time.Since(from)) / 1e6
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
