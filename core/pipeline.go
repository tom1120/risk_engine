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

	fMutex   sync.RWMutex
	features map[string]*Feature
}

type Track struct {
	Id    int64
	Name  string
	Label string
	Tag   string
	Kind  NodeType
}

func NewPipelineContext() *PipelineContext {
	return &PipelineContext{features: make(map[string]*Feature), hitRules: make(map[string]*Rule)}
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

func (ctx *PipelineContext) SetFeatures(features map[string]*Feature) {
	if len(features) == 0 {
		return
	}
	ctx.fMutex.Lock()
	defer ctx.fMutex.Unlock()
	for k, v := range features {
		ctx.features[k] = v //override the same key feature
	}
}

func (ctx *PipelineContext) SetFeature(feature *Feature) {
	ctx.fMutex.Lock()
	defer ctx.fMutex.Unlock()
	ctx.features[feature.GetName()] = feature //override the same key feature
}

func (ctx *PipelineContext) GetFeature(name string) (result *Feature, ok bool) {
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

func (ctx *PipelineContext) GetFeatures(depends []string) (result map[string]*Feature) {
	if len(depends) == 0 {
		return
	}

	//from local
	ctx.fMutex.RLock()
	localFeatures := ctx.features
	ctx.fMutex.RUnlock()

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
		return
	}

	//curl
	return
}

func (ctx *PipelineContext) GetAllFeatures() map[string]*Feature {
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

type DecisionResult struct {
	HitRules map[string]*Rule
	Tracks   []*Track
	Features map[string]*Feature
}

func (ctx *PipelineContext) GetDecisionResult() *DecisionResult {
	return &DecisionResult{
		HitRules: ctx.GetHitRules(),
		Tracks:   ctx.GetTracks(),
		Features: ctx.GetAllFeatures(),
	}
}
