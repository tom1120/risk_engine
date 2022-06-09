package main

import (
	"github.com/skyhackvip/risk_engine/api"
	"github.com/skyhackvip/risk_engine/core"
	"github.com/skyhackvip/risk_engine/global"
	"io/ioutil"
	"log"
)

func main() {
	loadDsl()
	router := api.InitRouter()
	router.Run(":8889")
}

func loadDsl() {
	paths := []string{"/home/rong/go/src/github.com/skyhackvip/risk_engine/demo/flow_abtest.yaml",
		"/home/rong/go/src/github.com/skyhackvip/risk_engine/demo/flow_simple.yaml"}
	for _, path := range paths {
		load(path)
	}
	go func() {
		//checkChange(paths)
	}()
}

func load(path string) {
	if !checksum(path) {
		log.Println("check file error", path)
		return
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("load file error", err)
		return
	}
	dsl, err := core.LoadDsl(yamlFile) //todo 干掉core依赖
	if err != nil {
		log.Println("yaml convert error", err)
		return
	}
	global.AddDecisionFlow(dsl.Key, dsl.ConvertToDecisionFlow())
}

func checksum(path string) bool {
	return true
}
