package errless

import (
	"errors"
	"fmt"
	"strings"
)

// HandlerFunc defines the signature for an error handler.
type HandlerFunc func(error) error

type IfFunc func(error) bool

func Is(target error) IfFunc {
	return func(err error) bool {
		return errors.Is(err, target)
	}

}

func IsNot(target error) IfFunc {
	return func(err error) bool {
		return !errors.Is(err, target)
	}

}

func Contains(str string) IfFunc {
	return func(err error) bool {
		return strings.Contains(err.Error(), str)
	}

}

func messageHandler(message string) HandlerFunc {
	return func(err error) error {
		return fmt.Errorf(message+" - error: %s", err)
	}
}

func wrapHandler(message string) HandlerFunc {
	return func(err error) error {
		return fmt.Errorf(message+" - error: %s", err)
	}
}

// zero parameter functions
// --------------------------

// Check checks the error with default error handler.
func Check(err error, handles ...HandlerFunc) {
	if err != nil {
		for _, handle := range handles {
			err = handle(err)
			if err == nil {
				// get nil error no need to call the rest of the handler
				break
			}
		}
		if err != nil {
			panic(err)
		}
	}
}

// Params0 hold function parameters and error.
type Params0 struct {
	err          error
	skipNextStep bool
}

func CheckW(err error) *Params0 {
	return &Params0{err: err}
}

func (r *Params0) Err(handle ...HandlerFunc) {
	if !r.skipNextStep {
		Check(r.err, handle...)
	}
}
func (r *Params0) Ok(handle func(error)) {
	handle(r.err)
}
func (r *Params0) If(handle ...IfFunc) *Params0 {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params0) ErrMessage(message string) {
	r.Err(messageHandler(message))
}
func (r *Params0) ErrWrap(message string) {
	r.Err(wrapHandler(message))
}

// one parameter functions
// --------------------------

// Check1 checks the error with default error handler.
func Check1[A any](a A, err error) A {
	Check(err)
	return a
}

// Params1 hold function parameters and error.
type Params1[A any] struct {
	paramA       A
	err          error
	skipNextStep bool
}

// Try1 is an alias for Check1W.
func Try1[A any](a A, err error) *Params1[A] {
	return Check1W(a, err)
}

func Check1W[A any](a A, err error) *Params1[A] {
	return &Params1[A]{paramA: a, err: err}
}

// Err applies an error handler to the Result.
func (r *Params1[A]) Err(handle ...HandlerFunc) A {
	if !r.skipNextStep {
		Check(r.err, handle...)
	}
	return r.paramA
}

func (r *Params1[A]) Ok(handle func(error) A) A {
	return handle(r.err)
}
func (r *Params1[A]) If(handle ...IfFunc) *Params1[A] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

// W is an alias for Err.
func (r *Params1[A]) W(handle ...HandlerFunc) A {
	return r.Err(handle...)
}

func (r *Params1[A]) ErrMessage(message string) A {
	return r.Err(messageHandler(message))
}
func (r *Params1[A]) ErrWrap(message string) A {
	return r.Err(wrapHandler(message))
}

// 2 parameter functions
// --------------------------

// Check2 checks the error with default error handler.
func Check2[A, B any](a A, b B, err error) (A, B) {
	Check(err)
	return a, b
}

// Params2 hold function parameters and error.
type Params2[A, B any] struct {
	paramA       A
	paramB       B
	err          error
	skipNextStep bool
}

func Check2W[A, B any](a A, b B, err error) *Params2[A, B] {
	return &Params2[A, B]{paramA: a, paramB: b, err: err}
}

// Err  applies an error handler to the Result.
func (r *Params2[A, B]) Err(handle ...HandlerFunc) (A, B) {
	if !r.skipNextStep {
		Check(r.err, handle...)
	}
	return r.paramA, r.paramB
}

func (r *Params2[A, B]) Ok(handle func(error) (A, B)) (A, B) {
	return handle(r.err)
}

func (r *Params2[A, B]) If(handle ...IfFunc) *Params2[A, B] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params2[A, B]) ErrMessage(message string) (A, B) {
	return r.Err(messageHandler(message))
}
func (r *Params2[A, B]) ErrWrap(message string) (A, B) {
	return r.Err(wrapHandler(message))
}

// three parameter functions
// --------------------------

// Check3 checks the error with default error handler.
func Check3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	Check(err)
	return a, b, c
}

// Params3 hold function parameters and error.
type Params3[A, B, C any] struct {
	paramA       A
	paramB       B
	paramC       C
	err          error
	skipNextStep bool
}

func Check3W[A, B, C any](a A, b B, c C, err error) *Params3[A, B, C] {
	return &Params3[A, B, C]{paramA: a, paramB: b, paramC: c, err: err}
}

// Err applies an error handler to the Result.
func (r *Params3[A, B, C]) Err(handle ...HandlerFunc) (A, B, C) {
	if !r.skipNextStep {
		Check(r.err, handle...)
	}
	return r.paramA, r.paramB, r.paramC
}

func (r *Params3[A, B, C]) Ok(handle func(error) (A, B, C)) (A, B, C) {
	return handle(r.err)
}

func (r *Params3[A, B, C]) If(handle ...IfFunc) *Params3[A, B, C] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params3[A, B, C]) ErrMessage(message string) (A, B, C) {
	return r.Err(messageHandler(message))
}
func (r *Params3[A, B, C]) ErrWrap(message string) (A, B, C) {
	return r.Err(wrapHandler(message))
}

// four parameter functions
// --------------------------

// Check4 checks the error with default error handler.
func Check4[A, B, C, D any](a A, b B, c C, d D, err error) (A, B, C, D) {
	Check(err)
	return a, b, c, d
}

// Params4  hold function parameters and error.
type Params4[A, B, C, D any] struct {
	paramA       A
	paramB       B
	paramC       C
	paramD       D
	err          error
	skipNextStep bool
}

func Check4W[A, B, C, D any](a A, b B, c C, d D, err error) *Params4[A, B, C, D] {
	return &Params4[A, B, C, D]{paramA: a, paramB: b, paramC: c, paramD: d, err: err}
}

// Err applies an error handler to the Result.
func (r *Params4[A, B, C, D]) Err(handle ...HandlerFunc) (A, B, C, D) {
	if !r.skipNextStep {
		Check(r.err, handle...)
	}
	return r.paramA, r.paramB, r.paramC, r.paramD
}

func (r *Params4[A, B, C, D]) Ok(handle func(error) (A, B, C, D)) (A, B, C, D) {
	return handle(r.err)
}

func (r *Params4[A, B, C, D]) If(handle ...IfFunc) *Params4[A, B, C, D] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params4[A, B, C, D]) ErrMessage(message string) (A, B, C, D) {
	return r.Err(messageHandler(message))
}
func (r *Params4[A, B, C, D]) ErrWrap(message string) (A, B, C, D) {
	return r.Err(wrapHandler(message))
}

// five parameter functions
// --------------------------

// Check5 checks the error with default error handler.
func Check5[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) (A, B, C, D, E) {
	Check(err)
	return a, b, c, d, e
}

// Params5  hold function parameters and error.
type Params5[A, B, C, D, E any] struct {
	paramA       A
	paramB       B
	paramC       C
	paramD       D
	paramE       E
	err          error
	skipNextStep bool
}

func Check5W[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) *Params5[A, B, C, D, E] {
	return &Params5[A, B, C, D, E]{paramA: a, paramB: b, paramC: c, paramD: d, paramE: e, err: err}
}

// Err applies an error handler to the Result.
func (r *Params5[A, B, C, D, E]) Err(handle ...HandlerFunc) (A, B, C, D, E) {
	if !r.skipNextStep {
		Check(r.err, handle...)
	}
	return r.paramA, r.paramB, r.paramC, r.paramD, r.paramE
}
func (r *Params5[A, B, C, D, E]) Ok(handle func(error) (A, B, C, D, E)) (A, B, C, D, E) {
	return handle(r.err)
}

func (r *Params5[A, B, C, D, E]) If(handle ...IfFunc) *Params5[A, B, C, D, E] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func applyNextStep(handle []IfFunc, err error, skipNextStep bool) bool {

	applyNext := !skipNextStep
	if err == nil {
		return applyNext
	}
	apply := false
	for _, filter := range handle {
		if filter(err) {
			apply = true
			break
		}
	}
	return applyNext && apply
}

func (r *Params5[A, B, C, D, E]) ErrMessage(message string) (A, B, C, D, E) {
	return r.Err(messageHandler(message))
}
func (r *Params5[A, B, C, D, E]) ErrWrap(message string) (A, B, C, D, E) {
	return r.Err(wrapHandler(message))
}

// ReturnErr is set caught error to passed named error
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
