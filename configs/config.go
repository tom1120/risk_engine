package configs

//策略
type Strategy struct {
	Name     string `yaml:"name"`
	Priority int    `yaml:"priority"` //越大越优先
	Score    int    `yaml:"score"`    //策略分
}
