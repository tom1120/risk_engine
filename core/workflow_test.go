package core

import (
	"log"
	"testing"
)

func generateFlow(path string) *DecisionFlow {
	//构建flow
	flow := NewFlow()
	dsl := LoadDslFromFile(path)

	//map
	rulesetMap := make(map[string]INode)
	for _, ruleset := range dsl.Rulesets {
		rulesetMap[ruleset.Name] = ruleset
	}

	abtestMap := make(map[string]INode)
	for _, abtest := range dsl.Abtests {
		abtestMap[abtest.Name] = abtest
	}

	//flow
	for _, flowNode := range dsl.DecisionFlow {
		newNode := flowNode //need set new variable
		switch flowNode.NodeType {
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

func TestWorkflow(t *testing.T) {
	//flow := generateFlow("../test/yaml/flow_simple.yaml")
	flow := generateFlow("../test/yaml/flow_abtest.yaml")

	log.Println("=========all node========")
	a := flow.GetAllNodes()
	for k, v := range a {
		log.Println(k, v)
	}

	log.Println("--------start run----------")
	ctx := NewPipelineContext()
	//ctx初始化时候将所有特征都new加载了, 然后再每一部set 值
	fMap := map[string]interface{}{"feature_1": 60, "feature_2": 5, "feature_3": 80, "feature_4": 1, "feature_5": 2, "feature_6": 8}
	features := make(map[string]*Feature)
	for k, v := range fMap {
		feature := NewFeature(k, TypeInt, -9999)
		feature.SetValue(v)
		features[k] = feature
	}

	ctx.SetFeatures(features)
	flow.Run(ctx)

	for _, track := range ctx.GetTracks() {
		log.Println(track)
	}
}
