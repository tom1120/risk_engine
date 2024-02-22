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
package configs

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Conf struct {
	Server ServerConf `yaml:"Server"`
	App    AppConf    `yaml:"App"`
}

type ServerConf struct {
	Env          string        `yaml:"Env"`
	Port         int           `yaml:"Port"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

type AppConf struct {
	LogMethod     string `yaml:"LogMethod"` //console,file
	LogPath       string `yaml:"LogPath"`
	DslLoadMethod string `yaml:"DslLoadMethod"` //file,db
	DslLoadPath   string `yaml:"DslLoadPath"`
}

func LoadConfig(path string) (*Conf, error) {
	conf := new(Conf)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}

//策略
type Strategy struct {
	Name     string `yaml:"name"`
	Priority int    `yaml:"priority"` //越大越优先
	Score    int    `yaml:"score"`    //策略分
}

//keywords for execute
const (
	CONSOLE  = "console"
	FILE     = "file"
	DB       = "db"
	PARALLEL = "parallel"
)
