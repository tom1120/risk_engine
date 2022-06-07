package core

type FeatureType int

const (
	TypeInt FeatureType = iota
	TypeFloat
	TypeString
	TypeBool
	TypeEnum
)

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
