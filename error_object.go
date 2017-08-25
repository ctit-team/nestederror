package nestederror

import (
	"fmt"
)

// NestedError is a type that holding multiple errors.
type NestedError struct {
	outer error
	inner error
}

// New create a new NestedError object.
//
// outer can be either error or string. If it is string,
// it will be convert to error with fmt.Errorf.
func New(inner error, outer interface{}, args ...interface{}) *NestedError {
	// check arguments
	if inner == nil {
		panic("inner is nil")
	}

	if outer == nil {
		panic("outer is nil")
	}

	e := &NestedError{
		inner: inner,
	}

	// get the outer error
	switch outer := outer.(type) {
	case error:
		e.outer = outer
	case string:
		e.outer = fmt.Errorf(outer, args...)
	default:
		panic("outer is not either error or string")
	}

	return e
}

func (e *NestedError) Error() string {
	return Flatten(e).String()
}

// Inner return the inner error.
func (e *NestedError) Inner() error {
	return e.inner
}

// Outer return the outer error.
func (e *NestedError) Outer() error {
	return e.outer
}
