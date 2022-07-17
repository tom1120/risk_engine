package core

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/operator"
	"strings"
)

type FeatureType int

const (
	TypeInt FeatureType = iota
	TypeFloat
	TypeString
	TypeBool
	//	TypeStrategy //策略结构体
	TypeDefault
)

var FeatureTypeMap = map[string]FeatureType{
	"int":     TypeInt,
	"float":   TypeFloat,
	"string":  TypeString,
	"bool":    TypeBool,
	"default": TypeDefault,
	//	"strategy": TypeStrategy,
}

var FeatureStrMap = map[FeatureType]string{
	TypeInt:     "int",
	TypeFloat:   "float",
	TypeString:  "string",
	TypeBool:    "bool",
	TypeDefault: "default",
	//	TypeStrategy: "strategy",
}

func GetFeatureType(name string) FeatureType {
	return FeatureTypeMap[name]
}

func (featureType FeatureType) String() string {
	return FeatureStrMap[featureType]
}

type Feature struct {
	Id    int    `yaml:"id"`
	Name  string `yaml:"name"`
	Tag   string `yaml:"tag"`
	Label string `yaml:"label"`
	Kind  string `yaml:"kind"`
}

type IFeature interface {
	GetName() string
	SetValue(value interface{})
	GetValue() (interface{}, bool)
	SupportOperators() map[string]struct{}
	Compare(op string, value interface{}) (bool, error)
}

func NewFeature(name string, kind FeatureType) (feature IFeature) {
	switch kind {
	case TypeInt:
		fallthrough
	case TypeFloat:
		feature = &TypeNumFeature{
			Name: name,
			Kind: kind,
		}
	case TypeString:
		feature = &TypeStringFeature{
			Name: name,
			Kind: kind,
		}
	case TypeBool:
		feature = &TypeBoolFeature{
			Name: name,
			Kind: kind,
		}
	default:
		feature = &TypeDefaultFeature{
			Name: name,
			Kind: kind,
		}
	}
	return
}

//数值类型
type TypeNumFeature struct {
	Name         string
	Kind         FeatureType
	Value        interface{}
	DefaultValue interface{}
}

func (feature *TypeNumFeature) SupportOperators() map[string]struct{} {
	return configs.NumSupportOperator
}

func (feature *TypeNumFeature) SetValue(value interface{}) {
	feature.Value = value
}

func (feature *TypeNumFeature) GetValue() (interface{}, bool) {
	if feature.Value == nil { //取不到走默认值
		return feature.DefaultValue, false
	}
	return feature.Value, true
}

func (feature *TypeNumFeature) GetName() string {
	return feature.Name
}

func (feature *TypeNumFeature) Compare(op string, target interface{}) (bool, error) {
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errcode.ParseErrorNotSupportOperator
	}
	value, _ := feature.GetValue() //默认值处理

	switch op {
	case "GT":
		fallthrough
	case "LT":
		fallthrough
	case "GE":
		fallthrough
	case "LE":
		fallthrough
	case "EQ":
		fallthrough
	case "NEQ":
		rs, err := operator.Compare(op, value, target)
		return rs, err
	case "BETWEEN":
		if t, ok := target.([]int); !ok {
			return false, errcode.ParseErrorTargetMustBeArray
		} else {
			if len(t) != 2 {
				return false, errcode.ParseErrorTargetMustBeArray
			}
			if t[0] > t[1] {
				t[0], t[1] = t[1], t[0]
			}
			rs1, err := operator.Compare("GT", value, t[0])
			if err != nil {
				return false, err
			}
			rs2, err := operator.Compare("LT", value, t[1])
			if err != nil {
				return false, err
			}
			return rs1 && rs2, nil

		}
	case "IN":
		if t, ok := target.([]interface{}); !ok {
			return false, errcode.ParseErrorTargetMustBeArray
		} else {
			return operator.InArray(t, value), nil
		}
	default:
		return false, errcode.ParseErrorNotSupportOperator
	}
	return false, errcode.ParseErrorNotSupportOperator
}

//字符串类型
type TypeStringFeature struct {
	Name         string
	Kind         FeatureType
	Value        interface{}
	DefaultValue interface{}
}

func (feature *TypeStringFeature) SupportOperators() map[string]struct{} {
	return configs.StringSupportOperator
}

func (feature *TypeStringFeature) SetValue(value interface{}) {
	feature.Value = value
}

func (feature *TypeStringFeature) GetValue() (interface{}, bool) {
	if feature.Value == nil { //取不到走默认值
		return feature.DefaultValue, false
	}
	return feature.Value, true
}

func (feature *TypeStringFeature) GetName() string {
	return feature.Name
}

func (feature *TypeStringFeature) Compare(op string, target interface{}) (bool, error) {
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errcode.ParseErrorNotSupportOperator
	}
	value, _ := feature.GetValue() //默认值处理

	switch op {
	case "EQ":
		fallthrough
	case "NEQ":
		rs, err := operator.Compare(op, value, target)
		return rs, err
	case "LIKE":
		if ok := strings.Contains(value.(string), target.(string)); ok {
			return true, nil
		} else {
			return false, nil
		}
	case "IN":
		if t, ok := target.([]interface{}); !ok {
			return false, errcode.ParseErrorTargetMustBeArray
		} else {
			return operator.InArray(t, value), nil
		}
	default:
		return false, errcode.ParseErrorNotSupportOperator
	}
	return false, errcode.ParseErrorNotSupportOperator
}

//bool类型
type TypeBoolFeature struct {
	Name         string
	Kind         FeatureType
	Value        interface{}
	DefaultValue interface{}
}

func (feature *TypeBoolFeature) SupportOperators() map[string]struct{} {
	return configs.BoolSupportOperator
}

func (feature *TypeBoolFeature) SetValue(value interface{}) {
	feature.Value = value
}

func (feature *TypeBoolFeature) GetValue() (interface{}, bool) {
	return feature.Value, true
}

func (feature *TypeBoolFeature) GetName() string {
	return feature.Name
}

func (feature *TypeBoolFeature) Compare(op string, target interface{}) (bool, error) {
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errcode.ParseErrorNotSupportOperator
	}
	value, _ := feature.GetValue()
	return operator.Compare(op, value, target)
}

//默认类型
type TypeDefaultFeature struct {
	Name         string
	Kind         FeatureType
	Value        interface{}
	DefaultValue interface{}
}

func (feature *TypeDefaultFeature) SupportOperators() map[string]struct{} {
	return configs.DefaultSupportOperator
}

func (feature *TypeDefaultFeature) SetValue(value interface{}) {
	feature.Value = value
}

func (feature *TypeDefaultFeature) GetValue() (interface{}, bool) {
	if feature.Value == nil { //取不到走默认值
		return feature.DefaultValue, false
	}
	return feature.Value, true
}

func (feature *TypeDefaultFeature) GetName() string {
	return feature.Name
}

func (feature *TypeDefaultFeature) Compare(op string, target interface{}) (bool, error) {
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errcode.ParseErrorNotSupportOperator
	}
	value, _ := feature.GetValue() //默认值处理
	return operator.Compare(op, value, target)
}
