package core

import (
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/log"
)

type TreeNode struct {
	Info     NodeInfo `yaml:"info"`
	Blocks   []Block  `yaml:"blocks,flow"`
	Strategy Strategy `yaml:"strategy"`
}

type Strategy struct {
	OutputName string `yaml:"output_name"`
	OutputKind string `yaml:"output_kind"`
	Start      string `yaml:"start"`
	Logic      string `yaml:"logic"`
}

func (treeNode TreeNode) GetName() string {
	return treeNode.Info.Name
}

func (treeNode TreeNode) GetType() NodeType {
	return GetNodeType(treeNode.Info.Kind)
}

func (treeNode TreeNode) GetInfo() NodeInfo {
	return treeNode.Info
}

func (treeNode TreeNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (treeNode TreeNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (treeNode TreeNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	info := treeNode.GetInfo()
	log.Infof("======[trace] Tree %s start======", info.Label, treeNode.GetName())
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: treeNode.GetType(), Tag: info.Tag, Label: info.Label}
	var resultErr error = nil
	var result interface{}
	blockMap := treeNode.init()

	depends := ctx.GetFeatures(info.Depends)
	block, gotoNext := blockMap[treeNode.Strategy.Start]
	for gotoNext {
		ret, gotoNext, err := block.parse(depends)
		if err != nil {
			log.Error(err)
			resultErr = err
			break
		}
		if gotoNext {
			if b, ok := blockMap[ret.(string)]; !ok {
				resultErr = errcode.ParseErrorTreeNotMatch
			} else {
				block = b
			}
		} else { //finish
			result = ret
			break
		}
	}

	//save into ctx feature
	kind := treeNode.Strategy.OutputKind
	feature := NewFeature(treeNode.GetName(), GetFeatureType(kind))
	feature.SetValue(result)
	ctx.SetFeature(feature)
	if treeNode.Strategy.OutputName != "" { //extra
		extraFeature := NewFeature(treeNode.Strategy.OutputName, GetFeatureType(kind))
		extraFeature.SetValue(result)
		ctx.SetFeature(extraFeature)
	}

	//output
	nodeResult.Value = result
	log.Infof("======[trace] Tree %s end======", info.Label, treeNode.GetName())
	return nodeResult, resultErr
}

func (treeNode TreeNode) init() map[string]Block {
	blockMap := make(map[string]Block)
	for _, block := range treeNode.Blocks {
		blockMap[block.Name] = block
	}
	return blockMap
}
