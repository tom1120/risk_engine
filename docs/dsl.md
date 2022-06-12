# DSL语法详解
完整案例参见 /demo 目录

# common部分
- NodeInfo
- BlockStrategy
- Rule
- Condition
- Decision
- Output


## NodeInfo
```golang
type NodeInfo struct {
    Id      int64    `yaml:"id"`
    Name    string   `yaml:"name"`
    Tag     string   `yaml:"tag"`
    Label   string   `yaml:"label"`
    Kind    NodeType `yaml:"kind"`
    Depends []string `yaml:"depends,flow"`
}
```
```yaml
info:
      id: 111
      name: ruleset_1
      tag: internal
      label: 内部规则集1
      kind: ruleset
      depends: [feature_1,feature_2,feature_e]
```

## BlockStrategy
```golang
type BlockStrategy struct {
	IsBlock  bool        `yaml:"is_block"`
    HitRule  []string    `yaml:"hit_rule,flow"`
	Operator string      `yaml:"operator"`
    Value    interface{} `yaml:"value"`
}
```
```yaml
block_strategy:
      is_block: true
      hit_rule: [rule_1] #命中该规则就中断
      operator: EQ
      value: reject  #=reject中断
```


## Rule
```golang
type Rule struct {
    Name       string      `yaml:"name"`
    Tag        string      `yaml:"tag"`
    label      string      `yaml:"label"`
    Conditions []Condition `yaml:"conditions,flow"`
    Decision   Decision    `yaml:"decision"`
    Depends    []string    `yaml:"depends"`
}
```
```yaml
    - rule:
      name: rule_4
      tag: rule_4
      label: 规则4
      depends: [feature_2, feature_3]
      conditions:
      - condition:
        name: c4
        feature: feature_2
        operator: LT
        value: 8
      - condition:
        name: c5
        feature: feature_3
        operator: GT
        value: 9
      decision:
        depends: [c4, c5]
        logic: c4 || c5
        output:
          name:
          value: record
          kind: TypeStrategy
        assign:
          feat1: aa
          feat2: bb
```
规则名称全局唯一
todo: 增加规则优先级

## Condition

```golang
type Condition struct {
    Feature  string      `yaml:"feature"`
    Operator string      `yaml:"operator"`
    Value    interface{} `yaml:"value"`
    Result   string      `yaml:"result"`
    Name     string      `yaml:"name"`
}
```

```yaml
 - condition:
        name: c4
        feature: feature_2
        operator: LT
        value: 8
```

后面要扩展特征和特征比对。

## Decision

```golang
type Decision struct {
    Depends []string               `yaml:"depends,flow"` //依赖condition结果
    Logic   string                 `yaml:"logic"`
    Output  Output                 `yaml:"output"`
    Assign  map[string]interface{} `yaml:"assign"` //赋值更多变量
}
```



```yaml
decision:
        depends: [c4, c5]
        logic: c4 || c5
        output:
          name:
          value: record
          kind: TypeStrategy
        assign:
          feat1: aa
          feat2: bb
```

## Output

```golang
type Output struct {
    Name  string      `yaml:"name"` //该节点输出值重命名，如果无则以（节点类型+节点名）赋值变量
    Value interface{} `yaml:"value"`
    Kind  string      `yaml:"kind"`
    Hit   bool  //是否命中
}
```
```yaml
 output:
          name:
          value: record
          kind: TypeStrategy
```

kind 包括 featureType 和 NodeType 两种不同的
条件/ab节点 decision.output 是分支 类型是 NodeType
规则集节点。decision.output 是 featureType.TypeStrategy 策略




节点返回值：
1、是否阻断
2、对于ab节点和条件节点，要带下个节点的信息
3、对于规则集节点，决策节点 好像没啥重要的
```golang
//节点返回内容 是否阻断 下一个节点信息(ab,条件节点）
type NodeResult struct {
    Id           int64
    Name         string
    Label        string
    Tag          string
    Kind         NodeType
    IsBlock      bool
    Score        int
    Value        interface{}
    NextNodeName string //ab,条件节点有用
    NextNodeType NodeType
}
```

# 各节点实现接口
```golang
//各类型节点实现该接口
type INode interface {
    GetName() string
    GetType() NodeType
    GetInfo() NodeInfo
    Parse(*PipelineContext) (*NodeResult, error)
}
```

# 规则集
```golang
type RulesetNode struct {
    Info          NodeInfo      `yaml:"info"`
    ExecPlan      string        `yaml:"exec_plan"`
    BlockStrategy BlockStrategy `yaml:"block_strategy"`
    Rules         []Rule        `yaml:"rules,flow"`
}
```

```yaml
 - ruleset:
    info:
      id: 111
      name: ruleset_1
      tag: internal
      label: 内部规则集1
      kind: ruleset
      depends: [feature_1,feature_2,feature_e]
    exec_plan: parallel
    block_strategy:
      is_block: true
      hit_rule: [rule_1] #命中该规则就中断
      operator: EQ
      value: reject  #=reject中断
    rules:
    - rule:
      name: rule_1
      tag: tag1
      label: 规则1
      conditions:
      - condition:
        name: c1
        feature: feature_1
        operator: GT
        value: 50
      decision:
        depends: [c1]
        logic: c1
        output:
          name: #rename 如果为空，则给rule.name赋值
          value: reject
          kind: TypeStrategy
        assign:
          feature_x: 111
```

# abtest

```golang
type AbtestNode struct {
    Info    NodeInfo `yaml:"info"`
    Branchs []Branch `yaml:"branchs,flow"`
}
```


```yaml
- abtest:
    info:
      id: 11
      name: abtest_1
      tag: tag_ab
      label: 分流实验
      kind: abtest
      depends:
    branchs:
    - branch:
      name: branch_1
      percent: 44.5
      decision:
        depends:
        logic: random
        output:
          name:
          value: ruleset_1
          kind: ruleset
        assign:
    - branch:
      name: branch_2
      percent: 55.5
      decision:
        depends:
        logic: random
        output:
          name:
          value: ruleset_2
          kind: ruleset
```


# 策略
```golang
//策略
type Strategy struct {
    Name     string `yaml:"name"`
    Priority int    `yaml:"priority"` //越大越优先
    Score    int    `yaml:"score"`    //策略分
}

var Strategys = map[string]Strategy{
    "reject":  {"reject", 9, 100},
    "approve": {"approve", 5, 5},
    "record":  {"record", 1, 1},
}

//阻断策略
var BlockStrategy = []string{"reject"}
```

