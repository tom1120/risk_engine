package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"log"
)

type StartNode struct {
	Name string
	Type NodeType
}

func NewStartNode(name string) *StartNode {
	return &StartNode{
		Name: name,
		Type: TypeStart,
	}
}

func (node *StartNode) GetName() string {
	return node.Name
}

func (node *StartNode) GetType() NodeType {
	return node.Type
}

func (node *StartNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Println("====trace===start parse=", node.Name)
	return nil, nil
}
