package core

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	//"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
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
	log.Printf("======[trace]abtest(%s, %s) start======\n", info.Label, abtest.GetName())
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: abtest.GetType(), Tag: info.Tag, Label: info.Label, IsBlock: false}
	rand.Seed(time.Now().UnixNano())
	winNum := rand.Float64() * 100
	var counter float64 = 0
	for _, branch := range abtest.Branchs {
		counter += branch.Percent
		if counter > winNum {
			//feature global.Features.Set(dto.Feature{Name: abtest.Name, Value: branch.Name})
			log.Printf("abtest %v : %v, %v, output:%v \n", abtest.GetName(), branch.Name, winNum, branch.Decision.Output)
			nextNodeName := branch.Decision.Output.Value.(string)
			nextNodeType := GetNodeType(branch.Decision.Output.Kind)
			nodeResult.NextNodeName = nextNodeName
			nodeResult.NextNodeType = nextNodeType
			/*if res, ok := branch.Decision.Output.([]interface{}); ok {
				if len(res) == 2 {
					log.Println("abtest result", res)
					return res, nil
				}
			}*/
		}
	}
	nodeResult.Value = winNum
	log.Printf("======[trace]abtest(%s, %s) end======\n", info.Label, abtest.GetName())
	return nodeResult, nil
}
