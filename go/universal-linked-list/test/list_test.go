package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/BaiYiZi/magic-box/go/linked-list/list"
)

func TestGenerateInitList(t *testing.T) {
	{
		type S struct{}
		values := []any{20021011, 1.83, "140", S{}, struct{ Name string }{"BaiYiZi"}}

		lst, err := list.GenerateInitList(&values, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
		}

		fmt.Println(lst)
	}

	fmt.Println("--------------------")

	{
		values := []int{1, 2, 3, 4, 5}
		lst, err := list.GenerateInitList(&values, reflect.TypeOf(values[0]))

		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
		}

		fmt.Println(lst)
	}

	fmt.Println("--------------------")

	{
		values := []int{110, 20, 30, 40, 50}
		values = nil
		slice := &values
		slice = nil

		lst, err := list.GenerateInitList(slice, nil)

		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
		}

		fmt.Println(lst)
	}
}

func TestGenerateEmptyList(t *testing.T) {
	{
		lst := list.GenerateEmptyList(nil)
		fmt.Println(lst)

		lst.AddSliceToList(&[]any{1, "2", 3.4})
		fmt.Println(lst)
	}

	fmt.Println("--------------------")

	{
		lst := list.GenerateEmptyList(reflect.TypeOf(0))
		fmt.Println(lst)

		lst.AddSliceToList(&[]any{1, 2, 3, 4})
		fmt.Println(lst)
	}
}

func TestAddSliceToList(t *testing.T) {
	lst := list.GenerateEmptyList(nil)
	fmt.Println(lst)

	arr := []int{}
	arr = nil
	values := &arr
	values = nil
	err := lst.AddSliceToList(values)

	fmt.Println(err)
}

func TestEqual(t *testing.T) {
	nodeA := list.GenerateNode(10)
	nodeB := list.GenerateNode(10)
	judgeAddr := true

	fmt.Println(nodeA.Equal(nodeB, judgeAddr))
	fmt.Println(nodeA.Equal(nodeA, judgeAddr))

	fmt.Println(nodeA.Equal(nodeB, false))
}

func TestDeleteNode(t *testing.T) {
	type S struct{}
	values := []any{20021011, 1.83, "140", S{}, struct{ Name string }{"BaiYiZi"}}

	lst, err := list.GenerateInitList(&values, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err))
	}

	fmt.Println(lst)

	node := lst.HeadPointer
	lst.DeleteNode(node)
	fmt.Println(lst)

	node = lst.HeadPointer.Next.Next.Next
	lst.DeleteNode(node)
	fmt.Println(lst)
}
