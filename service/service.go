package service

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/core"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"log"
	"time"
)

type EngineService struct {
	StartTime int64
	EndTime   int64
}

func NewEngineService() *EngineService {
	return &EngineService{}
}

//dto.DslRunResponse
func (service *EngineService) Run(c *gin.Context, req *dto.EngineRunRequest) (*dto.EngineRunResponse, error) {
	service.StartTime = time.Now().UnixNano() / 1e6 //ms
	flow, err := global.GetDecisionFlow(req.Key)
	if err != nil {
		return (*dto.EngineRunResponse)(nil), err
	}
	ctx := core.NewPipelineContext()
	features := make(map[string]*core.Feature)
	for k, v := range req.Features {
		feature := core.NewFeature(k, core.TypeInt, -9999) //todo
		feature.SetValue(v)
		features[k] = feature
	}

	/*
		fMap := map[string]interface{}{"feature_1": 60, "feature_2": 5, "feature_3": 80, "feature_4": 1, "feature_5": 2, "feature_6": 8}
		for k, v := range fMap {
			feature := NewFeature(k, TypeInt, -9999)
			feature.SetValue(v)
			features[k] = feature
		}*/

	ctx.SetFeatures(features)
	flow.Run(ctx)

	result := ctx.GetDecisionResult() //将req放入
	return service.dataAdapter(req, result), nil
}

//todo
func (service *EngineService) dataAdapter(req *dto.EngineRunRequest, result *core.DecisionResult) *dto.EngineRunResponse {
	resp := &dto.EngineRunResponse{
		Key:   req.Key,
		ReqId: req.ReqId,
		Uid:   req.Uid,
		//		HitRules: result.HitRules,
	}
	features := make([]map[string]interface{}, 0)
	for _, feature := range result.Features {
		value, ok := feature.GetValue()
		features = append(features, map[string]interface{}{"name": feature.GetName(),
			"value":     value,
			"isDefault": !ok,
		})
	}
	resp.Features = features
	tracks := make([]map[string]interface{}, 0)
	i := 1
	for _, track := range result.Tracks {
		tracks = append(tracks, map[string]interface{}{"index": i,
			"name":  track.Name,
			"label": track.Label,
		})
		i++
	}
	resp.Tracks = tracks
	hitRules := make([]map[string]interface{}, 0)
	for _, rule := range result.HitRules {
		hitRules = append(hitRules, map[string]interface{}{"id": rule.Id,
			"name":  rule.Name,
			"label": rule.Label,
		})
	}
	resp.HitRules = hitRules
	service.EndTime = time.Now().UnixNano() / 1e6
	resp.RunTime = service.EndTime - service.StartTime
	resp.StartTime = time.Unix(service.StartTime/1000, 0).Format("2006-01-02 15:04:05")
	resp.EndTime = time.Unix(service.EndTime/1000, 0).Format("2006-01-02 15:04:05")
	log.Println(service.StartTime, service.EndTime)

	return resp
}
