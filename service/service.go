package service

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/core"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/internal/util"
	"log"
	"time"
)

type EngineService struct {
	startTime int64
	endTime   int64
	kernel    *core.Kernel
}

func NewEngineService(kernel *core.Kernel) *EngineService {
	return &EngineService{kernel: kernel}
}

//dto.DslRunResponse
func (service *EngineService) Run(c *gin.Context, req *dto.EngineRunRequest) (*dto.EngineRunResponse, error) {
	service.startTime = time.Now().UnixNano() / 1e6 //ms
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
	}()
	flow, err := service.kernel.GetDecisionFlow(req.Key, req.Version)
	if err != nil {
		return (*dto.EngineRunResponse)(nil), err
	}

	ctx := core.NewPipelineContext()

	//fill feature value from request features
	features := make(map[string]core.IFeature)
	for name, feature := range flow.FeatureMap {
		if val, ok := req.Features[name]; ok { //in request params
			featureType, err := util.GetType(val) //check data
			if err != nil {                       //warning: unknow type
				log.Println("type check error: ", err)
			}
			if !util.MatchType(featureType, feature.GetType().String()) {
				log.Printf("request feature type is not match! [%s] type is (%s), required (%s)\n", name, core.GetFeatureType(featureType), feature.GetType())
				continue
			}
			features[name] = feature
			if feature.GetType() == core.TypeDate {
				valDate, _ := util.StringToDate(val.(string))
				features[name].SetValue(valDate)
			} else {
				features[name].SetValue(val)
			}
		} else {
			log.Println("request lack feature: ", name)
		}
	}
	log.Println("======[trace]request features ======", features)
	ctx.SetFeatures(features)
	flow.Run(ctx)

	result := ctx.GetDecisionResult()
	return service.dataAdapter(req, result), nil
}

//adapte the result and output
func (service *EngineService) dataAdapter(req *dto.EngineRunRequest, result *core.DecisionResult) *dto.EngineRunResponse {
	resp := &dto.EngineRunResponse{
		Key:       req.Key,
		ReqId:     req.ReqId,
		Uid:       req.Uid,
		StartTime: time.Unix(service.startTime/1000, 0).Format("2006-01-02 15:04:05"),
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
	nodeResults := make([]map[string]interface{}, 0)
	for _, nodeResult := range result.NodeResults {
		if nodeResult == nil {
			continue
		}
		nodeResults = append(nodeResults, map[string]interface{}{
			"name":    nodeResult.Name,
			"id":      nodeResult.Id,
			"Kind":    nodeResult.Kind.String(),
			"tag":     nodeResult.Tag,
			"label":   nodeResult.Label,
			"IsBlock": nodeResult.IsBlock,
			"Value":   nodeResult.Value,
			"Score":   nodeResult.Score,
		})
		i++
	}
	resp.NodeResults = nodeResults

	service.endTime = time.Now().UnixNano() / 1e6
	resp.RunTime = service.endTime - service.startTime
	resp.EndTime = time.Unix(service.endTime/1000, 0).Format("2006-01-02 15:04:05")
	return resp
}
