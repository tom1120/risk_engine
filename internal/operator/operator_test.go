package operator

import (
	"github.com/skyhackvip/risk_engine/internal/log"
	"testing"
	"time"
)

func TestAInB(t *testing.T) {
	t.Log(AInB([]interface{}{1, 3, 5}, []interface{}{1, 2, 3, 4, 5}))
	t.Log(AInB([]interface{}{1, 3, 5}, []interface{}{1, 4, 5, 6}))
	t.Log(AInB([]interface{}{1, 3, 5}, []interface{}{1, 3}))
	t.Log(AInB([]interface{}{1, 3, 5}, []interface{}{5, 8, 7, 3, 1}))
}

func TestCompare(t *testing.T) {
	log.InitLogger("console", "")
	t.Log("------eq-------")
	t.Log(Compare("EQ", 3, 3.0))
	t.Log(Compare("EQ", 3, "3.000"))
	t.Log(Compare("EQ", "3", "3.000"))
	t.Log(Compare("EQ", "3.0", "3.000"))
	t.Log(Compare("EQ", "3.0", "3.0r00"))
	t.Log(Compare("EQ", "aa", "aa"))
	t.Log(Compare("EQ", "true", true))
	t.Log(Compare("EQ", "true", "true"))
	t.Log(Compare("EQ", true, true))

	t.Log("------ge-------")
	t.Log(Compare("GE", "3.0", "2.99999999"))
	t.Log(Compare("GE", "3", "2.99999999"))
	t.Log(Compare("GE", "3", "2"))
	t.Log(Compare("GE", 3, 2))
	t.Log(Compare("GE", 3, 2.99999999))
	t.Log(Compare("GE", 3.0, 2.99999999))
	t.Log(Compare("GE", 3.0, 3.000000000))

	t.Log("------date-----")
	t.Log(Compare("EQ", "2022-10-10", "2022-10-10 00:00:00"))
	now := time.Now()
	t.Log(Compare("EQ", now, now))

	t.Log("------array-----")
	t.Log(Compare("EQ", []interface{}{3, 8, 7, 6, 9}, []interface{}{9, 6, 7, 3, 8}))
	t.Log(Compare("EQ", []interface{}{"a", "b", "d", "c", "e"}, []interface{}{"a", "b", "c", "d", "e"}))
}

func TestBoolExpr(t *testing.T) {
	variables := map[string]bool{
		"foo": true,
		"bar": false,
		"a1":  true,
	}
	expr := " foo && bar"
	result, err := EvaluateBoolExpr(expr, variables)
	t.Log(expr, result, err)
	t.Log(EvaluateBoolExpr("!(foo&&bar)||!a1", variables))
}
