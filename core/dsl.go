package core

import (
	"log"
)

type Dsl struct {
	Key          string                 `yaml:"key"`
	Version      string                 `yaml:"version"`
	Metadata     map[string]interface{} `yaml:"metadata"`
	DecisionFlow []FlowNode             `yaml:"decision_flow,flow"`
	Rulesets     []RulesetNode          `yaml:"rulesets,flow"`
	Abtests      []AbtestNode           `yaml:"abtests,flow"`
	//	Conditionals    []Conditional    `yaml:"conditionals,flow"`
	//	DecisionTrees   []DecisionTree   `yaml:"decisiontrees,flow"`
	//	DecisionMatrixs []DecisionMatrix `yaml:"decisionmatrixs,flow"`
	//	ScoreCards      []ScoreCard      `yaml:"scorecards,flow"`
}

func (dsl *Dsl) CheckValid() bool {
	if dsl.Key == "" {
		return false
	}
	if len(dsl.DecisionFlow) == 0 {
		return false
	}
	return true
}

//dsl to decisionflow
func (dsl *Dsl) ConvertToDecisionFlow() (*DecisionFlow, error) {
	flow := NewDecisionFlow()
	flow.Key = dsl.Key
	flow.Version = dsl.Version
	flow.Metadata = dsl.Metadata

	//map
	rulesetMap := make(map[string]INode)
	for _, ruleset := range dsl.Rulesets {
		rulesetMap[ruleset.GetName()] = ruleset
	}
	abtestMap := make(map[string]INode)
	for _, abtest := range dsl.Abtests {
		abtestMap[abtest.GetName()] = abtest
	}

	//flow
	for _, flowNode := range dsl.DecisionFlow {
		newNode := flowNode //need set new variable
		switch GetNodeType(newNode.NodeKind) {
		case TypeRuleset:
			newNode.SetElem(rulesetMap[newNode.NodeName])
			flow.AddNode(&newNode)
		case TypeAbtest:
			newNode.SetElem(abtestMap[newNode.NodeName])
			flow.AddNode(&newNode)
		case TypeStart:
			newNode.SetElem(NewStartNode(newNode.NodeName))
			flow.SetStartNode(&newNode)
			flow.AddNode(&newNode)
		case TypeEnd:
			newNode.SetElem(NewEndNode(newNode.NodeName))
			flow.AddNode(&newNode)
		default:
			log.Printf("dsl (%s-%s) convert warning: unkown node type (%s)\n", dsl.Key, dsl.Version, newNode.NodeKind)
		}
	}
	return flow, nil
}
