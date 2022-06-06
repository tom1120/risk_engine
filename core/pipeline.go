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

	//track

	mutex sync.RWMutex

	features map[string]Feature
}

func NewPipelineContext() *PipelineContext {
	return &PipelineContext{features: make(map[string]Feature)}
}

func (ctx *PipelineContext) SetFeatures(features map[string]Feature) {
	if len(features) == 0 {
		return
	}
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	for k, v := range features {
		ctx.features[k] = v //override the same key feature
	}
}

func (ctx *PipelineContext) SetFeature(name string, value Feature) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	ctx.features[name] = value //override the same key feature
}

func (ctx *PipelineContext) GetFeature(name string) (result Feature, ok bool) {
	//local
	ctx.mutex.RLock()
	localFeatures := ctx.features
	ctx.mutex.RUnlock()
	if result, ok = localFeatures[name]; ok {
		return
	}
	//from remote

	return
}

func (ctx *PipelineContext) GetFeatures(depends []string) (result map[string]Feature) {
	if len(depends) == 0 {
		return
	}

	//from local
	ctx.mutex.RLock()
	localFeatures := ctx.features
	ctx.mutex.RUnlock()

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

func (ctx *PipelineContext) GetAllFeatures() map[string]Feature {
	ctx.mutex.RLock()
	features := ctx.features
	ctx.mutex.RUnlock()
	return features
}
