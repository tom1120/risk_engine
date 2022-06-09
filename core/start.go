package core

import (
	"log"
)

type StartNode struct {
	Info NodeInfo
}

func NewStartNode(name string) *StartNode {
	return &StartNode{
		Info: NodeInfo{Name: name, Kind: TypeStart.String()},
	}
}

func (node StartNode) GetName() string {
	return node.Info.Name
}

func (node StartNode) GetType() NodeType {
	return GetNodeType(node.Info.Kind)
}

func (node StartNode) GetInfo() NodeInfo {
	return node.Info
}

func (node StartNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	log.Println("====[trace]start=====", node.GetName())
	return (*NodeResult)(nil), nil
}
