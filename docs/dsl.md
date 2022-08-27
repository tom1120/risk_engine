# DSL语法详解
完整案例参见 /demo 目录

# 规则基础部分
- Rule 规则
- Condition 表达式
- Decision 决策
- Output 结果

## 规则定义 Rule
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

Yaml DSL 描述一条规则
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

规则解释：

```shell
如果：
    feature_2 < 8  || feature_3 > 9
那么：
     输出结果：record
     变量赋值：feat1 = aa   feat2 = bb
```

- 规则名称 rule.name 全局唯一
- todo: 增加规则优先级

## 条件表达式定义 Condition

```golang
type Condition struct {
    Feature  string      `yaml:"feature"`
    Operator string      `yaml:"operator"`
    Value    interface{} `yaml:"value"`
    Result   string      `yaml:"result"`
    Name     string      `yaml:"name"`
}
```

Yaml DSL 描述表达式
```yaml
- condition:
  name: c4
  feature: feature_2
  operator: LT
  value: 8
```
- 表达式解释：feature_2 < 8
- 条件表达式是最基础单位，由特征、阈值、操作符组成。
- 条件表达式可以组成规则。


## 决策 Decision

```golang
type Decision struct {
    Depends []string               `yaml:"depends,flow"` //依赖condition结果
    Logic   string                 `yaml:"logic"`
    Output  Output                 `yaml:"output"`
    Assign  map[string]interface{} `yaml:"assign"` //赋值更多变量
}
```
Yaml DSL 描述决策内容

```yaml
decision:
    depends: [c4, c5]
    logic: c4 || c5
    output:
      name:
      value: record
      kind: string
    assign:
      feat1: aa
      feat2: bb
```
- 决策可以输出指定内容，也可以给特征变量赋值
- 决策结果：输出字符串 record，赋值 feat1 = aa   feat2 = bb


## 输出结果 Output

```golang
type Output struct {
    Name  string      `yaml:"name"` //该节点输出值重命名，如果无则以（节点类型+节点名）赋值变量
    Value interface{} `yaml:"value"`
    Kind  string      `yaml:"kind"`
    Hit   bool  //是否命中
}
```
Yaml DSL 描述输出内容
```yaml
output:
  name:
  value: record
  kind: string
```

- kind 包括 FeatureType 和 NodeType 两种
- 条件/ab节点 decision.output 是分支类型 NodeType
- 规则集节点  decision.output 是特征类型 FeatureType


# 特征 feature

```golang
type Feature struct {
  Id    int    `yaml:"id"`
  Name  string `yaml:"name"`
  Tag   string `yaml:"tag"`
  Label string `yaml:"label"`
  Kind  string `yaml:"kind"`
}
```
Yaml DSL 描述特征

```yaml
features:
  - feature:
    id: 1
    name: num_feature
    tag: aa
    label: 数值特征
    kind: int
  - feature:
    id: 2
    name: str_feature
    tag: aa
    label: 字符串特征
    kind: string
  - feature:
    id: 3
    name: bool_feature
    tag: aa
    label: 布尔特征
    kind: bool
  - feature:
    id: 4
    name: date_feature
    tag: aa
    label: 日期特征
    kind: date
  - feature:
    id: 5
    name: array_feature
    tag: aa
    label: 数组特征
    kind: array
  - feature:
    id: 6
    name: map_feature
    tag: aa
    label: 字典特征
    kind: map
```

不同特征支持不同的操作符：

```
var OperatorMap = map[string]string{
  GT:         ">",
  LT:         "<",
  GE:         ">=",
  LE:         "<=",
  EQ:         "==",
  NEQ:        "!=",
  BETWEEN:    "between",
  LIKE:       "like",
  IN:         "in",
  CONTAIN:    "contain",
  BEFORE:     "before",
  AFTER:      "after",
  KEYEXIST:   "keyexist",
  VALUEEXIST: "valueexist",
}
```
- 数值型支持： >   <   >=   <=    ==   !=    between    in
- 字符串支持： ==    !=   like   in
- 布尔型支持： ==    !=
- 日期型支持： ==    !=    before   after   between
- 数组型支持： ==    !=    contain   in
- 字典型支持： keyexist    valueexist


# 决策流节点实现
```golang
//各类型节点实现该接口
type INode interface {
    GetName() string
    GetType() NodeType
    GetInfo() NodeInfo
    Parse(*PipelineContext) (*NodeResult, error)
}
```

# 规则集类型节点
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

# 冠军挑战者类型节点 abtest

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
## 阻断策略 BlockStrategy
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

