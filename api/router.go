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
