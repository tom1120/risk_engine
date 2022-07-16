package configs

var LogicMap = map[string]string{
	"OR":  "||",
	"AND": "&&",
}

var OperatorMap = map[string]string{
	"GT":      ">",
	"LT":      "<",
	"GE":      ">=",
	"LE":      "<=",
	"EQ":      "==",
	"NEQ":     "!=",
	"BETWEEN": "between",
	"LIKE":    "like",
	"IN":      "in",
	"CONTAIN": "contain",
	//todo add string like
}

var NumSupportOperator = map[string]struct{}{
	"GT":      struct{}{},
	"LT":      struct{}{},
	"GE":      struct{}{},
	"LE":      struct{}{},
	"EQ":      struct{}{},
	"NEQ":     struct{}{},
	"BETWEEN": struct{}{},
}
var StringSupportOperator = map[string]struct{}{
	"EQ":      struct{}{},
	"NEQ":     struct{}{},
	"LIKE":    struct{}{},
	"IN":      struct{}{},
	"CONTAIN": struct{}{},
}
var DefaultSupportOperator = map[string]struct{}{
	"EQ":  struct{}{},
	"NEQ": struct{}{},
}

//todo
var DecisionMap = map[string]int{
	"reject": 100, //first priority
	"pass":   0,
	"record": 1,
}

const (
	ScoreReplace = "((score))"
)

const (
	Sum = "SUM"
	Min = "MIN"
	Max = "MAX"
	Avg = "AVG"
)

//decision
const (
	NilDecision   = 0        //not hit rules strategy
	BreakDecision = "reject" //if hit,break at once
)

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
