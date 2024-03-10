package errless

import "fmt"

// HandlerFunc defines the signature for an error handler.
type HandlerFunc func(error) error

// zero parameter functions
// --------------------------

// Check checks the error with default error handler.
func Check(err error, handles ...HandlerFunc) {
	if err != nil {
		for _, handle := range handles {
			err = handle(err)
			if err != nil {
				err = handle(err)
			}
		}
		if err != nil {
			panic(err)
		}
	}
}

// Params0 hold function parameters and error.
type Params0 struct {
	err error
}

func CheckW(err error) *Params0 {
	return &Params0{err: err}
}

func (r *Params0) With(handle ...HandlerFunc) {
	Check(r.err, handle...)
}

func (r *Params0) WithMessage(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
}
func (r *Params0) WithWrap(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
}

// one parameter functions
// --------------------------

// Check1 checks the error with default error handler.
func Check1[A any](a A, err error) A {
	if err != nil {
		Check(err)
	}
	return a
}

// Params1 hold function parameters and error.
type Params1[A any] struct {
	paramA A
	err    error
}

// Try1 is an alias for Check1W.
func Try1[A any](a A, err error) *Params1[A] {
	return Check1W(a, err)
}

func Check1W[A any](a A, err error) *Params1[A] {
	return &Params1[A]{paramA: a, err: err}
}

// With applies an error handler to the Result.
func (r *Params1[A]) With(handle ...HandlerFunc) A {
	Check(r.err, handle...)
	return r.paramA
}

// W is an alias for With.
func (r *Params1[A]) W(handle ...HandlerFunc) A {
	return r.With(handle...)
}

func (r *Params1[A]) WithMessage(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
}
func (r *Params1[A]) WithWrap(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
}

// 2 parameter functions
// --------------------------

// Check2 checks the error with default error handler.
func Check2[A, B any](a A, b B, err error) (A, B) {
	if err != nil {
		Check(err)
	}
	return a, b
}

// Params2 hold function parameters and error.
type Params2[A, B any] struct {
	paramA A
	paramB B
	err    error
}

func Check2W[A, B any](a A, b B, err error) *Params2[A, B] {
	return &Params2[A, B]{paramA: a, paramB: b, err: err}
}

// HandleD With applies an error handler to the Result.
func (r *Params2[A, B]) With(handle ...HandlerFunc) (A, B) {
	Check(r.err, handle...)
	return r.paramA, r.paramB
}

func (r *Params2[A, B]) WithMessage(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
}
func (r *Params2[A, B]) WithWrap(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
}

// three parameter functions
// --------------------------

// Check3 checks the error with default error handler.
func Check3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	if err != nil {
		Check(err)
	}
	return a, b, c
}

// Params3 hold function parameters and error.
type Params3[A, B, C any] struct {
	paramA A
	paramB B
	paramC C
	err    error
}

func Check3W[A, B, C any](a A, b B, c C, err error) *Params3[A, B, C] {
	return &Params3[A, B, C]{paramA: a, paramB: b, paramC: c, err: err}
}

// With applies an error handler to the Result.
func (r *Params3[A, B, C]) With(handle ...HandlerFunc) (A, B, C) {
	Check(r.err, handle...)
	return r.paramA, r.paramB, r.paramC
}

func (r *Params3[A, B, C]) WithMessage(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
}
func (r *Params3[A, B, C]) WithWrap(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
} // four parameter functions
// --------------------------

// Check4 checks the error with default error handler.
func Check4[A, B, C, D any](a A, b B, c C, d D, err error) (A, B, C, D) {
	if err != nil {
		Check(err)
	}
	return a, b, c, d
}

// Params4  hold function parameters and error.
type Params4[A, B, C, D any] struct {
	paramA A
	paramB B
	paramC C
	paramD D
	err    error
}

func Check4W[A, B, C, D any](a A, b B, c C, d D, err error) *Params4[A, B, C, D] {
	return &Params4[A, B, C, D]{paramA: a, paramB: b, paramC: c, paramD: d, err: err}
}

// With applies an error handler to the Result.
func (r *Params4[A, B, C, D]) With(handle ...HandlerFunc) (A, B, C, D) {
	Check(r.err, handle...)
	return r.paramA, r.paramB, r.paramC, r.paramD
}

func (r *Params4[A, B, C, D]) WithMessage(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
}
func (r *Params4[A, B, C, D]) WithWrap(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
} // five parameter functions
// --------------------------

// Check5 checks the error with default error handler.
func Check5[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) (A, B, C, D, E) {
	if err != nil {
		Check(err)
	}
	return a, b, c, d, e
}

// Params5  hold function parameters and error.
type Params5[A, B, C, D, E any] struct {
	paramA A
	paramB B
	paramC C
	paramD D
	paramE E
	err    error
}

func Check5W[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) *Params5[A, B, C, D, E] {
	return &Params5[A, B, C, D, E]{paramA: a, paramB: b, paramC: c, paramD: d, paramE: e, err: err}
}

// With applies an error handler to the Result.
func (r *Params5[A, B, C, D, E]) With(handle ...HandlerFunc) (A, B, C, D, E) {
	Check(r.err, handle...)
	return r.paramA, r.paramB, r.paramC, r.paramD, r.paramE
}

func (r *Params5[A, B, C, D, E]) WithMessage(message string) {
	Check(fmt.Errorf(message+"error: %s", r.err))
}
func (r *Params5[A, B, C, D, E]) WithWrap(message string) {
	Check(fmt.Errorf(message+"error: %w", r.err))
} // ReturnErr is set caught error to passed named error
func ReturnErr(namedErr *error) {
	Handle(namedErr, EmptyHandler)
}

// Handle is used to handle to catch panics and handle errors with a custom error handling function.
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

func HandleReturn(onError func(e error)) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			onError(err)
		} else {
			// This was not an error panic; re-panic with the original value.
			panic(r)
		}
	}
}

func EmptyHandler(err error) error {
	return err
}
