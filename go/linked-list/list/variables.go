package list

import (
	"errors"
	"fmt"
)

var (
	errInitNilShouldNil = func() error {
		return errors.New("if you want an empty linked list, please change to nil")
	}

	errInitValuesNotMatchType = func(x int) error {
		return fmt.Errorf("input values in index[%d] type does not match restriction type", x)
	}

	errListIsEmpty = func() error {
		return errors.New("not has node")
	}

	errNodeMatchTypeInList = func() error {
		return errors.New("input not match type in list")
	}
)

const (
	stringEmpty = ""
)
