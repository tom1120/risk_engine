package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/core"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/service"
	"log"
	"net/http"
)

type EngineHandler struct {
	kernel *core.Kernel
}

func NewEngineHandler(kernel *core.Kernel) *EngineHandler {
	return &EngineHandler{kernel: kernel}
}

func (handler *EngineHandler) Run(c *gin.Context) {
	code := 200
	errs := ""
	var request dto.EngineRunRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		code = 500
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  code, //todo
			"error": err.Error(),
		})
		return
	}
	log.Printf("======[trace]request start req_id (%s)======\n", request.ReqId)
	svr := service.NewEngineService(handler.kernel)
	result, err := svr.Run(c, &request)
	if err != nil {
		code = 501
		errs = err.Error()
	}
	log.Printf("======[trace]request end req_id (%s)======\n", request.ReqId)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"result": result,
		"error":  errs,
	})
}

func (handler *EngineHandler) List(c *gin.Context) {
	data := make([]*dto.Dsl, 0)
	for _, flow := range handler.kernel.GetAllDecisionFlow() {
		dsl := &dto.Dsl{Key: flow.Key, Version: flow.Version, Md5: flow.Md5}
		data = append(data, dsl)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": data,
		"error":  "",
	})
}
