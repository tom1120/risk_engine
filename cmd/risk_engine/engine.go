package main

import (
	"github.com/skyhackvip/risk_engine/api"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	api.Init()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Println("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//cancel()
			time.Sleep(time.Second)
			log.Println("risk_engine quit!")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
