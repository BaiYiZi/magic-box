package list

import (
	"fmt"
	"reflect"
)

type Node struct {
	Next  *Node
	Value any
	Type  reflect.Type
}

func GenerateNode[T any](value T) *Node {
	node := &Node{nil, value, nil}

	node.Type = reflect.TypeOf(node.Value)

	return node
}

func (node *Node) GetAddr() string {
	if node == nil {
		return StringEmpty
	}
	return fmt.Sprintf("%p", node)
}

func (node *Node) Equal(nodeT *Node, judgeAddr bool) bool {
	if node.Type != nodeT.Type {
		return false
	}

	if node.Value != nodeT.Value {
		return false
	}

	if judgeAddr {
		if !(node.GetAddr() == nodeT.GetAddr()) {
			return false
		}
	}

	return true
}
