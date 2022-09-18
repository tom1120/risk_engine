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
