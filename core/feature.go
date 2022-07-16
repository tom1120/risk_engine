package core

import (
	"errors"
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/operator"
)

type FeatureType int

const (
	TypeInt FeatureType = iota
	TypeFloat
	TypeString
	TypeBool
	TypeEnum
	TypeStrategy //策略结构体
	TypeDefault
)

var FeatureTypeMap = map[string]FeatureType{
	"int":      TypeInt,
	"float":    TypeFloat,
	"string":   TypeString,
	"bool":     TypeBool,
	"enum":     TypeEnum,
	"default":  TypeDefault,
	"strategy": TypeStrategy,
}

var FeatureStrMap = map[FeatureType]string{
	TypeInt:      "int",
	TypeFloat:    "float",
	TypeString:   "string",
	TypeBool:     "bool",
	TypeEnum:     "enum",
	TypeDefault:  "default",
	TypeStrategy: "strategy",
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

	//in, like
}

//default
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
		return false, errors.New("not support operator")
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
	default:
		return false, errors.New("not support operator1")
	}
	return false, errors.New("not support operator2")
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
	//todo
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errors.New("not support operator")
	}
	value, _ := feature.GetValue() //默认值处理
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
		return false, errors.New("not support operator")
	}
	value, _ := feature.GetValue() //默认值处理
	return operator.Compare(op, value, target)
}
