package global

import (
	"github.com/skyhackvip/risk_engine/configs"
)

//from configs
var Strategys = map[string]configs.Strategy{
	"reject":  {"reject", 9, 100},
	"approve": {"approve", 5, 5},
	"record":  {"record", 1, 1},
}
