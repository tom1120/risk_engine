package util

import (
	"testing"
)

func TestNumberic(t *testing.T) {
	t.Log(GetFeatureType(20))
	t.Log(GetFeatureType(30))
	t.Log(GetFeatureType("1111"))
	t.Log(GetFeatureType("11.11"))
	t.Log(GetFeatureType("aa11"))
	t.Log(GetFeatureType(true))
	t.Log(GetFeatureType("true"))
	t.Log(GetFeatureType(1111))
	t.Log(GetFeatureType(1111.1111))
}
