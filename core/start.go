package core

import (
	"github.com/skyhackvip/risk_engine/internal/log"
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

func (node StartNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (node StartNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (node StartNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	log.Info("======[trace] Start======")
	info := node.GetInfo()
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: node.GetType(), Tag: info.Tag, Label: info.Label}
	return nodeResult, nil
}
