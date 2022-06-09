package core

type FeatureType int

const (
	TypeInt FeatureType = iota
	TypeFloat
	TypeString
	TypeBool
	TypeEnum
	TypeStrategy //策略结构体
)

var FeatureTypeMap = map[string]FeatureType{
	"TypeInt":      TypeInt,
	"TypeFloat":    TypeFloat,
	"TypeString":   TypeString,
	"TypeBool":     TypeBool,
	"TypeEnum":     TypeEnum,
	"TypeStrategy": TypeStrategy,
}

var FeatureStrMap = map[FeatureType]string{
	TypeInt:      "TypeInt",
	TypeFloat:    "TypeFloat",
	TypeString:   "TypeString",
	TypeBool:     "TypeBool",
	TypeEnum:     "TypeEnum",
	TypeStrategy: "TypeStrategy",
}

func (featureType FeatureType) Get(name string) FeatureType {

	return FeatureTypeMap[name]
}

func (featureType FeatureType) String() string {
	return FeatureStrMap[featureType]
}

type Feature struct {
	name             string
	kind             FeatureType
	value            interface{}
	defaultValue     interface{}
	supportOperators []string
}

func NewFeature(name string, kind FeatureType, defaultValue interface{}) *Feature {
	return &Feature{
		name:         name,
		kind:         kind,
		defaultValue: defaultValue,
	}
}

func (feature *Feature) SetValue(value interface{}) {
	feature.value = value
}

func (feature *Feature) GetValue() (interface{}, bool) {
	if feature.value == nil { //取不到走默认值
		return feature.defaultValue, false
	}
	return feature.value, true
}

func (feature *Feature) GetName() string {
	return feature.name
}
