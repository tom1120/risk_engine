package core

import (
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
)

type TreeNode struct {
	Info     NodeInfo `yaml:"info"`
	Blocks   []Block  `yaml:"blocks,flow"`
	Strategy Strategy `yaml:"strategy"`
}

type Block struct {
	Name       string      `yaml:"name"`
	Feature    string      `yaml:"feature"`
	Conditions []Condition `yaml:"conditions,flow"`
}

type Strategy struct {
	OutputName string `yaml:"output_name"`
	OutputKind string `yaml:"output_kind"`
	Start      string `yaml:"start"`
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
	log.Printf("======[trace]Tree(%s, %s) start======\n", info.Label, treeNode.GetName())
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: treeNode.GetType(), Tag: info.Tag, Label: info.Label}
	var resultErr error = nil
	var result interface{}
	blockMap := treeNode.init()

	depends := ctx.GetFeatures(info.Depends)
	block, gotoNext := blockMap[treeNode.Strategy.Start]
	for gotoNext {
		ret, gotoNext, err := treeNode.parseBlock(block, depends)
		if err != nil {
			log.Println(err)
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
	log.Printf("======[trace]Tree(%s, %s) end======\n", info.Label, treeNode.GetName())
	return nodeResult, resultErr
}

func (treeNode TreeNode) parseBlock(block Block, depends map[string]IFeature) (interface{}, bool, error) {
	for _, condition := range block.Conditions {
		if feature, ok := depends[block.Feature]; ok {
			hit, err := feature.Compare(condition.Operator, condition.Value)
			if err != nil {
				log.Println("parse error", err)
				continue
			}
			if hit {
				if condition.Goto != "" {
					return condition.Goto, true, nil
				} else {
					return condition.Result, false, nil
				}
			}
		} else {
			log.Printf("error lack of feature: %s\n", block.Feature)
			continue
		}
	}
	return nil, false, errcode.ParseErrorTreeNotMatch
}

func (treeNode TreeNode) init() map[string]Block {
	log.Println("init", treeNode.Blocks)
	blockMap := make(map[string]Block)
	for _, block := range treeNode.Blocks {
		blockMap[block.Name] = block
	}
	return blockMap
}
