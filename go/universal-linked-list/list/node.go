package list

import (
	"fmt"
	"reflect"
)

type Node struct {
	Next     *Node
	Value    any
	nodeType reflect.Type
}

func GenerateNode[T any](value T) *Node {
	nde := &Node{nil, value, nil}

	nde.nodeType = reflect.TypeOf(nde.Value)

	return nde
}

func (nde *Node) getAddr() string {
	if nde == nil {
		return stringEmpty
	}

	return fmt.Sprintf("%p", nde)
}

func (nde *Node) Equal(ndeT *Node, judgeAddr bool) bool {
	if nde.nodeType != ndeT.nodeType {
		return false
	}

	if nde.Value != ndeT.Value {
		return false
	}

	if judgeAddr {
		if !(nde.getAddr() == ndeT.getAddr()) {
			return false
		}
	}

	return true
}
