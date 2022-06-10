package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/service"
	"net/http"
)

func EngineHandler(c *gin.Context) {
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
	svr := service.NewEngineService(kernel)
	result, err := svr.Run(c, &request)
	if err != nil {
		code = 501
		errs = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"result": result,
		"error":  errs,
	})
}
