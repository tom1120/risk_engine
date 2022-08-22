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
	rs4, err := f4.Compare("AFTER", "2022-03-01 10:00:04")
	t.Log(rs4, err)
	targetTime, _ := time.Parse("2006-01-02 15:04:05", "2022-03-01 10:00:04")
	rs4, err = f4.Compare("AFTER", targetTime)
	t.Log(rs4, err)
	rs4, err = f4.Compare("BETWEEN", []string{"2022-01-01 00:00:00", "2022-08-01 12:00:00"})
	t.Log(rs4, err)
}
