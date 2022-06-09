package core

//各类型节点实现该接口
type INode interface {
	GetName() string
	GetType() NodeType
	GetInfo() NodeInfo
	Parse(*PipelineContext) (*NodeResult, error)
}

//节点返回内容 是否阻断 下一个节点信息(ab,条件节点）
type NodeResult struct {
	IsBlock      bool
	NextNodeName string
	NextNodeType NodeType
}

//all support node
type NodeType int

const (
	TypeStart NodeType = iota
	TypeEnd
	TypeRuleset
	TypeAbtest
	TypeConditional
	TypeDecisiontree
	TypeDecisionmartix
	TypeScorecard
)

var nodeStrMap = map[NodeType]string{
	TypeStart:          "start",
	TypeEnd:            "end",
	TypeRuleset:        "ruleset",
	TypeAbtest:         "abtest",
	TypeConditional:    "conditional",
	TypeDecisiontree:   "decisiontree",
	TypeDecisionmartix: "decisionmartix",
	TypeScorecard:      "scorecard",
}

func (nodeType NodeType) String() string {
	return nodeStrMap[nodeType]
}

var nodeTypeMap map[string]NodeType = map[string]NodeType{
	"start":          TypeStart,
	"end":            TypeEnd,
	"ruleset":        TypeRuleset,
	"abtest":         TypeAbtest,
	"conditional":    TypeConditional,
	"decisiontree":   TypeDecisiontree,
	"decisionmartix": TypeDecisionmartix,
	"scorecard":      TypeScorecard,
}

func GetNodeType(name string) NodeType {
	return nodeTypeMap[name]
}
