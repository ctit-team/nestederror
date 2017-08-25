package nestederror

import (
	"strings"
)

var (
	// Separator is a string to separate between error message, default is ' -> '.
	Separator = " -> "
)

// Flatten flatten all errors into an Errors.
func Flatten(err error) Errors {
	if err == nil {
		return nil
	}

	errs := make([]error, 0, 16)

	return flatten(errs, err)
}

// Errors is a collection of error.
type Errors []error

func (r Errors) String() string {
	return strings.Join(r.Strings(), Separator)
}

// Strings return all error's string.
func (r Errors) Strings() []string {
	l := make([]string, 0, len(r))
	for _, e := range r {
		l = append(l, e.Error())
	}
	return l
}

func flatten(errs []error, err error) []error {
	switch err := err.(type) {
	case *NestedError:
		errs = flatten(errs, err.Outer())
		return flatten(errs, err.Inner())
	default:
		return append(errs, err)
	}
}
