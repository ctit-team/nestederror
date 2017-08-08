package nestederror

// Flatten flatten all nested errors into a slice.
func Flatten(err error) []error {
	if err == nil {
		return nil
	}

	errs := make([]error, 0, 16)

	return flatten(errs, err)
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
