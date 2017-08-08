package nestederror

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

// Strings return all error's string.
func (s Errors) Strings() []string {
	l := make([]string, 0, len(s))
	for _, e := range s {
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
