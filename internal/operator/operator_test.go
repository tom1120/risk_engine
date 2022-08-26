package operator

import (
	"testing"
)

func TestCompareArray(t *testing.T) {
	t.Log(CompareArray([]interface{}{3, 8, 7, 6, 9}, []interface{}{9, 6, 7, 3, 8}))
	t.Log(CompareArray([]interface{}{"a", "b", "d", "c", "e"}, []interface{}{"a", "b", "c", "d", "e"}))
}

func TestAInB(t *testing.T) {
	t.Log(AInB([]interface{}{1, 3, 5}, []interface{}{1, 2, 3, 4, 5}))
	t.Log(AInB([]interface{}{1, 3, 5}, []interface{}{1, 4, 5, 6}))
	t.Log(AInB([]interface{}{1, 3, 5}, []interface{}{1, 3}))
	t.Log(AInB([]interface{}{1, 3, 5}, []interface{}{5, 8, 7, 3, 1}))
}
