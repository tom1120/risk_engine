package core

import (
	"github.com/skyhackvip/risk_engine/internal/util"
	"log"
)

type ScorecardNode struct {
	Info     NodeInfo `yaml:"info"`
	Blocks   []Block  `yaml:"blocks,flow"`
	Strategy Strategy `yaml:"strategy"`
}

func (scorecardNode ScorecardNode) GetName() string {
	return scorecardNode.Info.Name
}

func (scorecardNode ScorecardNode) GetType() NodeType {
	return GetNodeType(scorecardNode.Info.Kind)
}

func (scorecardNode ScorecardNode) GetInfo() NodeInfo {
	return scorecardNode.Info
}

func (scorecardNode ScorecardNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (scorecardNode ScorecardNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (scorecardNode ScorecardNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	info := scorecardNode.GetInfo()
	log.Printf("======[trace]Scorecard(%s, %s) start======\n", info.Label, scorecardNode.GetName())
	depends := ctx.GetFeatures(info.Depends)
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: scorecardNode.GetType(), Tag: info.Tag, Label: info.Label}
	retArr := make([]interface{}, 0)
	for _, block := range scorecardNode.Blocks {
		ret, _, err := block.parse(depends)
		if err != nil {
			log.Println(err)
			break
		}
		retArr = append(retArr, ret)
	}
	//total score
	var score float64 = 0
	if len(retArr) > 0 && scorecardNode.Strategy.Logic == "sum" {
		for _, ret := range retArr {
			curScore, err := util.ToFloat64(ret)
			if err != nil {
				log.Println(err)
				break
			}
			score += curScore
		}
	}
	//save into ctx feature
	kind := scorecardNode.Strategy.OutputKind
	feature := NewFeature(scorecardNode.GetName(), GetFeatureType(kind))
	feature.SetValue(score)
	ctx.SetFeature(feature)
	if scorecardNode.Strategy.OutputName != "" { //extra
		extraFeature := NewFeature(scorecardNode.Strategy.OutputName, GetFeatureType(kind))
		extraFeature.SetValue(score)
		ctx.SetFeature(extraFeature)
	}

	//output
	nodeResult.Value = score
	log.Printf("======[trace]Scorecard(%s, %s) end======\n", info.Label, scorecardNode.GetName())
	return nodeResult, nil
}
