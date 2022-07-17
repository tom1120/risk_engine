package core

import (
	"testing"
)

func TestFeature(t *testing.T) {
	f1 := NewFeature("f1", GetFeatureType("int"))
	f1.SetValue(10)
	//rs1, _ := f1.Compare("GT", 101)
	target := []int{1, 20}
	between, _ := f1.Compare("BETWEEN", target)
	t.Log("between:", between)

	f2 := NewFeature("f2", GetFeatureType("string"))
	f2.SetValue("abc")                  //left,value
	like, _ := f2.Compare("LIKE", "ab") //target
	t.Log("like:", like)

	in, _ := f2.Compare("IN", []interface{}{"a", "ab", "abc", 11})
	t.Log("in:", in)

	f3 := NewFeature("f3", GetFeatureType("bool"))
	f3.SetValue(true)
	rs3, _ := f3.Compare("EQ", true)
	t.Log(rs3)

}
