package core

type Decision struct {
	Depends []string               `yaml:"depends,flow"`
	Logic   string                 `yaml:"logic"`
	Output  interface{}            `yaml:"output"` //该节点输出值
	Assign  map[string]interface{} `yaml:"assign"` //赋值
}
