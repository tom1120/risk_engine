package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/core"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/log"
)

func Init() { //conf
	kernel := core.NewKernel()
	kernel.LoadDsl(global.AppConf.DslLoadMethod, global.AppConf.DslLoadPath)

	engineHandler := NewEngineHandler(kernel)

	router := gin.Default()
	router.POST("/engine/run", engineHandler.Run)
	router.GET("/engine/list", engineHandler.List)

	router.Run(fmt.Sprintf(":%d", global.ServerConf.Port)) //conf

	log.Infof("[HTTP] Listening on %d", global.ServerConf.Port)
}
