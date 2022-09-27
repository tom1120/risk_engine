package udf

import (
//	"sync"
)

type Udf func(interface{}) (interface{}, error)

/*
var udfMap map[string]Udf = make(map[string]Udf)
var mu sync.RWMutex

func RegisterUdf(key string, fn Udf) {
	mu.Lock()
	defer mu.Unlock()
	udfMap[key] = fn //override if exists
}

func GetUdf(key string) Udf {
	mu.RLock()
	defer mu.RUnlock()
	return udfMap[key]
}
*/
