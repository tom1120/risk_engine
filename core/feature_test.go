package core

import (
	"testing"
)

func TestFeature(t *testing.T) {
	feature := NewFeature("f1", GetFeatureType("int"))
	feature.SetValue(10)
	rs, err := feature.Compare("EQ", 101)
	t.Log(rs, err)
}
