package core

import (
	"errors"
	"fmt"
	"github.com/skyhackvip/risk_engine/internal/operator"
	"log"
)

type Rule struct {
	Name       string      `yaml:"name"`
	Tag        string      `yaml:"tag"`
	Label      string      `yaml:"label"`
	Conditions []Condition `yaml:"conditions,flow"`
	Decision   Decision    `yaml:"decision"`
	Depends    []string    `yaml:"depends"`
}

//parse rule
func (rule *Rule) Parse(ctx *PipelineContext) (interface{}, error) {
	//rule.Conditions
	if len(rule.Conditions) == 0 {
		return nil, errors.New(fmt.Sprintf("rule (%s) condition is empty", rule.Name))
	}
	var conditionRet = make(map[string]interface{}, 0)
	for _, condition := range rule.Conditions {
		if data, ok := ctx.GetFeature(condition.Feature); ok {
			if data.Name == "" {
				log.Println("data error : data name is empty")
				continue
			}
			rs, err := operator.Compare(condition.Operator, data.Value, condition.Value)
			if err != nil {
				return nil, err
			}
			conditionRet[condition.Name] = rs
		} else {
			//lack of feature
			log.Printf("error lack of feature: %s\n", condition.Feature)
			continue
		}
	}
	if len(conditionRet) == 0 {
		return nil, errors.New(fmt.Sprintf("rule (%s) condition is empty", rule.Name))
	}

	//rule.Decision
	expr := rule.Decision.Logic
	logicRet, err := operator.Evaluate(expr, conditionRet)
	if err != nil {
		return nil, err
	}
	log.Printf("rule %s (%s) decision is: %v\n", rule.Label, rule.Name, logicRet)
	return logicRet, nil
}

type Condition struct {
	Feature  string      `yaml:"feature"`
	Operator string      `yaml:"operator"`
	Value    interface{} `yaml:"value"`
	Result   string      `yaml:"result"`
	Name     string      `yaml:"name"`
}

type Decision struct {
	Depends []string               `yaml:"depends,flow"`
	Logic   string                 `yaml:"logic"`
	Output  interface{}            `yaml:"output"` //该节点输出值
	Assign  map[string]interface{} `yaml:"assign"` //赋值
}

type Branch struct {
	BranchName string      `yaml:"branch_name"`
	Conditions []Condition `yaml:"conditions"` //used by conditional
	Logic      string      `yaml:"logic"`      //used by conditional
	Percent    float64     `yaml:"percent"`    //used by abtest
	Decision   Decision    `yaml:"decision"`
}
