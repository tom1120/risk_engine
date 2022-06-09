package configs

//策略
type Strategy struct {
	Name     string `yaml:"name"`
	Priority int    `yaml:"priority"` //越大越优先
	Score    int    `yaml:"score"`    //策略分
}

var Strategys = map[string]Strategy{
	"reject":  {"reject", 9, 100},
	"approve": {"approve", 5, 5},
	"record":  {"record", 1, 1},
}

//阻断策略
var BlockStrategy = []string{"reject"}
