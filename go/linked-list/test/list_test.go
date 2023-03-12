package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/BaiYiZi/magic-box/go/linked-list/list"
)

func TestGenerateList(t *testing.T) {
	{
		type S struct{}
		values := []any{20021011, 1.83, "140", S{}, struct{ Name string }{"BaiYiZi"}}

		lst, err := list.GenerateList(&values, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
		}

		fmt.Println(lst)
	}

	fmt.Println("--------------------")

	{
		values := []int{1, 2, 3, 4, 5}
		lst, err := list.GenerateList(&values, reflect.TypeOf(values[0]))

		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
		}

		fmt.Println(lst)
	}

	fmt.Println("--------------------")

	{
		values := []int{110, 20, 30, 40, 50}
		values = nil
		arg := &values
		arg = nil

		lst, err := list.GenerateList(arg, nil)

		if err != nil {
			fmt.Println(fmt.Errorf("error: %s", err))
		}

		fmt.Println(lst)
	}
}