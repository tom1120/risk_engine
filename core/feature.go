package core

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/operator"
	"strings"
	"time"
)

type FeatureType int

const (
	TypeInt FeatureType = iota
	TypeFloat
	TypeString
	TypeBool
	TypeList
	TypeMap
	TypeDate
	TypeDefault
)

var FeatureTypeMap = map[string]FeatureType{
	"int":     TypeInt,
	"float":   TypeFloat,
	"string":  TypeString,
	"bool":    TypeBool,
	"list":    TypeList,
	"map":     TypeMap,
	"date":    TypeDate,
	"default": TypeDefault,
}

var FeatureStrMap = map[FeatureType]string{
	TypeInt:     "int",
	TypeFloat:   "float",
	TypeString:  "string",
	TypeBool:    "bool",
	TypeList:    "list",
	TypeMap:     "map",
	TypeDate:    "date",
	TypeDefault: "default",
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
	GetType() FeatureType
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
	case TypeDate:
		feature = &TypeDateFeature{
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

func (feature *TypeNumFeature) GetType() FeatureType {
	return feature.Kind
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

func (feature *TypeStringFeature) GetType() FeatureType {
	return feature.Kind
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

func (feature *TypeBoolFeature) GetType() FeatureType {
	return feature.Kind
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

//date类型
type TypeDateFeature struct {
	Name         string
	Kind         FeatureType
	Value        interface{}
	DefaultValue interface{}
}

func (feature *TypeDateFeature) GetType() FeatureType {
	return feature.Kind
}

func (feature *TypeDateFeature) SupportOperators() map[string]struct{} {
	return configs.DateSupportOperator
}

func (feature *TypeDateFeature) SetValue(value interface{}) { //format error
	feature.Value = value
}

func (feature *TypeDateFeature) GetValue() (interface{}, bool) {
	return feature.Value, true
}

func (feature *TypeDateFeature) GetName() string {
	return feature.Name
}

func (feature *TypeDateFeature) Compare(op string, target interface{}) (bool, error) {
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errcode.ParseErrorNotSupportOperator
	}
	value, _ := feature.GetValue()
	valueTime, ok := value.(time.Time)
	if !ok {
		return false, errcode.ParseErrorTargetNotSupport
	}
	var (
		targetTime      time.Time
		targetTimeLeft  time.Time
		targetTimeRight time.Time
		isTargetArr     bool
	)
	var err error
	switch target.(type) {
	case string:
		targetTime, err = feature.strToTime(target.(string))
		if err != nil {
			return false, err
		}
	case time.Time:
		targetTime = target.(time.Time)
	case []string:
		if targetArr := target.([]string); len(targetArr) != 2 {
			return false, errcode.ParseErrorTargetNotSupport
		} else {
			targetTimeLeft, err = feature.strToTime(targetArr[0])
			if err != nil {
				return false, err
			}
			targetTimeRight, err = feature.strToTime(targetArr[1])
			if err != nil {
				return false, err
			}
			isTargetArr = true
		}
	case []time.Time:
		if targetArr := target.([]time.Time); len(targetArr) != 2 {
			return false, errcode.ParseErrorTargetNotSupport
		} else {
			targetTimeLeft = targetArr[0]
			targetTimeRight = targetArr[1]
			isTargetArr = true
		}
	default:
		return false, errcode.ParseErrorTargetNotSupport
	}
	if isTargetArr && op != "BETWEEN" || !isTargetArr && op == "BETWEEN" {
		return false, errcode.ParseErrorTargetNotSupport
	}
	switch op {
	case "BEFORE":
		return valueTime.Before(targetTime), nil
	case "AFTER":
		return valueTime.After(targetTime), nil
	case "EQ":
		return valueTime.Equal(targetTime), nil
	case "NEQ":
		return !valueTime.Equal(targetTime), nil
	case "BETWEEN":
		return valueTime.After(targetTimeLeft) && valueTime.Before(targetTimeRight), nil
	}
	return false, errcode.ParseErrorNotSupportOperator
}

func (feature TypeDateFeature) strToTime(str string) (time.Time, error) {
	return time.Parse(configs.DATE_FORMAT_DETAIL, str)
}

//List类型
type TypeListFeature struct {
	Name         string
	Kind         FeatureType
	Value        interface{}
	DefaultValue interface{}
}

func (feature *TypeListFeature) GetType() FeatureType {
	return feature.Kind
}

func (feature *TypeListFeature) SupportOperators() map[string]struct{} {
	return configs.ListSupportOperator
}

func (feature *TypeListFeature) SetValue(value interface{}) {
	feature.Value = value
}

func (feature *TypeListFeature) GetValue() (interface{}, bool) {
	if feature.Value == nil { //取不到走默认值
		return feature.DefaultValue, false
	}
	return feature.Value, true
}

func (feature *TypeListFeature) GetName() string {
	return feature.Name
}

func (feature *TypeListFeature) Compare(op string, target interface{}) (bool, error) {
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errcode.ParseErrorNotSupportOperator
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

func (feature *TypeDefaultFeature) GetType() FeatureType {
	return feature.Kind
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
