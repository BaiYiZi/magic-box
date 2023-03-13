package list

import (
	"fmt"
	"reflect"
)

type list struct {
	HeadPointer     *node
	Length          int
	restrictionType reflect.Type
}

func GenerateEmptyList(restrictionType reflect.Type) *list {
	lst := &list{}
	if restrictionType != nil {
		lst.restrictionType = restrictionType
	}

	return lst
}

func GenerateInitList[T any](values *[]T, restrictionType reflect.Type) (*list, error) {
	lst := &list{}

	if restrictionType != nil {
		lst.restrictionType = restrictionType
	}

	err := AddSliceToList(values, lst)
	if err != nil {
		return lst, err
	}

	return lst, nil
}

func AddSliceToList[T any](values *[]T, lst *list) error {
	if values == nil || *values == nil || len(*values) == 0 {
		return errInitNilShouldNil()
	}

	tmpLst := &list{
		HeadPointer:     nil,
		Length:          0,
		restrictionType: lst.restrictionType,
	}

	for i, v := range *values {
		err := tmpLst.AddValue(v)

		if err != nil {
			if err.Error() == errNodeMatchTypeInList().Error() {
				return errInitValuesNotMatchType(i)
			} else {
				return err
			}
		}
	}

	lst.HeadPointer = tmpLst.HeadPointer
	lst.Length += tmpLst.Length

	return nil
}

func (lst *list) AddValue(value any) error {
	if !lst.verifyType(value) {
		return errNodeMatchTypeInList()
	}

	nde := GenerateNode(value)

	if lst.HeadPointer == nil {
		lst.HeadPointer = nde
		lst.Length += 1

		return nil
	}

	last := lst.HeadPointer

	for last.Next != nil {
		last = last.Next
	}

	last.Next = nde
	lst.Length += 1

	return nil
}

func (lst *list) verifyType(value any) bool {
	if lst.restrictionType == nil {
		return true
	}

	return lst.restrictionType == reflect.TypeOf(value)
}

func (lst *list) String() string {
	if lst == nil {
		return StringEmpty
	}

	str := StringEmpty

	str += fmt.Sprintf("Length: %d\n", lst.Length)
	str += fmt.Sprintf("RestrictionType: %v\n", lst.restrictionType)
	str += fmt.Sprintf("Content: %v\n", lst.ContentString())

	return str
}

func (lst *list) ContentString() string {
	if lst.Length == 0 {
		return StringEmpty
	}

	nde := lst.HeadPointer
	result := StringEmpty

	for {
		if nde != nil {
			result += fmt.Sprintf("%v -> ", nde.Value)
		}

		nde = nde.Next

		if nde == nil {
			break
		}
	}

	result += "nil"

	return result
}

func (lst *list) DeleteNode(deleteNode *node) {
	if deleteNode == nil {
		return
	}

	if lst.HeadPointer == deleteNode {
		lst.HeadPointer = lst.HeadPointer.Next
		lst.Length--

		return
	}

	var pre, next *node

	lst.ForEach(func(i int, nde *node, lst *list) bool {
		if nde.Next.Equal(deleteNode, true) {
			pre = nde
			next = deleteNode.Next
			pre.Next = next
			lst.Length--

			return false
		}

		return true
	})
}

func (lst *list) ForEach(f func(int, *node, *list) bool) {
	if lst.HeadPointer == nil {
		return
	}

	index := 0
	nde := lst.HeadPointer

	for nde != nil {
		isContinue := f(index, nde, lst)
		if !isContinue {
			break
		}

		index++
		nde = nde.Next
	}
}
