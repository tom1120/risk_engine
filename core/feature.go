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
	Name    string
	Type    FeatureType
	Value   interface{}
	Default interface{}
}
