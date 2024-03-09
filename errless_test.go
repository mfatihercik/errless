//go:build test

package errless_test

import (
	"errors"
	"testing"

	"github.com/mfatihercik/errless"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {

	t.Run("function fail cases", func(t *testing.T) {
		t.Run("function with error return value zero parameter", func(t *testing.T) {
			expectedError := errors.New("test error")
			funcReturnError := func() (err error) {
				zeroParameterFunc := func() error {
					return expectedError
				}
				defer errless.Handle(&err, errless.EmptyHandler)
				errless.Check(zeroParameterFunc())
				return err
			}

			err := funcReturnError()
			assert.Equal(t, expectedError, err)
		})

		t.Run("void function with zero parameter", func(t *testing.T) {
			expectedError := errors.New("test error")
			voidFunc := func() {
				zeroParameterFunc := func() error {
					return expectedError
				}
				errorHandle := func(err error) error {
					assert.Equal(t, expectedError, err)
					return err
				}
				defer errless.Handle(nil, errorHandle)
				errless.Check(zeroParameterFunc())
			}

			voidFunc()
		})

	})

	t.Run("function success case", func(t *testing.T) {

		t.Run("function with error return value zero parameter", func(t *testing.T) {
			funcReturnError := func() (err error) {
				zeroParameterFunc := func() error {
					// This is a success case, so we should return nil
					return nil
				}
				defer errless.Handle(&err, errless.EmptyHandler)
				errless.Check(zeroParameterFunc())
				return err
			}

			err := funcReturnError()
			assert.Nil(t, err)
		})

	})
}

func TestCheck1(t *testing.T) {

	t.Run("function fail cases", func(t *testing.T) {
		t.Run("function with error return value and  1 parameter internal call", func(t *testing.T) {
			expectedError := errors.New("test error")
			funcReturnError := func(val int) (err error) {

				innerFunc := func(val int) (int, error) {
					return 0, expectedError
				}
				defer errless.Handle(&err, errless.EmptyHandler)
				errless.Check1(innerFunc(val))
				return err
			}

			err := funcReturnError(10)
			assert.Equal(t, expectedError, err)
		})

		t.Run("void function with zero parameter internal call", func(t *testing.T) {
			expectedError := errors.New("test error")
			voidFunc := func(val int) {

				innerFunc := func(val int) (int, error) {
					return 0, expectedError
				}
				defer errless.Handle(nil, errless.EmptyHandler)
				errless.Check1(innerFunc(val))
			}

			voidFunc(10)
		})

	})

	t.Run("function success case", func(t *testing.T) {

		t.Run("function with error return value and 1 parameter internal call", func(t *testing.T) {
			funcReturnError := func(val int) (err error) {

				innerFunc := func(val int) (int, error) {
					return val, nil
				}
				defer errless.Handle(&err, errless.EmptyHandler)
				errless.Check1(innerFunc(val))
				return err
			}

			err := funcReturnError(10)
			assert.Nil(t, err)
		})

	})
}

func TestCheck2(t *testing.T) {

	t.Run("function fail cases", func(t *testing.T) {
		t.Run("function with error return value and  2 parameter internal call", func(t *testing.T) {
			expectedError := errors.New("test error")
			funcReturnError := func(val int) (err error) {

				innerFunc := func(val int) (int, int, error) {
					return 0, 0, expectedError
				}
				defer errless.Handle(&err, errless.EmptyHandler)
				errless.Check2(innerFunc(val))
				return err
			}

			err := funcReturnError(10)
			assert.Equal(t, expectedError, err)
		})

		t.Run("void function with zero parameter internal call", func(t *testing.T) {
			expectedError := errors.New("test error")
			voidFunc := func(val int) {

				innerFunc := func(val int) (int, int, error) {
					return 0, 0, expectedError
				}
				defer errless.Handle(nil, errless.EmptyHandler)
				errless.Check2(innerFunc(val))
			}

			voidFunc(10)
		})

	})

	t.Run("function success case", func(t *testing.T) {

		t.Run("function with error return value and 1 parameter internal call", func(t *testing.T) {
			funcReturnError := func(val int) (err error) {

				innerFunc := func(val int) (int, int, error) {
					return val, val, nil
				}
				defer errless.Handle(&err, errless.EmptyHandler)
				errless.Check2(innerFunc(val))
				return err
			}

			err := funcReturnError(10)
			assert.Nil(t, err)
		})

	})
}

func TestCheck3(t *testing.T) {

	t.Run("function fail cases", func(t *testing.T) {
		t.Run("function with error return value and  2 parameter internal call", func(t *testing.T) {
			expectedError := errors.New("test error")
			funcReturnError := func(val int) (err error) {

				innerFunc := func(val int) (int, int, int, error) {
					return 0, 0, 0, expectedError
				}
				defer errless.Handle(&err, errless.EmptyHandler)
				errless.Check3(innerFunc(val))
				return err
			}

			err := funcReturnError(10)
			assert.Equal(t, expectedError, err)
		})

		t.Run("void function with zero parameter internal call", func(t *testing.T) {
			expectedError := errors.New("test error")
			voidFunc := func(val int) {

				innerFunc := func(val int) (int, int, int, error) {
					return 0, 0, 0, expectedError
				}
				defer errless.Handle(nil, errless.EmptyHandler)
				errless.Check3(innerFunc(val))
			}

			voidFunc(10)
		})

	})

	t.Run("function success case", func(t *testing.T) {

		t.Run("function with error return value and 1 parameter internal call", func(t *testing.T) {
			funcReturnError := func(val int) (err error) {

				innerFunc := func(val int) (int, int, int, error) {
					return 0, 0, 0, nil
				}
				defer errless.Handle(&err, errless.EmptyHandler)
				errless.Check3(innerFunc(val))
				return err
			}

			err := funcReturnError(10)
			assert.Nil(t, err)
		})

	})
}
