package core

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
	"sync"
)

type MatrixNode struct {
	Info           NodeInfo `yaml:"info"`
	Rules          []Rule   `yaml:"rules,flow"`
	DependFeatures map[string]IFeature
	ExecPlan       string         `yaml:"exec_plan"`
	MatrixStrategy MatrixStrategy `yaml:"matrix_strategy"`
}

type MatrixStrategy struct {
	OutputName string   `yaml:"output_name"`
	OutputKind string   `yaml:"output_kind"`
	Depend     []string `yaml:"depend"`
	Cases      []Case   `yaml:"cases"`
}

type Case struct {
	Case   []string `yaml:"case"`
	Output string   `yaml:"output"`
}

func (matrixNode MatrixNode) GetName() string {
	return matrixNode.Info.Name
}

func (matrixNode MatrixNode) GetType() NodeType {
	return GetNodeType(matrixNode.Info.Kind)
}

func (matrixNode MatrixNode) GetInfo() NodeInfo {
	return matrixNode.Info
}

func (matrixNode MatrixNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (matrixNode MatrixNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (matrixNode MatrixNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	info := matrixNode.GetInfo()
	log.Printf("======[trace]matrix(%s, %s) start======\n", info.Label, matrixNode.GetName())
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: matrixNode.GetType(), Tag: info.Tag, Label: info.Label}

	depends := ctx.GetFeatures(info.Depends)

	var outputs = make(map[string]*Output, 0)
	var xResult string
	var yResult string

	if matrixNode.ExecPlan == "parallel" { //并发分组执行
		ruleMap := make(map[string][]Rule)
		for _, rule := range matrixNode.Rules {
			if rule.Kind == configs.MATRIXX {
				ruleMap[configs.MATRIXX] = append(ruleMap[configs.MATRIXX], rule)
			} else if rule.Kind == configs.MATRIXY {
				ruleMap[configs.MATRIXY] = append(ruleMap[configs.MATRIXY], rule)
			}
		}
		var wg sync.WaitGroup
		var mu sync.Mutex
		for key, rules := range ruleMap {
			wg.Add(1)
			go func(key string, rules []Rule) {
				defer wg.Done()
				for _, rule := range rules {
					output, err := rule.Parse(ctx, depends)
					//continue if error
					if err != nil {
						log.Println(err)
						continue
					}
					//continue if miss hit
					if !output.GetHit() {
						continue
					}
					//break for loop if hit once (only hit one rule in one group)
					ctx.AddHitRule(&rule)
					mu.Lock()
					outputs[rule.Name] = output
					if key == configs.MATRIXX {
						xResult = rule.Name
					} else if key == configs.MATRIXY {
						yResult = rule.Name
					}
					mu.Unlock()
					break
				}
			}(key, rules)
		}
		wg.Wait()
	} else { //串行执行
		for _, rule := range matrixNode.Rules {
			output, err := rule.Parse(ctx, depends)
			if err != nil {
				log.Println(err)
				continue
			}
			if !output.GetHit() {
				continue
			}
			//if hit rule
			ctx.AddHitRule(&rule)
			outputs[rule.Name] = output
			if rule.Kind == configs.MATRIXX {
				xResult = rule.Name
			} else if rule.Kind == configs.MATRIXY {
				yResult = rule.Name
			}
		}
	}

	//match result
	if xResult == "" || yResult == "" {
		return nodeResult, errcode.ParseErrorMatrixNotMatch
	}
	val, kind, ok := matrixNode.matchResult(xResult, yResult)
	if !ok {
		return nodeResult, errcode.ParseErrorMatrixOutputEmpty
	}

	//save into ctx feature
	feature := NewFeature(matrixNode.GetName(), GetFeatureType(kind))
	feature.SetValue(val)
	ctx.SetFeature(feature)
	if matrixNode.MatrixStrategy.OutputName != "" { //extra
		extraFeature := NewFeature(matrixNode.MatrixStrategy.OutputName, GetFeatureType(kind))
		extraFeature.SetValue(val)
		ctx.SetFeature(extraFeature)
	}

	//output
	nodeResult.Value = val
	log.Printf("======[trace]matrix(%s, %s) end======\n", info.Label, matrixNode.GetName())
	return nodeResult, nil
}

func (matrixNode MatrixNode) matchResult(xResult, yResult string) (interface{}, string, bool) {
	caseMap := make(map[string]interface{})
	for _, c := range matrixNode.MatrixStrategy.Cases {
		key := c.Case[0] + c.Case[1]
		caseMap[key] = c.Output
	}
	if val, ok := caseMap[xResult+yResult]; ok {
		return val, matrixNode.MatrixStrategy.OutputKind, true
	}
	return nil, "", false
}
