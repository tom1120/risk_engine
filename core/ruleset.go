package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
	//	"sort"
	"sync"
)

type RulesetNode struct {
	Name     string   `yaml:"name"`
	Kind     NodeType `yaml:"kind"`
	Tag      string   `yaml:"tag"`
	Label    string   `yaml:"label"`
	ExecPlan string   `yaml:"exec_plan"`
	Depends  []string `yaml:"depends,flow"`
	Rules    []Rule   `yaml:"rules,flow"`
	decision Decision `yaml:"decision"`
}

func (node RulesetNode) GetName() string {
	return node.Name
}

func (node RulesetNode) GetKind() NodeType {
	return node.Kind
}

func (node RulesetNode) GetLabel() string {
	return node.Label
}

func (node RulesetNode) GetTag() string {
	return node.Tag
}

func (ruleset RulesetNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Printf("====[trace]ruleset %s start=====\n", ruleset.Name)

	var ruleResult = make([]Decision, 0)

	//ruleset 批量调用特征
	//depends := ctx.GetFeatures(ruleset.Depends) //global.Features.Get(ruleset.Depends)

	if ruleset.ExecPlan == "parallel" { //并发执行规则
		var wg sync.WaitGroup
		for _, rule := range ruleset.Rules {
			wg.Add(1)
			go func(rule Rule) { //rule
				defer wg.Done()
				rs, err := rule.Parse(ctx)
				if err != nil {
					log.Println(err)
				}

				//ruleDecision := configs.NilDecision todo nil
				if rs.(bool) { //HIT
					//nodeResult.Hits = append(nodeResult.Hits, rule.Name)
					//ruleDecision = configs.DecisionMap[rule.Decision]
					//assign
				}
				ruleResult = append(ruleResult, rule.Decision)
			}(rule)
		}
		wg.Wait()
	} else { //串行执行
		for _, rule := range ruleset.Rules {
			rs, err := rule.Parse(ctx)
			if err != nil {
				return nil, err
			}
			//ruleDecision := configs.NilDecision
			if rs.(bool) { //HIT
				//append hit rule
				//		nodeResult.Hits = append(nodeResult.Hits, rule.Name)
				//ruleDecision = configs.DecisionMap[rule.Decision]
			}
			ruleResult = append(ruleResult, rule.Decision)
		}
	}

	if len(ruleResult) == 0 {
		log.Printf("ruleset %s parse no result\n", ruleset.Name)
		return nil, errcode.ParseErrorRulesetOutputEmpty
	}

	//get max value result, reject is 100, record is 1, pass or no result is 0
	/*	sort.Sort(sort.Reverse(sort.IntSlice(ruleResult)))
		log.Printf("ruleset %s result is :%v\n", ruleset.Name, ruleResult[0])
		nodeResult.Decision = ruleResult[0]
	*/

	//result.AddDetail(*nodeResult)

	log.Printf("====[trace]ruleset %s end=====\n", ruleset.Name)
	return ruleResult[0], nil
}
