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
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/log"
	"math/rand"
	"time"
)

type AbtestNode struct {
	Info    NodeInfo `yaml:"info"`
	Branchs []Branch `yaml:"branchs,flow"`
}

func (abtest AbtestNode) GetName() string {
	return abtest.Info.Name
}

func (abtest AbtestNode) GetType() NodeType {
	return GetNodeType(abtest.Info.Kind)
}

func (abtest AbtestNode) GetInfo() NodeInfo {
	return abtest.Info
}

func (node AbtestNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (node AbtestNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (abtest AbtestNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	info := abtest.GetInfo()
	log.Infof("======[trace] abtest %s start======", info.Label, abtest.GetName())
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: abtest.GetType(), Tag: info.Tag, Label: info.Label, IsBlock: false}

	rand.Seed(time.Now().UnixNano())
	winNum := rand.Float64() * 100
	var counter float64 = 0
	var matchBranch bool
	for _, branch := range abtest.Branchs {
		counter += branch.Percent
		if counter > winNum {
			log.Infof("abtest name %s, branch %s, randomNum %v, output %v", abtest.GetName(), branch.Name, winNum, branch.Decision.Output)
			nodeResult.NextNodeName = branch.Decision.Output.Value.(string)
			nodeResult.NextNodeType = GetNodeType(branch.Decision.Output.Kind)
			matchBranch = true
			break //break loop
		}
	}
	nodeResult.Value = winNum

	log.Infof("======[trace] abtest %s end======", info.Label, abtest.GetName())
	if matchBranch {
		return nodeResult, nil
	}
	return nodeResult, errcode.ParseErrorNoBranchMatch
}
