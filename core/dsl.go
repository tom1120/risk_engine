package core

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
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

//load dsl from file
func LoadDslFromFile(file string) (*Dsl, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return (*Dsl)(nil), nil
	}
	return LoadDsl(yamlFile)
}

func LoadDsl(file []byte) (*Dsl, error) {
	dsl := new(Dsl)
	err := yaml.Unmarshal(file, dsl)
	return dsl, err
}

func (dsl *Dsl) ConvertToDecisionFlow() *DecisionFlow {
	flow := NewDecisionFlow()

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
		switch GetNodeType(flowNode.NodeKind) {
		case TypeRuleset:
			newNode.SetElem(rulesetMap[flowNode.NodeName])
			flow.AddNode(&newNode)
		case TypeAbtest:
			newNode.SetElem(abtestMap[flowNode.NodeName])
			flow.AddNode(&newNode)
		case TypeStart:
			newNode.SetElem(NewStartNode(newNode.NodeName))
			flow.SetStartNode(&newNode)
			flow.AddNode(&newNode)
		case TypeEnd:
			newNode.SetElem(NewEndNode(newNode.NodeName))
			flow.AddNode(&newNode)
		default:
			log.Println("unkown node type")
		}
	}
	return flow
}
