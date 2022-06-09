package global

import (
	"github.com/skyhackvip/risk_engine/core"
	"github.com/skyhackvip/risk_engine/internal/errcode"
)

var DecisionFlowMap = make(map[string]*core.DecisionFlow)

func AddDecisionFlow(key string, flow *core.DecisionFlow) {
	if !checkDecisionFlow(flow) {
		return
	}
	if _, ok := DecisionFlowMap[key]; !ok {
		DecisionFlowMap[key] = flow
	} else {
		//todo已存在，是否覆盖
	}
}

//check checksum
func checkDecisionFlow(flow *core.DecisionFlow) bool {

	//todo
	return true
}

func GetDecisionFlow(key string) (*core.DecisionFlow, error) {
	if flow, ok := DecisionFlowMap[key]; ok {
		return flow, nil
	}
	return (*core.DecisionFlow)(nil), errcode.DslErrorNotFound
}

func GetAllDecisionFlow() map[string]*core.DecisionFlow {
	return DecisionFlowMap
}
