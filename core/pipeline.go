package core

import (
	"sync"
)

type PipelineContext struct {
	//dsl
	currentNode INode

	gotoNext bool

	nextNodeKey string

	//next node category

	//request params

	//proccess result

	tMutex sync.RWMutex
	tracks []*Track

	fMutex   sync.RWMutex
	features map[string]*Feature
}

type Track struct {
	name  string
	label string
	tag   string
	kind  NodeType
}

func NewPipelineContext() *PipelineContext {
	return &PipelineContext{features: make(map[string]*Feature)}
}

func (ctx *PipelineContext) AddTrack(node INode) {
	ctx.tMutex.Lock()
	defer ctx.tMutex.Unlock()
	ctx.tracks = append(ctx.tracks, &Track{name: node.GetName(),
		label: node.GetLabel(),
		tag:   node.GetTag(),
		kind:  node.GetKind(),
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
	features := ctx.features
	ctx.fMutex.RUnlock()
	return features
}
