package main

import (
	"context"
	"flag"
	"github.com/skyhackvip/risk_engine/api"
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/global"
	"log"
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

	api.Init()

	//graceful restart
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	log.Println("shutdown risk engine...")
	//cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	/*if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error:", err)
	}*/
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}
	log.Println("server exiting")
}
