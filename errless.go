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

func Message(message string) HandlerFunc {
	return func(err error) error {
		return fmt.Errorf(message+" - error: %s", err)
	}
}

func Wrap(message string) HandlerFunc {
	return func(err error) error {
		return fmt.Errorf(message+" - error: %s", err)
	}
}

// zero parameter functions
// --------------------------

// Throw checks the error with default error handler.
func Throw(err error, handles ...HandlerFunc) {
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

func Try(err error) *Params0 {
	return &Params0{err: err}
}

func (r *Params0) Err(handle ...HandlerFunc) {
	if !r.skipNextStep {
		Throw(r.err, handle...)
	}
}
func (r *Params0) Or(handle func(error)) {
	handle(r.err)
}
func (r *Params0) If(handle ...IfFunc) *Params0 {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params0) ErrMessage(message string) {
	r.Err(Message(message))
}
func (r *Params0) ErrWrap(message string) {
	r.Err(Wrap(message))
}

// one parameter functions
// --------------------------

// Throw1 checks the error with default error handler.
func Throw1[A any](a A, err error) A {
	Throw(err)
	return a
}

// Params1 hold function parameters and error.
type Params1[A any] struct {
	paramA         A
	err            error
	skipNextHandle bool
}

func Try1[A any](a A, err error) *Params1[A] {
	return &Params1[A]{paramA: a, err: err}
}

// Err applies an error handler to the Result.
func (r *Params1[A]) Err(handle ...HandlerFunc) A {
	if !r.skipNextHandle {
		Throw(r.err, handle...)
	}
	return r.paramA
}

func (r *Params1[A]) Fallback(handle func(error) A) A {
	if r.skipNextHandle {
		Throw(r.err)
	}
	return handle(r.err)
}
func (r *Params1[A]) If(handle ...IfFunc) *Params1[A] {
	r.skipNextHandle = !applyNextStep(handle, r.err, r.skipNextHandle)
	return r
}

func (r *Params1[A]) IfIs(err error) *Params1[A] {
	return r.If(Is(err))
}
func (r *Params1[A]) IfNot(err error) *Params1[A] {
	return r.If(IsNot(err))
}

// E is an alias for Err.
func (r *Params1[A]) E(handle ...HandlerFunc) A {
	return r.Err(handle...)
}

func (r *Params1[A]) ErrMessage(message string) A {
	return r.Err(Message(message))
}
func (r *Params1[A]) ErrWrap(message string) A {
	return r.Err(Wrap(message))
}

// 2 parameter functions
// --------------------------

// Throw2 checks the error with default error handler.
func Throw2[A, B any](a A, b B, err error) (A, B) {
	Throw(err)
	return a, b
}

// Params2 hold function parameters and error.
type Params2[A, B any] struct {
	paramA       A
	paramB       B
	err          error
	skipNextStep bool
}

func Try2[A, B any](a A, b B, err error) *Params2[A, B] {
	return &Params2[A, B]{paramA: a, paramB: b, err: err}
}

// Err  applies an error handler to the Result.
func (r *Params2[A, B]) Err(handle ...HandlerFunc) (A, B) {
	if !r.skipNextStep {
		Throw(r.err, handle...)
	}
	return r.paramA, r.paramB
}

func (r *Params2[A, B]) Fallback(handle func(error) (A, B)) (A, B) {
	if r.skipNextStep {
		Throw(r.err)
	}
	return handle(r.err)
}

func (r *Params2[A, B]) If(handle ...IfFunc) *Params2[A, B] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params2[A, B]) IfIs(err error) *Params2[A, B] {
	return r.If(Is(err))
}
func (r *Params2[A, B]) IfNot(err error) *Params2[A, B] {
	return r.If(IsNot(err))
}

func (r *Params2[A, B]) ErrMessage(message string) (A, B) {
	return r.Err(Message(message))
}
func (r *Params2[A, B]) ErrWrap(message string) (A, B) {
	return r.Err(Wrap(message))
}

// three parameter functions
// --------------------------

// Throw3 checks the error with default error handler.
func Throw3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	Throw(err)
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

func Try3[A, B, C any](a A, b B, c C, err error) *Params3[A, B, C] {
	return &Params3[A, B, C]{paramA: a, paramB: b, paramC: c, err: err}
}

// Err applies an error handler to the Result.
func (r *Params3[A, B, C]) Err(handle ...HandlerFunc) (A, B, C) {
	if !r.skipNextStep {
		Throw(r.err, handle...)
	}
	return r.paramA, r.paramB, r.paramC
}

func (r *Params3[A, B, C]) Fallback(handle func(error) (A, B, C)) (A, B, C) {
	if r.skipNextStep {
		Throw(r.err)
	}
	return handle(r.err)
}

func (r *Params3[A, B, C]) If(handle ...IfFunc) *Params3[A, B, C] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params3[A, B, C]) IfIs(err error) *Params3[A, B, C] {
	return r.If(Is(err))
}
func (r *Params3[A, B, C]) IfNot(err error) *Params3[A, B, C] {
	return r.If(IsNot(err))
}

func (r *Params3[A, B, C]) ErrMessage(message string) (A, B, C) {
	return r.Err(Message(message))
}
func (r *Params3[A, B, C]) ErrWrap(message string) (A, B, C) {
	return r.Err(Wrap(message))
}

// four parameter functions
// --------------------------

// Throw4 checks the error with default error handler.
func Throw4[A, B, C, D any](a A, b B, c C, d D, err error) (A, B, C, D) {
	Throw(err)
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

func Try4[A, B, C, D any](a A, b B, c C, d D, err error) *Params4[A, B, C, D] {
	return &Params4[A, B, C, D]{paramA: a, paramB: b, paramC: c, paramD: d, err: err}
}

// Err applies an error handler to the Result.
func (r *Params4[A, B, C, D]) Err(handle ...HandlerFunc) (A, B, C, D) {
	if !r.skipNextStep {
		Throw(r.err, handle...)
	}
	return r.paramA, r.paramB, r.paramC, r.paramD
}

func (r *Params4[A, B, C, D]) Fallback(handle func(error) (A, B, C, D)) (A, B, C, D) {
	if r.skipNextStep {
		Throw(r.err)
	}
	return handle(r.err)
}

func (r *Params4[A, B, C, D]) If(handle ...IfFunc) *Params4[A, B, C, D] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params4[A, B, C, D]) IfIs(err error) *Params4[A, B, C, D] {
	return r.If(Is(err))
}
func (r *Params4[A, B, C, D]) IfNot(err error) *Params4[A, B, C, D] {
	return r.If(IsNot(err))
}

func (r *Params4[A, B, C, D]) ErrMessage(message string) (A, B, C, D) {
	return r.Err(Message(message))
}
func (r *Params4[A, B, C, D]) ErrWrap(message string) (A, B, C, D) {
	return r.Err(Wrap(message))
}

// five parameter functions
// --------------------------

// Throw5 checks the error with default error handler.
func Throw5[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) (A, B, C, D, E) {
	Throw(err)
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

func Try5[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) *Params5[A, B, C, D, E] {
	return &Params5[A, B, C, D, E]{paramA: a, paramB: b, paramC: c, paramD: d, paramE: e, err: err}
}

// Err applies an error handler to the Result.
func (r *Params5[A, B, C, D, E]) Err(handle ...HandlerFunc) (A, B, C, D, E) {
	if !r.skipNextStep {
		Throw(r.err, handle...)
	}
	return r.paramA, r.paramB, r.paramC, r.paramD, r.paramE
}
func (r *Params5[A, B, C, D, E]) Fallback(handle func(error) (A, B, C, D, E)) (A, B, C, D, E) {
	if r.skipNextStep {
		Throw(r.err)
	}
	if r.skipNextStep {
		Throw(r.err)
	}
	return handle(r.err)
}

func (r *Params5[A, B, C, D, E]) If(handle ...IfFunc) *Params5[A, B, C, D, E] {
	r.skipNextStep = !applyNextStep(handle, r.err, r.skipNextStep)
	return r
}

func (r *Params5[A, B, C, D, E]) IfIs(err error) *Params5[A, B, C, D, E] {
	return r.If(Is(err))
}
func (r *Params5[A, B, C, D, E]) IfNot(err error) *Params5[A, B, C, D, E] {
	return r.If(IsNot(err))
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
	return r.Err(Message(message))
}
func (r *Params5[A, B, C, D, E]) ErrWrap(message string) (A, B, C, D, E) {
	return r.Err(Wrap(message))
}

// HandleErr is set caught error to passed named error
func HandleErr(namedErr *error) {
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

func Catch(onError func(e error)) {
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
