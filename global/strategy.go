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
	"github.com/skyhackvip/risk_engine/configs"
)

//from configs
var Strategys = map[string]configs.Strategy{
	"reject":  {"reject", 9, 100},
	"approve": {"approve", 5, 5},
	"record":  {"record", 1, 1},
}
