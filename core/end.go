package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"log"
)

type EndNode struct {
	Name string   `yaml:"name"`
	Kind NodeType `yaml:"kind"`
}

func NewEndNode(name string) *EndNode {
	return &EndNode{
		Name: name,
		Kind: TypeEnd,
	}
}

func (node *EndNode) GetName() string {
	return node.Name
}

func (node *EndNode) GetKind() NodeType {
	return node.Kind
}

func (node *EndNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Println("======[trace]End=====" + node.Name)
	return nil, nil
}
