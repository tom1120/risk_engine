package core

import (
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/operator"
	"log"
)

type ConditionalNode struct {
	Name    string   `yaml:"name"`
	Kind    NodeType `yaml:"kind"`
	Label   string   `yaml:"label"`
	Tag     string   `yaml:"tag"`
	Depends []string `yaml:"depends"`
	Branchs []Branch `yaml:"branchs,flow"`
}

func (node ConditionalNode) GetName() string {
	return node.Name
}

func (node ConditionalNode) GetKind() NodeType {
	return node.Kind
}

func (node ConditionalNode) GetLabel() string {
	return node.Label
}

func (node ConditionalNode) GetTag() string {
	return node.Tag
}

func (node ConditionalNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (node ConditionalNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (conditional ConditionalNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Println("====[trace]conditional start=====", conditional.Name)

	for _, branch := range conditional.Branchs { //loop all the branch
		var conditionRs = make([]bool, 0)
		for _, condition := range branch.Conditions {
			if feature, ok := ctx.GetFeature(condition.Feature); ok {
				value, _ := feature.GetValue()
				rs, err := operator.Compare(condition.Operator, value, condition.Value)
				if err != nil {
					return nil, err
				}
				conditionRs = append(conditionRs, rs)
			} else { //get feature fail
				continue //can modify according scene
			}
		}
		logicRs, _ := operator.Boolean(conditionRs, branch.Logic)
		if logicRs { //if true, choose the branch and break
			//nodeResult.SetDecision(branch.Decision)
			//result.AddDetail(*nodeResult)
			return branch.Decision.Output, nil
		} else {
			continue
		}
	}

	log.Println("====[trace]conditional end=====", conditional.Name)
	return nil, errcode.ParseErrorNoBranchMatch //can't find any branch
}
