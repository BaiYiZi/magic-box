package list

import (
	"fmt"
	"reflect"
)

type node struct {
	Next     *node
	Value    any
	nodeType reflect.Type
}

func GenerateNode[T any](value T) *node {
	nde := &node{nil, value, nil}

	nde.nodeType = reflect.TypeOf(nde.Value)

	return nde
}

func (nde *node) getAddr() string {
	if nde == nil {
		return StringEmpty
	}

	return fmt.Sprintf("%p", nde)
}

func (nde *node) Equal(ndeT *node, judgeAddr bool) bool {
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
