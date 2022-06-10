package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/core"
	"log"
)

var (
	kernel *core.Kernel
)

func Init() { //conf
	kernel = core.NewKernel()
	kernel.LoadDsl("file", "")

	router := gin.Default()
	router.POST("/engine/run", EngineHandler)
	router.Run(":8889") //conf

	log.Printf("[HTTP] Listening on: %s\n", ":8889")
}
