package list

import (
	"fmt"
)

var (
	errValueNotIsPtr = func() error {
		return fmt.Errorf("value not is ptr")
	}

	errSliceIsInvalid = func() error {
		return fmt.Errorf("slice is nil")
	}

	errValueNotArrayOrSlice = func() error {
		return fmt.Errorf("not is slice or array")
	}

	errLengthEqualZero = func() error {
		return fmt.Errorf("length = 0 with array or slice ")
	}

	errInitValuesNotMatchType = func(x int) error {
		return fmt.Errorf("input values in index[%d] type does not match restriction type", x)
	}

	errNodeMatchTypeInList = func() error {
		return fmt.Errorf("input not match type in list")
	}
)

const (
	stringEmpty = ""
)
