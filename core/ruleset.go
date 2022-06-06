package core

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
	"sort"
)

type RulesetNode struct {
	Name     string   `yaml:"ruleset_name"`
	Type     NodeType `yaml:"node_type"`
	Category string   `yaml:"ruleset_category"`
	RuleExec string   `yaml:"rule_exec"`
	Rules    []Rule   `yaml:"rules,flow"`
	Depends  []string `yaml:"depends,flow"`
}

func NewRulesetNode(name string) RulesetNode {
	return RulesetNode{
		Name: name,
		Type: TypeRuleset,
	}
}

func (node RulesetNode) GetName() string {
	return node.Name
}

func (node RulesetNode) GetType() NodeType {
	return node.Type
}

func (ruleset RulesetNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Printf("====trace : ruleset %s start=====\n", ruleset.Name)

	nodeResult := dto.NewNodeResult(ruleset.Name)

	var ruleResult = make([]int, 0)
	//depends := ctx.GetFeatures(ruleset.Depends) //global.Features.Get(ruleset.Depends)

	//nodeResult.AddFactor(depends)
	for _, rule := range ruleset.Rules {
		rs, err := rule.Parse(ctx)
		if err != nil {
			return nil, err
		}
		ruleDecision := configs.NilDecision
		if rs.(bool) { //HIT
			nodeResult.Hits = append(nodeResult.Hits, rule.RuleName)
			ruleDecision = configs.DecisionMap[rule.Decision]
		}
		ruleResult = append(ruleResult, ruleDecision)
	}

	if len(ruleResult) == 0 {
		log.Printf("ruleset %s parse no result\n", ruleset.Name)
		return nil, errcode.ParseErrorRulesetOutputEmpty
	}

	//get max value result, reject is 100, record is 1, pass or no result is 0
	sort.Sort(sort.Reverse(sort.IntSlice(ruleResult)))
	log.Printf("ruleset %s result is :%v\n", ruleset.Name, ruleResult[0])
	nodeResult.Decision = ruleResult[0]

	//result.AddDetail(*nodeResult)

	log.Printf("====trace : ruleset %s end=====\n", ruleset.Name)
	return ruleResult[0], nil
}
