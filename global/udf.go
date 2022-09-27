package global

import (
	"github.com/skyhackvip/risk_engine/internal/udf"
	"sync"
)

var udfMap map[string]udf.Udf = make(map[string]udf.Udf)
var mu sync.RWMutex

func RegisterUdf(key string, fn udf.Udf) {
	mu.Lock()
	defer mu.Unlock()
	udfMap[key] = fn //override if exists
}

func GetUdf(key string) udf.Udf {
	mu.RLock()
	defer mu.RUnlock()
	return udfMap[key]
}
