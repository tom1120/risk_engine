package core

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/operator"
	"github.com/skyhackvip/risk_engine/internal/util"
	"strings"
	"time"
)

type FeatureType int

const (
	TypeInt FeatureType = iota
	TypeFloat
	TypeString
	TypeBool
	TypeDate
	TypeArray
	TypeMap
	TypeDefault
)

var FeatureTypeMap = map[string]FeatureType{
	"int":     TypeInt,
	"float":   TypeFloat,
	"string":  TypeString,
	"bool":    TypeBool,
	"date":    TypeDate,
	"array":   TypeArray,
	"map":     TypeMap,
	"default": TypeDefault,
}

var FeatureStrMap = map[FeatureType]string{
	TypeInt:     "int",
	TypeFloat:   "float",
	TypeString:  "string",
	TypeBool:    "bool",
	TypeDate:    "date",
	TypeArray:   "array",
	TypeMap:     "map",
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
	SetValue(value interface{}) error
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
	case TypeArray:
		feature = &TypeArrayFeature{
			Name: name,
			Kind: kind,
		}
	case TypeMap:
		feature = &TypeMapFeature{
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

func (feature *TypeNumFeature) SetValue(value interface{}) error {
	if err := checkValue(value, feature.GetType()); err != nil {
		return err
	}
	feature.Value = value
	return nil
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
	case configs.EQ:
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

func (feature *TypeStringFeature) SetValue(value interface{}) error {
	if err := checkValue(value, feature.GetType()); err != nil {
		return err
	}
	feature.Value = value
	return nil
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
	case configs.EQ:
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

func (feature *TypeBoolFeature) SetValue(value interface{}) error {
	if err := checkValue(value, feature.GetType()); err != nil {
		return err
	}
	feature.Value = value
	return nil
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

func (feature *TypeDateFeature) SetValue(value interface{}) error {
	if err := checkValue(value, feature.GetType()); err != nil {
		return err
	}
	feature.Value = value
	return nil
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
		return false, errcode.ParseErrorFeatureTypeNotMatch
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
		targetTime, err = util.StringToDate(target.(string))
		if err != nil {
			return false, err
		}
	case time.Time:
		targetTime = target.(time.Time)
	case []string:
		if targetArr := target.([]string); len(targetArr) != 2 {
			return false, errcode.ParseErrorTargetNotSupport
		} else {
			targetTimeLeft, err = util.StringToDate(targetArr[0])
			if err != nil {
				return false, err
			}
			targetTimeRight, err = util.StringToDate(targetArr[1])
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
	case configs.EQ:
		return valueTime.Equal(targetTime), nil
	case "NEQ":
		return !valueTime.Equal(targetTime), nil
	case "BETWEEN":
		return valueTime.After(targetTimeLeft) && valueTime.Before(targetTimeRight), nil
	}
	return false, errcode.ParseErrorNotSupportOperator
}

//Array类型
type TypeArrayFeature struct {
	Name         string
	Kind         FeatureType
	Value        interface{}
	DefaultValue interface{}
}

func (feature *TypeArrayFeature) GetType() FeatureType {
	return feature.Kind
}

func (feature *TypeArrayFeature) SupportOperators() map[string]struct{} {
	return configs.ArraySupportOperator
}

func (feature *TypeArrayFeature) SetValue(value interface{}) error {
	if err := checkValue(value, feature.GetType()); err != nil {
		return err
	}
	valueArr, ok := value.([]interface{})
	if !ok {
		return errcode.ParseErrorFeatureSetValue
	}
	for k, v := range valueArr {
		if util.IsInt(v) {
			valueArr[k], _ = util.ToInt(v)
		}
	}
	feature.Value = valueArr
	return nil
}

func (feature *TypeArrayFeature) GetValue() (interface{}, bool) {
	if feature.Value == nil { //取不到走默认值
		return feature.DefaultValue, false
	}
	return feature.Value, true
}

func (feature *TypeArrayFeature) GetName() string {
	return feature.Name
}

func (feature *TypeArrayFeature) Compare(op string, target interface{}) (bool, error) {
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errcode.ParseErrorNotSupportOperator
	}
	value, _ := feature.GetValue()
	valueArr, ok := value.([]interface{})
	if !ok {
		return false, errcode.ParseErrorFeatureTypeNotMatch
	}
	targetArr, isTargetArray := target.([]interface{})
	if !isTargetArray && op != configs.CONTAIN { //contain target can be simple data
		return false, errcode.ParseErrorTargetMustBeArray
	}
	switch op {
	case configs.EQ:
		return operator.CompareArray(valueArr, targetArr), nil
	case configs.NEQ:
		return !operator.CompareArray(valueArr, targetArr), nil
	case configs.IN:
		return operator.AInB(valueArr, targetArr), nil
	case configs.CONTAIN:
		if isTargetArray {
			return operator.AInB(targetArr, valueArr), nil
		}
		return operator.InArray(valueArr, target), nil
	}
	return false, errcode.ParseErrorNotSupportOperator
}

//map类型
type TypeMapFeature struct {
	Name         string
	Kind         FeatureType
	Value        interface{}
	DefaultValue interface{}
}

func (feature *TypeMapFeature) GetType() FeatureType {
	return feature.Kind
}

func (feature *TypeMapFeature) SupportOperators() map[string]struct{} {
	return configs.MapSupportOperator
}

func (feature *TypeMapFeature) SetValue(value interface{}) error {
	if err := checkValue(value, feature.GetType()); err != nil {
		return err
	}
	feature.Value = value
	return nil
}

func (feature *TypeMapFeature) GetValue() (interface{}, bool) {
	if feature.Value == nil { //取不到走默认值
		return feature.DefaultValue, false
	}
	return feature.Value, true
}

func (feature *TypeMapFeature) GetName() string {
	return feature.Name
}

func (feature *TypeMapFeature) Compare(op string, target interface{}) (bool, error) {
	if _, ok := feature.SupportOperators()[op]; !ok {
		return false, errcode.ParseErrorNotSupportOperator
	}
	value, _ := feature.GetValue() //默认值处理
	valueMap, ok := value.(map[string]interface{})
	if !ok {
		return false, errcode.ParseErrorFeatureTypeNotMatch
	}
	switch op {
	case configs.KEYEXIST:
		targetStr, err := util.ToString(target)
		if err != nil {
			return false, err
		}
		if _, ok := valueMap[targetStr]; ok {
			return true, nil
		}
		return false, nil
	case configs.VALUEEXIST:
		for _, v := range valueMap {
			if v == target {
				return true, nil
			}
		}
		return false, nil
	}
	return false, errcode.ParseErrorNotSupportOperator
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

func (feature *TypeDefaultFeature) SetValue(value interface{}) error {
	feature.Value = value
	return nil
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

func checkValue(value interface{}, featureType FeatureType) error {
	valueType, err := util.GetType(value)
	if err != nil {
		return errcode.ParseErrorFeatureSetValue
	}
	if GetFeatureType(valueType) != featureType {
		return errcode.ParseErrorFeatureSetValue
	}
	return nil
}
