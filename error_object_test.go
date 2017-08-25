package nestederror

import (
	"errors"
	"testing"
)

func TestNestedError(t *testing.T) {
	inner := errors.New("foo")
	outer := errors.New("bar")

	e := New(inner, outer)

	if e := e.Error(); e != "bar -> foo" {
		t.Fatalf("expected 'bar -> foo' from String, got '%v'", e)
	}

	if e.Inner() == nil {
		t.Fatal("expected non-nil from Inner")
	}

	if e.Outer() == nil {
		t.Fatal("expected non-nil from Outer")
	}
}

func TestNestedError_InnerIsNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.FailNow()
		}
	}()

	New(nil, "foo")
}

func TestNestedError_OuterIsNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.FailNow()
		}
	}()

	New(errors.New("foo"), nil)
}

func TestNestedError_OuterIsString(t *testing.T) {
	e := New(errors.New("foo"), "bar %v", 7)

	if e := e.Outer().Error(); e != "bar 7" {
		t.Fatalf("expect 'bar 7' from outer, got %v", e)
	}
}
