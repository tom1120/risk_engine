package core

import (
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/operator"
	"log"
)

type ConditionalNode struct {
	Info    NodeInfo `yaml:"info"`
	Branchs []Branch `yaml:"branchs,flow"`
}

func (node ConditionalNode) GetName() string {
	return node.Info.Name
}

func (node ConditionalNode) GetType() NodeType {
	return GetNodeType(node.Info.Kind)
}

func (node ConditionalNode) GetInfo() NodeInfo {
	return node.Info
}

func (node ConditionalNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (node ConditionalNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (conditional ConditionalNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	info := conditional.GetInfo()
	log.Printf("====[trace]conditional (%s, %s) start=====\n", info.Label, conditional.GetName())
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: conditional.GetType(), Tag: info.Tag, Label: info.Label, IsBlock: false}

	depends := ctx.GetFeatures(info.Depends)
	var matchBranch bool
	for _, branch := range conditional.Branchs { //loop all the branch
		var conditionRet = make(map[string]interface{}, 0)
		for _, condition := range branch.Conditions {
			if feature, ok := depends[condition.Feature]; ok {
				rs, err := feature.Compare(condition.Operator, condition.Value)
				if err != nil {
					return nil, err
				}
				conditionRet[condition.Name] = rs
			} else { //get feature fail
				log.Printf("error lack of feature: %s\n", condition.Feature)
				continue
			}
		}
		if len(conditionRet) == 0 { //current branch not match
			continue
		}
		logicRs, err := operator.Evaluate(branch.Decision.Logic, conditionRet)
		if err != nil {
			continue
		}
		if logicRs { //if true, choose the branch and break
			log.Printf("conditional %v : %v,  output:%v \n", conditional.GetName(), branch.Name, branch.Decision.Output)
			nodeResult.Value = branch.Name
			nodeResult.NextNodeName = branch.Decision.Output.Value.(string)
			nodeResult.NextNodeType = GetNodeType(branch.Decision.Output.Kind)
			matchBranch = true
			break
		}
	}
	log.Printf("====[trace]conditional (%s, %s) end=====\n", info.Label, conditional.GetName())
	if matchBranch {
		return nodeResult, nil
	}
	return nodeResult, errcode.ParseErrorNoBranchMatch //can't find any branch

}
