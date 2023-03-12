package list

import (
	"fmt"
	"reflect"
)

func GenerateList[T any](values *[]T, restrictionType reflect.Type) (*List, error) {
	list := &List{}

	if restrictionType != nil {
		list.RestrictionType = restrictionType
		list.IsRestrictionType = true
	}

	err := AddSliceToList(values, list)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (list *List) AddValue(value any) error {
	flagMatch := list.verifyType(value)

	if !flagMatch {
		return errNodeMatchTypeInList()
	}

	node := GenerateNode(value)

	if list.HeadPointer == nil {
		list.HeadPointer = node
		list.Length += 1

		return nil
	}

	last := list.HeadPointer

	for last.Next != nil {
		last = last.Next
	}

	last.Next = node
	list.Length += 1

	return nil
}

func (list *List) verifyType(value any) bool {
	if !list.IsRestrictionType {
		return true
	}

	return list.RestrictionType == reflect.TypeOf(value)
}

func (list *List) String() string {
	if list == nil {
		return StringEmpty
	}

	str := StringEmpty

	str += fmt.Sprintf("Length: %d\n", list.Length)
	str += fmt.Sprintf("IsRestrictionType: %v\n", list.IsRestrictionType)
	str += fmt.Sprintf("RestrictionType: %v\n", list.RestrictionType)
	str += fmt.Sprintf("Content: %v\n", list.ToString())

	return str
}

func (list *List) ToString() string {
	if list.Length == 0 {
		return StringEmpty
	}

	node := list.HeadPointer
	result := StringEmpty

	for {
		if node != nil {
			result += fmt.Sprintf("%v -> ", node.Value)
		}

		node = node.Next

		if node == nil {
			break
		}
	}

	result += "nil"

	return result
}

func AddSliceToList[T any](values *[]T, list *List) error {
	if values == nil || *values == nil || len(*values) == 0 {
		return nil
	}

	tmpList := &List{
		HeadPointer:       nil,
		Length:            0,
		IsRestrictionType: list.IsRestrictionType,
		RestrictionType:   list.RestrictionType,
	}

	for i, v := range *values {
		err := tmpList.AddValue(v)

		if err != nil {
			if err.Error() == errNodeMatchTypeInList().Error() {
				return errInitValuesNotMatchType(i)
			} else {
				return err
			}
		}
	}

	list.HeadPointer = tmpList.HeadPointer
	list.Length += tmpList.Length

	return nil
}

func (list *List) DeleteNode(deleteNode *Node) {
	if deleteNode == nil {
		return
	}

	if list.HeadPointer == deleteNode {
		list.HeadPointer = list.HeadPointer.Next
		list.Length--

		return
	}

	var pre, next *Node

	list.ForEach(func(i int, node *Node, lst *List) bool {
		if node.Next.Equal(deleteNode, true) {
			pre = node
			next = deleteNode.Next
			pre.Next = next
			list.Length--

			return false
		}

		return true
	})
}

func (list *List) ForEach(f func(int, *Node, *List) bool) {
	if list.HeadPointer == nil {
		return
	}

	index := 0
	node := list.HeadPointer

	for node != nil {
		isContinue := f(index, node, list)
		if !isContinue {
			break
		}

		index++
		node = node.Next
	}
}
