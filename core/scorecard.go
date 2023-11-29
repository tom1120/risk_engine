// Copyright (c) 2023
//
// @author 贺鹏Kavin
// 微信公众号:技术岁月
// https://github.com/skyhackvip/risk_engine
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
package core

import (
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/log"
	"github.com/skyhackvip/risk_engine/internal/util"
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
	log.Infof("======[trace] Scorecard %s start======", info.Label, scorecardNode.GetName())
	depends := ctx.GetFeatures(info.Depends)
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: scorecardNode.GetType(), Tag: info.Tag, Label: info.Label}
	retArr := make([]interface{}, 0)
	for _, block := range scorecardNode.Blocks {
		ret, _, err := block.parse(depends)
		if err != nil {
			log.Error(err)
			break
		}
		retArr = append(retArr, ret)
	}

	//total score
	var score float64 = 0
	if len(retArr) > 0 {
		Fn := global.GetUdf(scorecardNode.Strategy.Logic)
		if ret, err := Fn(retArr); err != nil {
			log.Error(err)
		} else {
			score, err = util.ToFloat64(ret)
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
	nodeResult.Score = score
	log.Infof("======[trace] Scorecard %s end======", info.Label, scorecardNode.GetName())
	return nodeResult, nil
}
