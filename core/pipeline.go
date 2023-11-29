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
	"sync"
)

type PipelineContext struct {
	//request params

	//proccess result

	hMutex   sync.RWMutex
	hitRules map[string]*Rule

	tMutex sync.RWMutex
	tracks []*Track

	nMutex      sync.RWMutex
	nodeResults map[string]*NodeResult

	fMutex   sync.RWMutex
	features map[string]IFeature //保存所有上下文中已赋值的特征
}

type Track struct {
	Id    int64
	Name  string
	Label string
	Tag   string
	Kind  NodeType
}

func NewPipelineContext() *PipelineContext {
	return &PipelineContext{features: make(map[string]IFeature),
		hitRules:    make(map[string]*Rule),
		nodeResults: make(map[string]*NodeResult),
	}
}

func (ctx *PipelineContext) AddTrack(node INode) {
	ctx.tMutex.Lock()
	defer ctx.tMutex.Unlock()
	ctx.tracks = append(ctx.tracks, &Track{Name: node.GetName(),
		Id:    node.GetInfo().Id,
		Label: node.GetInfo().Label,
		Tag:   node.GetInfo().Tag,
		Kind:  node.GetType(),
	})
}

func (ctx *PipelineContext) GetTracks() []*Track {
	ctx.tMutex.RLock()
	defer ctx.tMutex.RUnlock()
	return ctx.tracks
}

func (ctx *PipelineContext) SetFeatures(features map[string]IFeature) {
	if len(features) == 0 {
		return
	}
	ctx.fMutex.Lock()
	defer ctx.fMutex.Unlock()
	for k, v := range features {
		ctx.features[k] = v //override the same key feature
	}
}

func (ctx *PipelineContext) SetFeature(feature IFeature) {
	ctx.fMutex.Lock()
	defer ctx.fMutex.Unlock()
	ctx.features[feature.GetName()] = feature //override the same key feature
}

func (ctx *PipelineContext) GetFeature(name string) (result IFeature, ok bool) {
	//local
	ctx.fMutex.RLock()
	localFeatures := ctx.features
	ctx.fMutex.RUnlock()
	if result, ok = localFeatures[name]; ok {
		return
	}
	//from remote

	return
}

func (ctx *PipelineContext) GetFeatures(depends []string) (result map[string]IFeature) {
	if len(depends) == 0 {
		return
	}

	//from local
	ctx.fMutex.RLock()
	localFeatures := ctx.features
	ctx.fMutex.RUnlock()

	result = make(map[string]IFeature)
	remoteList := make([]string, 0)
	for _, name := range depends {
		if v, ok := localFeatures[name]; ok {
			result[name] = v
		} else {
			remoteList = append(remoteList, name)
		}
	}

	//from remote
	if len(remoteList) == 0 {
		//远程调用后new的时候要知道类型
		return
	}

	//curl
	return
}

func (ctx *PipelineContext) GetAllFeatures() map[string]IFeature {
	ctx.fMutex.RLock()
	defer ctx.fMutex.RUnlock()
	return ctx.features
}

func (ctx *PipelineContext) AddHitRule(rule *Rule) {
	ctx.tMutex.Lock()
	defer ctx.tMutex.Unlock()
	ctx.hitRules[rule.Name] = rule
}

func (ctx *PipelineContext) GetHitRules() map[string]*Rule {
	ctx.tMutex.RLock()
	defer ctx.tMutex.RUnlock()
	return ctx.hitRules
}

func (ctx *PipelineContext) AddNodeResult(name string, nodeResult *NodeResult) {
	ctx.nMutex.Lock()
	defer ctx.nMutex.Unlock()
	ctx.nodeResults[name] = nodeResult
}

func (ctx *PipelineContext) GetNodeResults() map[string]*NodeResult {
	ctx.nMutex.RLock()
	defer ctx.nMutex.RUnlock()
	return ctx.nodeResults
}

type DecisionResult struct {
	HitRules    map[string]*Rule
	Tracks      []*Track
	Features    map[string]IFeature
	NodeResults map[string]*NodeResult
}

func (ctx *PipelineContext) GetDecisionResult() *DecisionResult {
	return &DecisionResult{
		HitRules:    ctx.GetHitRules(),
		Tracks:      ctx.GetTracks(),
		Features:    ctx.GetAllFeatures(),
		NodeResults: ctx.GetNodeResults(),
	}
}
