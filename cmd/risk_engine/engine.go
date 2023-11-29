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
package main

import (
	"context"
	"flag"
	"github.com/skyhackvip/risk_engine/api"
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := flag.String("c", "", "config file path")
	flag.Parse()
	conf, err := configs.LoadConfig(*c)
	if err != nil {
		panic(err)
	}
	global.ServerConf = &conf.Server
	global.AppConf = &conf.App

	log.InitLogger(global.AppConf.LogMethod, global.AppConf.LogPath)

	api.Init()

	//graceful restart
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	log.Info("shutdown risk engine...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Warn("timeout of 5 seconds")
	}
	log.Info("server exiting")
}
