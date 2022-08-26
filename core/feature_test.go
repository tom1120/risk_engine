package core

import (
	"testing"
	"time"
)

func TestFeature(t *testing.T) {
	t.Log("--------int--------")
	f1 := NewFeature("f1", GetFeatureType("int"))
	f1.SetValue(10)
	rs1, err := f1.Compare("GT", 101)
	t.Log(rs1, err)

	rs1, err = f1.Compare("BETWEEN", []int{1, 20})
	t.Log(rs1, err)

	t.Log("--------string--------")
	f2 := NewFeature("f2", GetFeatureType("string"))
	f2.SetValue("abc")                  //left,value
	like, _ := f2.Compare("LIKE", "ab") //target
	t.Log("like:", like)
	in, _ := f2.Compare("IN", []interface{}{"a", "ab", "abc", 11})
	t.Log("in:", in)

	t.Log("--------bool--------")
	f3 := NewFeature("f3", GetFeatureType("bool"))
	f3.SetValue(true)
	rs3, _ := f3.Compare("EQ", true)
	t.Log(rs3)

	t.Log("--------date--------")
	f4 := NewFeature("f4", GetFeatureType("date"))
	value, _ := time.Parse("2006-01-02 15:04:05", "2022-03-01 10:00:05")
	f4.SetValue(value)
	rs4, err := f4.Compare("AFTER", "2022-03-01")
	t.Log(rs4, err)
	targetTime, _ := time.Parse("2006-01-02 15:04:05", "2022-03-01 10:00:04")
	rs4, err = f4.Compare("AFTER", targetTime)
	t.Log(rs4, err)
	rs4, err = f4.Compare("BETWEEN", []string{"2022-01-01", "2022-08-01 12:00:00"})
	t.Log(rs4, err)

	t.Log("--------array--------")
	f5 := NewFeature("f5", GetFeatureType("array"))
	f5.SetValue([]interface{}{2, 3, 5, 7, 8})
	rs5, err := f5.Compare("EQ", []interface{}{1, 5, 7, 3, 8})
	t.Log(rs5, err)

	f6 := NewFeature("f6", GetFeatureType("array"))
	f6.SetValue([]interface{}{1, 3})
	rs6, err := f6.Compare("IN", []interface{}{1, 3, 5, 7, 9})
	rs6, err = f6.Compare("CONTAIN", []interface{}{1, 3})
	t.Log(rs6, err)
	rs6, err = f6.Compare("CONTAIN", 3)
	t.Log(rs6, err)
	f := NewFeature("f", GetFeatureType("array"))
	f.SetValue([]interface{}{"aa", "bb", 3, 4.0, 5.5, 6})
	rs6, err = f.Compare("IN", []interface{}{"aa", "bb", 3, 4.0, 5.5, 6, 7})
	t.Log("--x xxin-", rs6, err)

	t.Log("--------map--------")
	f7 := NewFeature("f7", GetFeatureType("map"))
	f7.SetValue(map[string]interface{}{"a": 1, "b": 2, "c": 3})
	rs7, err := f7.Compare("KEYEXIST", "b")
	t.Log(rs7, err)
	rs7, err = f7.Compare("VALUEEXIST", 3)
	t.Log(rs7, err)
}
