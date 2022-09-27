package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/core"
	"github.com/skyhackvip/risk_engine/core/udf"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/internal/log"
	"github.com/skyhackvip/risk_engine/internal/util"
	"time"
)

type EngineService struct {
	startTime time.Time
	kernel    *core.Kernel
}

func NewEngineService(kernel *core.Kernel) *EngineService {
	builtinUdf()
	return &EngineService{kernel: kernel}
}

//dto.DslRunResponse
func (service *EngineService) Run(c *gin.Context, req *dto.EngineRunRequest) (*dto.EngineRunResponse, error) {
	service.startTime = time.Now()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error(err)
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
				log.Errorf("type check error: %s", err)
			}
			if !util.MatchType(featureType, feature.GetType().String()) {
				log.Warnf("request feature type is not match! %s", fmt.Sprintf("%s type is %s, required %s", name, core.GetFeatureType(featureType), feature.GetType()))
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
			log.Warn("request lack feature: %s", name)
		}
	}
	log.Infof("======request features %v======", features)
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
		StartTime: util.TimeFormat(service.startTime),
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
	resp.RunTime = util.TimeSince(service.startTime)
	resp.EndTime = util.TimeFormat(time.Now())
	return resp
}

func builtinUdf() {
	global.RegisterUdf("sum", udf.Sum)
}
