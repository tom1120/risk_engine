package core

import (
	"log"
	"testing"
)

func TestFlow(t *testing.T) {
	//flow := generateFlow("../test/yaml/flow_simple.yaml")
	flow := generateFlow("../demo/flow_abtest.yaml")

	log.Println("=========all node========")
	a := flow.GetAllNodes()
	for k, v := range a {
		log.Println(k, v)
	}

	log.Println("--------start run----------")
	ctx := NewPipelineContext()
	//ctx初始化时候将所有特征都new加载了, 然后再每一部set 值
	fMap := map[string]interface{}{"feature_1": 60, "feature_2": 5, "feature_3": 80, "feature_4": 1, "feature_5": 2, "feature_6": 8}
	features := make(map[string]*Feature)
	for k, v := range fMap {
		feature := NewFeature(k, TypeInt, -9999)
		feature.SetValue(v)
		features[k] = feature
	}

	ctx.SetFeatures(features)
	flow.Run(ctx)

	for _, track := range ctx.GetTracks() {
		log.Println(track)
	}
}
