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
package errcode

var (
	ErrorFeatureTypeUnknow  = NewError(2000001, "feature type support int,float,bool,string,date,array,map")
	ErrorTypeConvert        = NewError(2000002, "type convert error")
	ErrorNotSupportOperator = NewError(2000003, "not support operator")
	ErrorNotANumber         = NewError(2000004, "not a number")
	ErrorBooleanValEmpty    = NewError(2000005, "boolean operator value is empty")
	ErrorBooleanValLack     = NewError(2000006, "boolean operator value lack")
)
