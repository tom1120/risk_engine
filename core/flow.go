package core

import (
	"errors"
	"fmt"
	"log"
)

type DecisionFlow struct {
	Key        string
	Version    string
	Metadata   map[string]interface{}
	Md5        string //yaml文件的md5值
	flowMap    map[string]*FlowNode
	startNode  *FlowNode
	FeatureMap map[string]IFeature
}

func NewDecisionFlow() *DecisionFlow {
	return &DecisionFlow{flowMap: make(map[string]*FlowNode)}
}

func (flow *DecisionFlow) AddNode(node *FlowNode) {
	key := flow.getNodeKey(node.NodeName, node.NodeKind)
	if _, ok := flow.flowMap[key]; !ok {
		flow.flowMap[key] = node
	} else {
		log.Println("repeat add node: " + key)
	}
}

//NodeType string
func (flow *DecisionFlow) GetNode(name string, nodeType interface{}) (*FlowNode, bool) {
	key := flow.getNodeKey(name, nodeType)
	if flowNode, ok := flow.flowMap[key]; ok {
		return flowNode, ok
	}
	return new(FlowNode), false
}

func (flow *DecisionFlow) GetAllNodes() map[string]*FlowNode {
	return flow.flowMap
}

func (flow *DecisionFlow) getNodeKey(name string, nodeType interface{}) string {
	return fmt.Sprintf("%s-%s", nodeType, name)
}

func (flow *DecisionFlow) SetStartNode(startNode *FlowNode) {
	flow.startNode = startNode
}

func (flow *DecisionFlow) GetStartNode() (*FlowNode, bool) {
	if flow.startNode == nil {
		return &FlowNode{}, false
	}
	return flow.startNode, true
}

func (flow *DecisionFlow) Run(ctx *PipelineContext) (err error) {
	//recover
	go func() {
		defer func() {
			if err := recover(); err != nil {
				err = err
				log.Println(err)
			}
		}()
	}()

	//find StartNode
	flowNode, ok := flow.GetStartNode()
	if !ok {
		err = errors.New("no start node")
		return
	}

	gotoNext := true
	for gotoNext {
		//ctx.SetCurrentNode(flowNode)
		flowNode, gotoNext = flow.parseNode(flowNode, ctx)
	}
	return
}

//parse current node and return next node
func (flow *DecisionFlow) parseNode(curNode *FlowNode, ctx *PipelineContext) (nextNode *FlowNode, gotoNext bool) {
	//parse current node
	ctx.AddTrack(curNode.GetElem())
	res, err := curNode.Parse(ctx)
	ctx.AddNodeResult(curNode.NodeName, res)

	//error break
	if err != nil {
		log.Println(err)
		return
	}

	//node is block
	if res.IsBlock {
		gotoNext = !res.IsBlock
		return
	}

	//goto next node
	switch curNode.GetNodeType() {
	case TypeEnd: //END:
		gotoNext = false
		return
	case TypeConditional:
		fallthrough
	case TypeAbtest: //ABTEST
		nextNode, gotoNext = flow.GetNode(res.NextNodeName, res.NextNodeType)
		return
	default: //start,matrix,ruleset,tree,scorecard
		nextNode, gotoNext = flow.GetNode(curNode.NextNodeName, curNode.NextNodeKind)
		return
	}
}

type FlowNode struct {
	NodeName     string `yaml:"node_name"`
	NodeKind     string `yaml:"node_kind"`
	NextNodeName string `yaml:"next_node_name"`
	NextNodeKind string `yaml:"next_node_kind"`

	elem     INode
	nextNode *FlowNode
}

func (flowNode *FlowNode) GetNodeType() NodeType {
	return GetNodeType(flowNode.NodeKind)
}

func (flowNode *FlowNode) GetNextNodeType() NodeType {
	return GetNodeType(flowNode.NextNodeKind)
}

func (flowNode *FlowNode) SetElem(elem INode) {
	flowNode.elem = elem
}

func (flowNode *FlowNode) GetElem() INode {
	return flowNode.elem
}

func (flowNode *FlowNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	//before hook
	err := flowNode.elem.BeforeParse(ctx)
	if err != nil {
		return (*NodeResult)(nil), err
	}
	//parse
	result, err := flowNode.elem.Parse(ctx)
	if err != nil {
		return (*NodeResult)(nil), err
	}
	//after hook
	err = flowNode.elem.AfterParse(ctx, result)
	if err != nil {
		return (*NodeResult)(nil), err
	}
	return result, nil
}
