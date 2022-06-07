package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
	"math/rand"
	"time"
)

type AbtestNode struct {
	Name       string   `yaml:"name"`
	Label      string   `yaml:"label"`
	Tag        string   `yaml:"tag"`
	Kind       NodeType `yaml:"kind"`
	OutputName string   `yaml:"output_name"`
	Branchs    []Branch `yaml:"branchs,flow"`
}

func (abtest AbtestNode) GetName() string {
	return abtest.Name
}

func (abtest AbtestNode) GetKind() NodeType {
	return abtest.Kind
}

func (abtest AbtestNode) GetLabel() string {
	return abtest.Label
}

func (abtest AbtestNode) GetTag() string {
	return abtest.Tag
}

func (abtest AbtestNode) Parse(ctx *PipelineContext) (interface{}, error) {
	log.Println("====[trace] abtest========")
	rand.Seed(time.Now().UnixNano())
	winNum := rand.Float64() * 100
	var counter float64 = 0
	for _, branch := range abtest.Branchs {
		counter += branch.Percent
		if counter > winNum {
			//feature global.Features.Set(dto.Feature{Name: abtest.Name, Value: branch.BranchName})
			log.Printf("abtest %v : %v, %v\n", abtest.Name, branch.BranchName, winNum)
			if res, ok := branch.Decision.Output.([]interface{}); ok {
				if len(res) == 2 {
					log.Println("abtest result", res)
					return res, nil
				}
			}
		}
	}
	return nil, errcode.ParseErrorNoBranchMatch
}
