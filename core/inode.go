package core

//各类型节点实现该接口
type INode interface {
	GetName() string
	GetKind() NodeType
	GetLabel() string
	GetTag() string
	Parse(*PipelineContext) (interface{}, error)
}

//all support node
type NodeType string

const (
	TypeStart          NodeType = "start"
	TypeEnd            NodeType = "end"
	TypeRuleset        NodeType = "ruleset"
	TypeAbtest         NodeType = "abtest"
	TypeConditional    NodeType = "conditional"
	TypeDecisiontree   NodeType = "decisiontree"
	TypeDecisionmartix NodeType = "decisionmartrix"
	TypeScorecard      NodeType = "scorecard"
)
