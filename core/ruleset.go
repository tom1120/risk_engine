package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
	//	"sort"
	"sync"
)

type RulesetNode struct {
	Info          NodeInfo      `yaml:"info"`
	ExecPlan      string        `yaml:"exec_plan"`
	BlockStrategy BlockStrategy `yaml:"block_strategy"`
	Rules         []Rule        `yaml:"rules,flow"`
}

func (rulesetNode RulesetNode) GetName() string {
	return rulesetNode.Info.Name
}

func (rulesetNode RulesetNode) GetType() NodeType {
	return GetNodeType(rulesetNode.Info.Kind)
}

func (rulesetNode RulesetNode) GetInfo() NodeInfo {
	return rulesetNode.Info
}

func (rulesetNode RulesetNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	log.Printf("====[trace]ruleset %s start=====\n", rulesetNode.GetName())

	var ruleResult = make([]*Output, 0)

	//ruleset 批量调用特征
	//depends := ctx.GetFeatures(ruleset.Depends) //global.Features.Get(ruleset.Depends)

	if rulesetNode.ExecPlan == "parallel" { //并发执行规则
		var wg sync.WaitGroup
		var mu sync.Mutex
		for _, rule := range rulesetNode.Rules {
			wg.Add(1)
			go func(rule Rule) { //rule
				defer wg.Done()

				output, err := rule.Parse(ctx)
				if err != nil { //todo 报错如何处理
					log.Println(err)
				}
				if output == (*Output)(nil) {
					return

				}

				//命中规则有结果
				//加入规则命中列表中，ctx.AddHitRule(rule) rule id,name,tag,label,feature
				log.Println("命中规则")
				mu.Lock() //使用channel取代锁
				ruleResult = append(ruleResult, output)
				mu.Unlock()
			}(rule)
		}
		wg.Wait()
	} else { //串行执行
		for _, rule := range rulesetNode.Rules {
			output, err := rule.Parse(ctx)
			if err != nil {
				return nil, err //todo报错如何处理
			}
			if output == (*Output)(nil) {
				continue
			}
			//命中规则有结果
			//加入规则命中列表中，ctx.AddHitRule(rule) rule id,name,tag,label,feature
			log.Println("命中规则")
			ruleResult = append(ruleResult, output)
		}
	}

	//无规则命中
	if len(ruleResult) == 0 {
		log.Printf("ruleset %s parse no result\n", rulesetNode.GetName())
		return (*NodeResult)(nil), errcode.ParseErrorRulesetOutputEmpty
	}

	//TypeStrategy
	//命中阻断规则

	//规则得分

	//最高优先级的规则
	for _, output := range ruleResult {
		log.Println(output)
	}
	//	log.Println(ruleResult)

	//get max value result, reject is 100, record is 1, pass or no result is 0
	/*	sort.Sort(sort.Reverse(sort.IntSlice(ruleResult)))
		log.Printf("ruleset %s result is :%v\n", ruleset.Name, ruleResult[0])
		nodeResult.Decision = ruleResult[0]
	*/

	//result.AddDetail(*nodeResult)

	log.Printf("====[trace]ruleset %s end=====\n", rulesetNode.GetName())
	return (*NodeResult)(nil), nil
}
