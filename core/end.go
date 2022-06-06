package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"log"
)

type EndNode struct {
	Name string
	Type NodeType
}

func NewEndNode(name string) *EndNode {
	return &EndNode{
		Name: name,
		Type: TypeEnd,
	}
}

func (node *EndNode) GetName() string {
	return node.Name
}

func (node *EndNode) GetType() NodeType {
	return node.Type
}

func (node *EndNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Println("End parse" + node.Name)
	return nil, nil
}
