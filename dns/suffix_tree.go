package dns

import (
	"strings"
)

type Node struct {
	Node map[string]*Node
}

func NewNode() *Node {
	return &Node{Node: map[string]*Node{}}
}
func NewNodeFromRule(rules ...string) *Node {
	node := NewNode()
	for i := range rules {
		node.Add(strings.Split(rules[i], "."))
	}
	return node
}

func (n *Node) String() string {
	return n.string("")
}
func (n *Node) string(prefix string) (out string) {
	for key, val := range n.Node {
		out += prefix + key + "\n" + val.string(prefix+"    ")
	}
	return
}

func (n *Node) Add(secs []string) {
	length := len(secs)
	switch length {
	case 0:
		return
	case 1:
		n.Node[secs[length-1]] = NewNode()
	default:
		subNode, ok := n.Node[secs[length-1]]
		if !ok {
			subNode = NewNode()
			n.Node[secs[length-1]] = subNode
		}
		subNode.Add(secs[:length-1])
	}
}

func (n *Node) Match(addr string) bool {
	return n.matchSecs(strings.Split(addr, "."))
}

func (n *Node) matchSecs(secs []string) bool {
	length := len(secs)
	if length == 0 {
		switch len(n.Node) {
		case 0:
			return true
		case 1:
			_, ok := n.Node["*"]
			return ok
		default:
			return false
		}
	}

	if n, ok := n.Node[secs[length-1]]; ok {
		return n.matchSecs(secs[:length-1])
	}

	_, ok := n.Node["*"]
	return ok
}
