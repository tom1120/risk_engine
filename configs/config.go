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

//策略
type Strategy struct {
	Name     string `yaml:"name"`
	Priority int    `yaml:"priority"` //越大越优先
	Score    int    `yaml:"score"`    //策略分
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

const (
	CONSOLE = "console"
	FILE    = "file"
	DB      = "db"
)
