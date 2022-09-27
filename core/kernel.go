package core

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/log"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Kernel struct {
	DecisionFlowMap map[string]*DecisionFlow
}

func NewKernel() *Kernel {
	return &Kernel{DecisionFlowMap: make(map[string]*DecisionFlow)}
}

//load dsl from method, default from file path. convert into flow
func (kernel *Kernel) LoadDsl(method, path string) {
	var yamls map[string][]byte
	var err error
	if method == configs.FILE {
		yamls, err = kernel.LoadFromFile(path)
	} else {
		yamls, err = kernel.LoadFromDb()
	}
	if err != nil {
		log.Errorf("load dsl fail, method %s, path %s, err %s", method, path, err)
		return
	}
	for k, v := range yamls {
		dsl := new(Dsl)
		err := yaml.Unmarshal(v, dsl)
		if err != nil {
			log.Errorf("file %s convert dsl error: %s", k, err)
			continue
		}
		if !dsl.CheckValid() {
			log.Errorf("file %s dsl check error: %s", k, err)
			continue
		}
		flow, err := dsl.ConvertToDecisionFlow()
		key := kernel.getMapKey(dsl.Key, dsl.Version)
		if err != nil {
			log.Errorf("dsl %s convert to flow error: %s", key, err)
			continue
		}
		if _, ok := kernel.DecisionFlowMap[key]; ok {
			log.Errorf("dsl load repeat %s", key)
		}
		flow.Md5 = fmt.Sprintf("%x", md5.Sum(v))
		kernel.DecisionFlowMap[key] = flow //重复后一个覆盖前一个
	}
}

func (kernel *Kernel) LoadFromFile(path string) (yamls map[string][]byte, err error) {
	//get file list
	files := make([]string, 0)
	err = filepath.Walk(path, func(fp string, info os.FileInfo, err error) error {
		if filepath.Ext(fp) == ".yaml" {
			files = append(files, fp)
		}
		return err
	})

	//read file
	yamls = make(map[string][]byte)
	for _, file := range files {
		yamlFile, err := ioutil.ReadFile(file)
		if err != nil {
			log.Errorf("load file %s error: %s", file, err)
			continue
		}
		yamls[file] = yamlFile
	}

	if len(yamls) == 0 {
		err = errors.New("no valid dsl") //errcode
		return
	}
	return
}

func (kernel *Kernel) getMapKey(key, version string) string {
	return fmt.Sprintf("%s-%s", key, version)
}

//校验dsl yaml完整性
func (kernel *Kernel) CheckDslValid(dsl *Dsl) bool {
	return true
}

func (kernel *Kernel) LoadFromDb() (yamls map[string][]byte, err error) {
	err = errors.New("not finished")
	return
}

func (kernel *Kernel) GetAllDecisionFlow() map[string]*DecisionFlow {
	return kernel.DecisionFlowMap
}

func (kernel *Kernel) GetDecisionFlow(key, version string) (*DecisionFlow, error) {
	if flow, ok := kernel.DecisionFlowMap[kernel.getMapKey(key, version)]; ok {
		return flow, nil
	}
	return (*DecisionFlow)(nil), errcode.DslErrorNotFound
}
