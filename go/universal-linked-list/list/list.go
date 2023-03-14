package list

import (
	"fmt"
	"reflect"
)

type List struct {
	HeadPointer     *Node
	Length          int
	restrictionType reflect.Type
}

func GenerateEmptyList(restrictionType reflect.Type) *List {
	lst := &List{}
	if restrictionType != nil {
		lst.restrictionType = restrictionType
	}

	return lst
}

func GenerateInitList[T any](values *[]T, restrictionType reflect.Type) (*List, error) {
	lst := &List{}

	if restrictionType != nil {
		lst.restrictionType = restrictionType
	}

	err := lst.AppendSlice(values)
	if err != nil {
		return nil, err
	}

	return lst, nil
}

func (lst *List) AppendSlice(slice any) error {
	if reflect.TypeOf(slice).Kind() != reflect.Ptr {
		return errValueNotIsPtr()
	}

	if reflect.ValueOf(slice).Elem().Kind() == reflect.Invalid {
		return errSliceIsInvalid()
	}

	v := reflect.ValueOf(slice).Elem().Interface()
	vtk := reflect.TypeOf(v).Kind()
	if !(vtk == reflect.Array || vtk == reflect.Slice) {
		return errValueNotArrayOrSlice()
	}

	tmpLst := &List{
		HeadPointer:     nil,
		Length:          0,
		restrictionType: lst.restrictionType,
	}

	sli := reflect.ValueOf(v)
	if sli.Len() == 0 {
		return errLengthEqualZero()
	}

	for i := 0; i < sli.Len(); i++ {
		content := sli.Index(i).Interface()
		err := tmpLst.AppendValue(content)

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

func (lst *List) AppendValue(value any) error {
	if lst == nil {
		return errListIsNil()
	}

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

func (lst *List) verifyType(value any) bool {
	if lst.restrictionType == nil {
		return true
	}

	return lst.restrictionType == reflect.TypeOf(value)
}

func (lst *List) String() string {
	if lst == nil {
		return stringEmpty
	}

	str := stringEmpty

	str += fmt.Sprintf("Length: %d\n", lst.Length)
	str += fmt.Sprintf("RestrictionType: %v\n", lst.restrictionType)
	str += fmt.Sprintf("Content: %v", lst.ContentString())

	return str
}

func (lst *List) ContentString() string {
	if lst.Length == 0 {
		return stringEmpty
	}

	nde := lst.HeadPointer
	result := stringEmpty

	for {
		if nde != nil {
			result += fmt.Sprintf("(%v, %v) -> ", nde.Value, nde.nodeType)
		}

		nde = nde.Next

		if nde == nil {
			break
		}
	}

	result += "nil"

	return result
}

func (lst *List) DeleteNode(deleteNode *Node) {
	if deleteNode == nil {
		return
	}

	if lst.HeadPointer.Equal(deleteNode, true) {
		lst.HeadPointer = lst.HeadPointer.Next
		lst.Length--

		return
	}

	var pre, next *Node

	lst.ForEach(func(i int, nde *Node, lst *List) bool {
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

func (lst *List) ForEach(f func(int, *Node, *List) bool) {
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

func (lst *List) FrontAddValue(value any) error {
	if lst == nil {
		return errListIsNil()
	}

	if !lst.verifyType(value) {
		return errNodeMatchTypeInList()
	}

	nde := GenerateNode(value)
	if lst.HeadPointer == nil {
		lst.HeadPointer = nde
		lst.Length += 1

		return nil
	}

	nde.Next = lst.HeadPointer
	lst.HeadPointer = nde
	lst.Length += 1

	return nil
}

func (lst *List) Reverse() error {
	if lst == nil {
		return errListIsNil()
	}

	if lst.Length == 0 {
		return errListLengthIsNil()
	}

	if lst.Length == 1 {
		return nil
	}

	tmpLst := &List{
		HeadPointer:     lst.HeadPointer,
		Length:          lst.Length,
		restrictionType: lst.restrictionType,
	}

	lst.Length = 0
	lst.HeadPointer = nil

	tmpLst.ForEach(func(i int, nde *Node, _ *List) bool {
		lst.FrontAddValue(nde.Value)

		return true
	})

	return nil
}
