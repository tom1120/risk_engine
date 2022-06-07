package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"log"
)

type StartNode struct {
	Name string
	Kind NodeType
}

func NewStartNode(name string) *StartNode {
	return &StartNode{
		Name: name,
		Kind: TypeStart,
	}
}

func (node *StartNode) GetName() string {
	return node.Name
}

func (node *StartNode) GetKind() NodeType {
	return node.Kind
}

func (node *StartNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Println("====[trace]start=====", node.Name)
	return nil, nil
}
