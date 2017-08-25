package nestederror

import (
	"errors"
	"testing"
)

func TestErrors_String_CustomSeparator(t *testing.T) {
	var errs Errors

	oldsep := Separator
	Separator = " > "
	defer func() {
		Separator = oldsep
	}()

	errs = []error{errors.New("foo"), errors.New("bar")}
	res := errs.String()

	if res != "foo > bar" {
		t.Fatalf("expected 'foo > bar, got '%v'", res)
	}
}

func TestErrors_String_DefaultSeparator(t *testing.T) {
	var errs Errors

	errs = []error{errors.New("foo"), errors.New("bar")}
	res := errs.String()

	if res != "foo -> bar" {
		t.Fatalf("expected 'foo -> bar', got '%v'", res)
	}
}

func TestErrors_Strings(t *testing.T) {
	var errs Errors

	errs = []error{errors.New("foo"), errors.New("bar")}
	res := errs.Strings()

	if l := len(res); l != 2 {
		t.Fatalf("expected slice with 2 elements, got %v", l)
	}

	if s := res[0]; s != "foo" {
		t.Fatalf("expected 'foo' at index 0, got '%v'", s)
	}

	if s := res[1]; s != "bar" {
		t.Fatalf("expected 'bar' at index 1, got '%v'", s)
	}
}

func TestFlatten(t *testing.T) {
	errs := []error{
		errors.New("1"),
		errors.New("2"),
		errors.New("3"),
		errors.New("4"),
	}

	outer := New(errs[1], errs[0])
	inner := New(errs[3], errs[2])
	err := New(inner, outer)

	res := Flatten(err)

	if l := len(res); l != len(errs) {
		t.Fatalf("expected slice with %v elements, got %v", len(errs), l)
	}

	for i := 0; i < len(errs); i++ {
		if res[i] != errs[i] {
			t.Fatalf("expected error number %v on index %v", i, i)
		}
	}
}

func TestFlatten_ErrIsNil(t *testing.T) {
	res := Flatten(nil)

	if res != nil {
		t.Fatal("expected nil, got non-nil")
	}
}

func TestFlatten_ErrIsNormal(t *testing.T) {
	err := errors.New("foo")

	res := Flatten(err)

	if l := len(res); l != 1 {
		t.Fatalf("expected slice with 1 elements, got %v", l)
	}

	if res[0] != err {
		t.Fatal("slice does not contains original error")
	}
}

func TestFlatten_ErrIsSimpleNested(t *testing.T) {
	inner := errors.New("foo")
	outer := errors.New("bar")
	err := New(inner, outer)

	res := Flatten(err)

	if l := len(res); l != 2 {
		t.Fatalf("expected slice with 2 elements, got %v", l)
	}

	if res[0] != outer {
		t.Fatal("expected outer error on index 0")
	}

	if res[1] != inner {
		t.Fatal("expected inner error on index 1")
	}
}
