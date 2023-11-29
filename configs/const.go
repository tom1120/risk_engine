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
package configs

//all operators
const (
	GT         = "GT"
	LT         = "LT"
	GE         = "GE"
	LE         = "LE"
	EQ         = "EQ"
	NEQ        = "NEQ"
	BETWEEN    = "BETWEEN"
	LIKE       = "LIKE"
	IN         = "IN"
	CONTAIN    = "CONTAIN"
	BEFORE     = "BEFORE"
	AFTER      = "AFTER"
	KEYEXIST   = "KEYEXIST"
	VALUEEXIST = "VALUEEXIST"
	AND        = "and"
	OR         = "or"
)

var OperatorMap = map[string]string{
	GT:         ">",
	LT:         "<",
	GE:         ">=",
	LE:         "<=",
	EQ:         "==",
	NEQ:        "!=",
	BETWEEN:    "between",
	LIKE:       "like",
	IN:         "in",
	CONTAIN:    "contain",
	BEFORE:     "before",
	AFTER:      "after",
	KEYEXIST:   "keyexist",
	VALUEEXIST: "valueexist",
	AND:        "&&",
	OR:         "||",
}

var NumSupportOperator = map[string]struct{}{
	GT:      struct{}{},
	LT:      struct{}{},
	GE:      struct{}{},
	LE:      struct{}{},
	EQ:      struct{}{},
	NEQ:     struct{}{},
	BETWEEN: struct{}{},
	IN:      struct{}{},
}
var StringSupportOperator = map[string]struct{}{
	EQ:   struct{}{},
	NEQ:  struct{}{},
	LIKE: struct{}{},
	IN:   struct{}{},
}
var EnumSupportOperator = map[string]struct{}{
	EQ:  struct{}{},
	NEQ: struct{}{},
}
var BoolSupportOperator = map[string]struct{}{
	EQ:  struct{}{},
	NEQ: struct{}{},
}
var DateSupportOperator = map[string]struct{}{
	BEFORE:  struct{}{},
	AFTER:   struct{}{},
	EQ:      struct{}{},
	NEQ:     struct{}{},
	BETWEEN: struct{}{},
}
var ArraySupportOperator = map[string]struct{}{
	EQ:      struct{}{},
	NEQ:     struct{}{},
	CONTAIN: struct{}{},
	IN:      struct{}{},
}
var MapSupportOperator = map[string]struct{}{
	KEYEXIST:   struct{}{},
	VALUEEXIST: struct{}{},
}
var DefaultSupportOperator = map[string]struct{}{
	EQ:  struct{}{},
	NEQ: struct{}{},
}

var CompareOperators = map[string]struct{}{
	EQ:  struct{}{},
	NEQ: struct{}{},
	GT:  struct{}{},
	LT:  struct{}{},
	GE:  struct{}{},
	LE:  struct{}{},
}

var BooleanOperators = map[string]string{
	AND: OperatorMap[AND],
	OR:  OperatorMap[OR],
}

//all support node
const (
	START          = "start"
	END            = "end"
	RULESET        = "ruleset"
	ABTEST         = "abtest"
	CONDITIONAL    = "conditional"
	DECISIONTREE   = "decisiontree"
	DECISIONMATRIX = "decisionmatrix"
	SCORECARD      = "scorecard"
)

//matrix
const (
	MATRIXX = "matrixX"
	MATRIXY = "matrixY"
)

//all type
const (
	INT     = "int"
	FLOAT   = "float"
	STRING  = "string"
	BOOL    = "bool"
	DATE    = "date"
	ARRAY   = "array"
	MAP     = "map"
	DEFAULT = "default"
)

//date type
const (
	DATE_FORMAT        = "2006-01-02"
	DATE_FORMAT_DETAIL = "2006-01-02 15:04:05"
)
