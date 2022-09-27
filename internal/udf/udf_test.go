package udf

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Log(Sum(100))
	t.Log(Sum(100.5))
	t.Log(Sum([]interface{}{3.3, 100, 5.0, 8}))
	t.Log(Sum([]interface{}{3.3, 100, 5.0, 8, "88"}))
	t.Log(Sum([]interface{}{3.3, 100, 5.0, 8, true}))
	t.Log(Sum([]interface{}{3.3, 100, 5.0, 8, "8ab"}))
	t.Log(Sum([3]int{3, 4, 5}))
	t.Log(Sum([3]float32{3.0, 4.1, 5}))
	t.Log(Sum([3]string{"3.0", "4.1", "5"}))
	t.Log(Sum([3]string{"3.a", "b.1", "5"}))
}
