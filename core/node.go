package core

import ()

type INode interface {
	GetName() string
	GetType() NodeType
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
