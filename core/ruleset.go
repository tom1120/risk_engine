package core

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/operator"
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
	info := rulesetNode.GetInfo()
	log.Printf("======[trace]ruleset(%s, %s) start======\n", info.Label, rulesetNode.GetName())
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: rulesetNode.GetType(), Tag: info.Tag, Label: info.Label}

	var ruleOutputs = make(map[string]*Output, 0)
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
					return
				}
				if !output.GetHit() { //未命中
					return

				}

				//命中规则有结果
				ctx.AddHitRule(&rule)
				mu.Lock() //使用channel取代锁
				ruleOutputs[rule.Name] = output
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
			if !output.GetHit() {
				continue
			}
			//命中规则有结果
			ctx.AddHitRule(&rule)
			ruleOutputs[rule.Name] = output
		}
	}

	//无规则命中
	if len(ruleOutputs) == 0 {
		//log.Printf("ruleset %s parse no result\n", rulesetNode.GetName())
		return nodeResult, nil
	}

	hitRules := make(map[string]struct{})
	if len(rulesetNode.BlockStrategy.HitRule) > 0 {
		for _, v := range rulesetNode.BlockStrategy.HitRule {
			hitRules[v] = struct{}{}
		}
	}

	var block bool
	var score int
	var nodeRt configs.Strategy
	for name, output := range ruleOutputs {
		//节点规则得分
		if s, ok := global.Strategys[output.Value.(string)]; ok {
			score += s.Score
			//根据优先级获取结果
			if nodeRt.Priority < s.Priority { //默认0
				nodeRt = s
			}
		}
		//是否允许提前中断
		if rulesetNode.BlockStrategy.IsBlock {
			//命中规则在 ruleset.block_strategy.hit_rule 列表中
			if _, ok := hitRules[name]; ok {
				block = true
			}
		}
	}
	if rulesetNode.BlockStrategy.IsBlock {
		ok, _ := operator.Compare(rulesetNode.BlockStrategy.Operator, nodeRt.Name, rulesetNode.BlockStrategy.Value)
		if ok {
			block = true
		}
	}
	nodeResult.IsBlock = block
	nodeResult.Score = score
	nodeResult.Value = nodeRt.Name
	log.Printf("======[trace]ruleset(%s, %s) end======\n", info.Label, rulesetNode.GetName())
	return nodeResult, nil
}

func in(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
