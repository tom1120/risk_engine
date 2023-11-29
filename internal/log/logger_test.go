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
package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	//InitLogger("console", "")
	InitLogger("file", "./out")
	Errorf("this is error %s", "aa")
	Error("this is error!")
	Debug("this is debug!")
	Warn("this is warn!")
	Info("this is info!")
}
