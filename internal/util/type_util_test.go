package util

import (
	"testing"
)

func TestNumberic(t *testing.T) {
	t.Log(GetType(20))
	t.Log(GetType(30))
	t.Log(GetType("1111"))
	t.Log(GetType("11.11"))
	t.Log(GetType("aa11"))
	t.Log(GetType(true))
	t.Log(GetType("true"))
	t.Log(GetType(1111))
	t.Log(GetType(1111.1111))
}
