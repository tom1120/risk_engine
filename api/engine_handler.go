// Copyright (c) 2023
//
// @author 贺鹏Kavin
// 微信公众号:技术岁月
// https://github.com/skyhackvip/risk_engine
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/core"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/internal/log"
	"github.com/skyhackvip/risk_engine/service"
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
	log.Infof("======[trace] request start req_id %s======", request.ReqId)
	svr := service.NewEngineService(handler.kernel)
	result, err := svr.Run(c, &request)
	if err != nil {
		code = 501
		errs = err.Error()
	}
	log.Infof("======[trace] request end req_id %s======", request.ReqId)
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
