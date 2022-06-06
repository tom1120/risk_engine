package core

import (
	"errors"
	"fmt"
	"log"
)

type Flow struct {
	flowMap   map[string]*FlowNode
	startNode *FlowNode
}

func NewFlow() *Flow {
	return &Flow{flowMap: make(map[string]*FlowNode)}
}

func (flow *Flow) AddNode(node *FlowNode) {
	key := flow.getNodeKey(node.NodeName, node.NodeType)
	if _, ok := flow.flowMap[key]; !ok {
		flow.flowMap[key] = node
	} else {
		log.Println("repeat add node: " + key)
	}
}

//NodeType string
func (flow *Flow) GetNode(name string, nodeType interface{}) (*FlowNode, bool) {
	key := flow.getNodeKey(name, nodeType)
	if flowNode, ok := flow.flowMap[key]; ok {
		return flowNode, ok
	}
	return new(FlowNode), false
}

func (flow *Flow) GetAllNodes() map[string]*FlowNode {
	return flow.flowMap
}

func (flow *Flow) getNodeKey(name string, nodeType interface{}) string {
	return fmt.Sprintf("%s-%s", nodeType, name)
}

func (flow *Flow) SetStartNode(startNode *FlowNode) {
	flow.startNode = startNode
}

func (flow *Flow) GetStartNode() (*FlowNode, bool) {
	return flow.startNode, true
}

func (flow *Flow) Run(ctx *PipelineContext) (err error) {

	//recover()

	defer func() {
		if err := recover(); err != nil {
			err = err
			log.Println(err)
			//return err //errors.New("engine run error")
		}
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

		//data, err := flowNode.Parse(flow.Ctx)

		//gotoNext: isbreak || nextnode empty/end

		/*		if flowNode.Node.Category == ABTEST {
				} else {
					flowNode = flowNode.GetNextNode()
				}
		*/

	}

	return
}

//parse current node and return next node
func (flow *Flow) parseNode(curNode *FlowNode, ctx *PipelineContext) (*FlowNode, bool) {

	//parse current node
	res, err := curNode.Parse(ctx)
	if err != nil {
		log.Println(err)
	}

	//get next node
	nextNode := new(FlowNode)
	switch curNode.NodeType { //string int
	case "end": //END:
		return nextNode, false
	case "abtest": //ABTEST:
		next := res.([]interface{})
		log.Println(next[0], next[1])
		return flow.GetNode(next[0].(string), next[1].(string))
	default: //start
		return flow.GetNode(curNode.NextNodeName, curNode.NextNodeType)
	}
	return nextNode, false
}

type FlowNode struct {
	NodeName string   `yaml:"node_name"`
	NodeType NodeType `yaml:"node_type"`

	NextNodeName string   `yaml:"next_node_name"`
	NextNodeType NodeType `yaml:"next_node_type"`

	elem     INode
	nextNode *FlowNode
}

//func NewFlowNode(node INode) *FlowNode {
//	return &FlowNode{NodeName: node.GetName(), Elem: node}
//category  Category: node.GetCategory(),
//}

func (flowNode *FlowNode) SetElem(elem INode) {
	flowNode.elem = elem
}

/*func (flowNode *FlowNode) SetNextNode(nextNode *FlowNode) {
	flowNode.nextNode = nextNode
}

func (flowNode *FlowNode) GetNextNode() *FlowNode {
	return flowNode.nextNode
}*/

func (flowNode *FlowNode) Parse(ctx *PipelineContext) (interface{}, error) {
	return flowNode.elem.Parse(ctx)
}
