package core

import (
	"log"
	"testing"
)

func TestWorkflow(t *testing.T) {

	flow := NewFlow()
	//dsl := LoadDslFromFile("../test/yaml/flow_simple1.yaml")
	dsl := LoadDslFromFile("../test/yaml/flow_abtest.yaml")

	//map
	rulesetMap := make(map[string]INode)
	for _, ruleset := range dsl.Rulesets {
		rulesetMap[ruleset.Name] = ruleset
	}

	abtestMap := make(map[string]INode)
	for _, abtest := range dsl.Abtests {
		log.Println("=====")
		log.Println(abtest)
		log.Println("=====")
		abtestMap[abtest.Name] = abtest
	}

	//flow
	for _, flowNode := range dsl.Workflow {
		newNode := flowNode
		log.Println(flowNode)
		switch flowNode.NodeType {
		case TypeRuleset:
			//node := rulesetMap[flowNode.NodeName]
			//flowNode.SetElem(node)
			newNode.SetElem(rulesetMap[flowNode.NodeName])
			flow.AddNode(&newNode)
			//next
			//t.Log("ruleset", flowNode.NodeName, flowNode.Category)
		case TypeAbtest:
			newNode.SetElem(abtestMap[flowNode.NodeName])
			flow.AddNode(&newNode)
		case TypeStart:
			//startNode := flowNode //need set new variable
			newNode.SetElem(NewStartNode(newNode.NodeName))
			flow.SetStartNode(&newNode)
			flow.AddNode(&newNode)
			//t.Log("start", flowNode.NodeName, flowNode.Category)
		case TypeEnd:
			newNode.SetElem(NewEndNode(newNode.NodeName))
			flow.AddNode(&newNode)
		default:
			log.Println("unkown node type")
		}
	}

	a := flow.GetAllNodes()
	log.Println(a)
	for k, v := range a {
		t.Log(k, v)
	}

	t.Log("--------start run----------")
	ctx := NewPipelineContext()
	features := map[string]interface{}{"feature_1": 60, "feature_2": 5, "feature_3": 80, "feature_4": 1, "feature_5": 2, "feature_6": 8}

	for k, v := range features {
		ctx.SetFeature(k, Feature{Name: k, Type: TypeInt, Value: v})
	}

	flow.Run(ctx)

	/*
				branch := []string{"rule1", "rule2"}
		abtestNode := NewFlowNode(NewAbtestNode("ab1", branch))
		flow.AddNode(abtestNode)

		startNode.SetNextNode(abtestNode)

		//start->ab->1->3->end
		//         ->2->end

		flowNode1.SetNextNode(flowNode3)

		flowNode3.SetNextNode(endNode)
		flowNode2.SetNextNode(endNode)

		//具体case
		features := map[string]interface{}{"a": 3, "b": "b"}
		ctx.SetFeatures(features)

		flow.Run()

		f := ctx.GetFeatures()
		for k, v := range f {
			t.Log(k, v)
		}
	*/
}
