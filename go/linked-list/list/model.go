package list

import (
	"reflect"
)

type List struct {
	HeadPointer       *Node
	Length            int
	IsRestrictionType bool
	RestrictionType   reflect.Type
}

type Node struct {
	Next  *Node
	Value any
	Type  reflect.Type
}
