package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"log"
)

type EndNode struct {
	Name  string   `yaml:"name"`
	Kind  NodeType `yaml:"kind"`
	Tag   string   `"yaml:"tag"`
	Label string   `yaml:"label"`
}

func NewEndNode(name string) *EndNode {
	return &EndNode{
		Name: name,
		Kind: TypeEnd,
	}
}

func (node EndNode) GetName() string {
	return node.Name
}

func (node EndNode) GetKind() NodeType {
	return node.Kind
}

func (node EndNode) GetTag() string {
	return node.Tag
}

func (node EndNode) GetLabel() string {
	return node.Label
}

func (node EndNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Println("======[trace]End=====" + node.Name)
	return nil, nil
}
