package global

import (
	"github.com/skyhackvip/risk_engine/configs"
)

var Strategys = map[string]configs.Strategy{
	"reject":  {"reject", 9, 100},
	"approve": {"approve", 5, 5},
	"record":  {"record", 1, 1},
}

//阻断策略
var BlockStrategy = map[string]struct{}{"reject": struct{}{}}
