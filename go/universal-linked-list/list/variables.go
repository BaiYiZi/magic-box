package list

import (
	"fmt"
)

var (
	errInitNilShouldNil = func() error {
		return fmt.Errorf("if you want an empty linked list, please call \"GenerateEmptyList()\"")
	}

	errInitValuesNotMatchType = func(x int) error {
		return fmt.Errorf("input values in index[%d] type does not match restriction type", x)
	}

	errNodeMatchTypeInList = func() error {
		return fmt.Errorf("input not match type in list")
	}
)

const (
	StringEmpty = ""
)
