package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"log"
)

type EndNode struct {
	Info NodeInfo
}

func NewEndNode(name string) *EndNode {
	return &EndNode{
		Info: NodeInfo{Name: name, Kind: TypeEnd.String()},
	}
}

func (node EndNode) GetName() string {
	return node.Info.Name
}

func (node EndNode) GetType() NodeType {
	return GetNodeType(node.Info.Kind)
}

func (node EndNode) GetInfo() NodeInfo {
	return node.Info
}

func (node EndNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	log.Println("======[trace]End======")
	return (*NodeResult)(nil), nil
}
