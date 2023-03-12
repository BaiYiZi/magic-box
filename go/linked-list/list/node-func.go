package list

import (
	"reflect"
)

func GenerateNode[T any](value T) *Node {
	node := &Node{nil, value, nil}

	node.Type = reflect.TypeOf(node.Value)

	return node
}
