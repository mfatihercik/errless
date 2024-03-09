package errless

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Check1[A any](a A, err error) A {
	if err != nil {
		Check(err)
	}
	return a
}
func Check2[A, B any](a A, b B, err error) (A, B) {
	if err != nil {
		Check(err)
	}
	return a, b
}

func Check3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	if err != nil {
		Check(err)
	}
	return a, b, c
}

func Check4[A, B, C, D any](a A, b B, c C, d D, err error) (A, B, C, D) {
	if err != nil {
		Check(err)
	}
	return a, b, c, d
}

func Check5[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) (A, B, C, D, E) {
	if err != nil {
		Check(err)
	}
	return a, b, c, d, e
}

// handleError is a deferred function that takes a custom error handling function as a parameter.
func Handle(namedErr *error, onError func(error) error) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			e := onError(err) // Use the provided custom error handling logic.
			if namedErr != nil {
				*namedErr = e
			}
		} else {
			// This was not an error panic; re-panic with the original value.
			panic(r)
		}
	}
}

func EmptyHandler(err error) error {
	return err
}
