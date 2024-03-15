//go:build test

package errless_test

import (
	"errors"
	"fmt"
	"testing"

	e "github.com/mfatihercik/errless"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name              string
	expectedErrorText string
	subErrorText      string
	errorHandle       func(err error) error
	withCustomHandler func(err error) error
	withMessage       string
}

func TestCheckFunctions(t *testing.T) {

	testcases := []TestCase{
		{name: "default handler should return error",
			expectedErrorText: "test error", // error case
			errorHandle: func(err error) error {
				assert.ErrorContains(t, err, "test error")
				return err
			},
			withCustomHandler: nil,
		},

		{name: "default handler should return success",
			expectedErrorText: "", // success case
			errorHandle: func(err error) error {
				t.Fatal("shouldn't be called")
				return err
			},
			withCustomHandler: nil,
		},
		{name: "custom handler: should return customized error",
			expectedErrorText: "custom error", // error in the custom handler
			errorHandle: func(err error) error {
				assert.ErrorContains(t, err, "custom error")
				return err
			},
			withCustomHandler: func(err error) error {
				return fmt.Errorf("custom error :%w ", err)
			},
		},
		{name: "Filter: should return customized error",
			expectedErrorText: "custom error", // error in the custom handler
			errorHandle: func(err error) error {
				assert.ErrorContains(t, err, "custom error")
				return err
			},
			withCustomHandler: func(err error) error {
				return fmt.Errorf("custom error :%w ", err)
			},
			subErrorText: "custom",
		},
		{name: "custom handler: should return success",
			expectedErrorText: "", // success case
			errorHandle: func(err error) error {
				t.Helper()
				t.Fatal("shouldn't be called error handler")
				return err
			},
			withCustomHandler: func(err error) error {
				t.Fatal("shouldn't be called custom handler")
				return err
			},
		},
		{name: "with Message: should return customized error",
			expectedErrorText: "custom message",
			errorHandle: func(err error) error {
				assert.ErrorContains(t, err, "custom message")
				return err
			},
			withMessage: "custom message",
		},
		{name: "with Message: should return success",
			expectedErrorText: "", // success case
			errorHandle: func(err error) error {
				t.Fatal("shouldn't be called")
				return err
			},
			withMessage: "custom message",
		},
	}
	for _, tt := range testcases {

		// zero return function
		functionWithZeroReturn(t, tt)
		functionWithOneReturn(t, tt)
		functionWithTwoReturn(t, tt)
		functionWithThreeReturn(t, tt)
		functionWithFourReturn(t, tt)
		functionWithFiveReturn(t, tt)

		// function with one return value and error

	}

}

func functionWithZeroReturn(t *testing.T, tt TestCase) {
	t.Run(tt.name+" functionWithZeroReturn", func(t *testing.T) {
		t.Helper()
		zeroParameterFunc := func() error {
			if tt.expectedErrorText == "" {
				return nil
			}
			return errors.New(tt.expectedErrorText)
		}
		funcWithErrorReturn := func() (err error) {
			defer e.Handle(&err, tt.errorHandle)
			// if custom handler is not nil, use it
			if tt.withCustomHandler != nil {
				e.Try(zeroParameterFunc()).If(e.Contains(tt.subErrorText)).Err(tt.withCustomHandler)
			} else if tt.withMessage != "" {
				e.Try(zeroParameterFunc()).ErrMessage(tt.withMessage)
			} else {
				e.Throw(zeroParameterFunc())
			}
			return err
		}

		funcWithOk := func() (err error) {
			defer e.Handle(&err, tt.errorHandle)
			// if custom handler is not nil, use it
			e.Try(zeroParameterFunc()).Or(func(err error) {

			})

			return nil
		}

		// call function
		err := funcWithErrorReturn()
		if tt.expectedErrorText == "" {
			assert.Nil(t, err)
		} else {
			assert.ErrorContains(t, err, tt.expectedErrorText)
		}

		err = funcWithOk()
		assert.Nil(t, err)
	})
}

func functionWithOneReturn(t *testing.T, tt TestCase) {
	t.Run(tt.name+" functionWithOneReturn", func(t *testing.T) {
		multiplyByTwo := func(a int) (int, error) {
			if tt.expectedErrorText == "" {
				return a * 2, nil
			}
			return 0, errors.New(tt.expectedErrorText)
		}
		funcWithErrorReturn := func(a int) (rez int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			// if custom handler is not nil, use it

			if tt.withCustomHandler != nil {
				return e.Try1(multiplyByTwo(a)).If(e.Contains(tt.subErrorText)).Err(tt.withCustomHandler), nil
			} else {
				return e.Throw1(multiplyByTwo(a)), nil
			}
		}

		// call function
		res, err := funcWithErrorReturn(2)
		if tt.expectedErrorText == "" {
			assert.Nil(t, err)
			assert.Equal(t, 4, res)
		} else {
			assert.ErrorContains(t, err, tt.expectedErrorText)
			assert.Equal(t, 0, res)
		}

		// check ok case

		funcWithOk := func(a int) (rez int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			res := e.Try1(multiplyByTwo(a)).Fallback(func(err error) int {
				return a * 2
			})
			return res, nil
		}
		res, err = funcWithOk(2)
		assert.Nil(t, err)
		assert.Equal(t, 4, res)

	})

}

func functionWithTwoReturn(t *testing.T, tt TestCase) {
	t.Run(tt.name+" functionWithTwoReturn", func(t *testing.T) {
		multiplyByTwo := func(a, b int) (int, int, error) {
			if tt.expectedErrorText == "" {
				return a * 2, b * 2, nil
			}
			return 0, 0, errors.New(tt.expectedErrorText)
		}
		funcWithErrorReturn := func(a, b int) (res1, res2 int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			// if custom handler is not nil, use it

			if tt.withCustomHandler != nil {
				res1, res2 = e.Try2(multiplyByTwo(a, b)).If(e.Contains(tt.subErrorText)).Err(tt.withCustomHandler)
				return res1, res2, nil
			} else if tt.withMessage != "" {
				res1, res2 = e.Try2(multiplyByTwo(a, b)).ErrMessage(tt.withMessage)
				return res1, res2, nil
			} else {
				res1, res2 = e.Throw2(multiplyByTwo(a, b))
				return res1, res2, nil
			}
		}

		// call function
		res1, res2, err := funcWithErrorReturn(2, 3)
		if tt.expectedErrorText == "" {
			assert.Nil(t, err)
			assert.Equal(t, 4, res1)
			assert.Equal(t, 6, res2)
		} else {
			assert.ErrorContains(t, err, tt.expectedErrorText)
			assert.Equal(t, 0, res1)
			assert.Equal(t, 0, res2)
		}

		funcWithOk := func(a, b int) (res1, res2 int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			res1, res2 = e.Try2(multiplyByTwo(a, b)).Fallback(func(err error) (int, int) {
				return a * 2, b * 2
			})
			return res1, res2, nil
		}
		res1, res2, err = funcWithOk(2, 3)
		assert.Nil(t, err)
		assert.Equal(t, 4, res1)
		assert.Equal(t, 6, res2)
	})

}

func functionWithThreeReturn(t *testing.T, tt TestCase) {
	t.Run(tt.name+" functionWithThreeReturn", func(t *testing.T) {
		multiplyByTwo := func(a, b, c int) (int, int, int, error) {
			if tt.expectedErrorText == "" {
				return a * 2, b * 2, c * 2, nil
			}
			return 0, 0, 0, errors.New(tt.expectedErrorText)
		}
		funcWithErrorReturn := func(a, b, c int) (res1, res2, res3 int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			// if custom handler is not nil, use it

			if tt.withCustomHandler != nil {
				res1, res2, res3 = e.Try3(multiplyByTwo(a, b, c)).If(e.Contains(tt.subErrorText)).Err(tt.withCustomHandler)
				return res1, res2, res3, nil
			} else if tt.withMessage != "" {
				res1, res2, res3 = e.Try3(multiplyByTwo(a, b, c)).ErrMessage(tt.withMessage)
				return res1, res2, res3, nil
			} else {
				res1, res2, res3 = e.Throw3(multiplyByTwo(a, b, c))
				return res1, res2, res3, nil
			}
		}

		// call function
		res1, res2, res3, err := funcWithErrorReturn(2, 3, 4)
		if tt.expectedErrorText == "" {
			assert.Nil(t, err)
			assert.Equal(t, 4, res1)
			assert.Equal(t, 6, res2)
			assert.Equal(t, 8, res3)
		} else {
			assert.ErrorContains(t, err, tt.expectedErrorText)
			assert.Equal(t, 0, res1)
			assert.Equal(t, 0, res2)
			assert.Equal(t, 0, res3)
		}

		funcWithOk := func(a, b, c int) (res1, res2, res3 int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			res1, res2, res3 = e.Try3(multiplyByTwo(a, b, c)).Fallback(func(err error) (int, int, int) {
				return a * 2, b * 2, c * 2
			})
			return res1, res2, res3, nil
		}
		res1, res2, res3, err = funcWithOk(2, 3, 4)
		assert.Nil(t, err)
		assert.Equal(t, 4, res1)
		assert.Equal(t, 6, res2)
		assert.Equal(t, 8, res3)
	})
}
func functionWithFourReturn(t *testing.T, tt TestCase) {
	t.Run(tt.name+" functionWithFourReturn", func(t *testing.T) {
		multiplyByTwo := func(a, b, c, d int) (int, int, int, int, error) {
			if tt.expectedErrorText == "" {
				return a * 2, b * 2, c * 2, d * 2, nil
			}
			return 0, 0, 0, 0, errors.New(tt.expectedErrorText)
		}
		funcWithErrorReturn := func(a, b, c, d int) (res1, res2, res3, res4 int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			// if custom handler is not nil, use it

			if tt.withCustomHandler != nil {
				res1, res2, res3, res4 = e.Try4(multiplyByTwo(a, b, c, d)).If(e.Contains(tt.subErrorText)).Err(tt.withCustomHandler)
				return res1, res2, res3, res4, nil
			} else if tt.withCustomHandler != nil {
				res1, res2, res3, res4 = e.Try4(multiplyByTwo(a, b, c, d)).ErrMessage(tt.withMessage)
				return res1, res2, res3, res4, nil
			} else {
				res1, res2, res3, res4 = e.Throw4(multiplyByTwo(a, b, c, d))
				return res1, res2, res3, res4, nil
			}
		}

		// call function
		res1, res2, res3, res4, err := funcWithErrorReturn(2, 3, 4, 5)
		if tt.expectedErrorText == "" {
			assert.Nil(t, err)
			assert.Equal(t, 4, res1)
			assert.Equal(t, 6, res2)
			assert.Equal(t, 8, res3)
			assert.Equal(t, 10, res4)
		} else {
			assert.ErrorContains(t, err, tt.expectedErrorText)
			assert.Equal(t, 0, res1)
			assert.Equal(t, 0, res2)
			assert.Equal(t, 0, res3)
			assert.Equal(t, 0, res4)
		}

		funcWithOk := func(a, b, c, d int) (res1, res2, res3, res4 int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			res1, res2, res3, res4 = e.Try4(multiplyByTwo(a, b, c, d)).Fallback(func(err error) (int, int, int, int) {
				return a * 2, b * 2, c * 2, d * 2
			})
			return res1, res2, res3, res4, nil
		}
		res1, res2, res3, res4, err = funcWithOk(2, 3, 4, 5)
		assert.Nil(t, err)
		assert.Equal(t, 4, res1)
		assert.Equal(t, 6, res2)
		assert.Equal(t, 8, res3)
		assert.Equal(t, 10, res4)
	})

}

func functionWithFiveReturn(t *testing.T, tt TestCase) {
	t.Run(tt.name+" functionWithFiveReturn", func(t *testing.T) {
		multiplyByTwo := func(a, b, c, d, e int) (int, int, int, int, int, error) {
			if tt.expectedErrorText == "" {
				return a * 2, b * 2, c * 2, d * 2, e * 2, nil
			}
			return 0, 0, 0, 0, 0, errors.New(tt.expectedErrorText)
		}
		funcWithErrorReturn := func(a, b, c, d, ee int) (res1, res2, res3, res4, res5 int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			// if custom handler is not nil, use it

			if tt.withCustomHandler != nil {
				res1, res2, res3, res4, res5 = e.Try5(multiplyByTwo(a, b, c, d, ee)).If(e.Contains(tt.subErrorText)).Err(tt.withCustomHandler)
				return res1, res2, res3, res4, res5, nil
			} else if tt.withCustomHandler != nil {
				res1, res2, res3, res4, res5 = e.Try5(multiplyByTwo(a, b, c, d, ee)).ErrMessage(tt.withMessage)
				return res1, res2, res3, res4, res5, nil
			} else {
				res1, res2, res3, res4, res5 = e.Throw5(multiplyByTwo(a, b, c, d, ee))
				return res1, res2, res3, res4, res5, nil
			}
		}

		// call function
		res1, res2, res3, res4, res5, err := funcWithErrorReturn(2, 3, 4, 5, 6)
		if tt.expectedErrorText == "" {
			assert.Nil(t, err)
			assert.Equal(t, 4, res1)
			assert.Equal(t, 6, res2)
			assert.Equal(t, 8, res3)
			assert.Equal(t, 10, res4)
			assert.Equal(t, 12, res5)
		} else {
			assert.ErrorContains(t, err, tt.expectedErrorText)
			assert.Equal(t, 0, res1)
			assert.Equal(t, 0, res2)
			assert.Equal(t, 0, res3)
			assert.Equal(t, 0, res4)
			assert.Equal(t, 0, res5)
		}

		funcWithOk := func(a, b, c, d, ee int) (res1, res2, res3, res4, res5 int, err error) {
			defer e.Handle(&err, tt.errorHandle)
			res1, res2, res3, res4, res5 = e.Try5(multiplyByTwo(a, b, c, d, ee)).Fallback(func(err error) (int, int, int, int, int) {
				return a * 2, b * 2, c * 2, d * 2, ee * 2
			})
			return res1, res2, res3, res4, res5, nil
		}
		res1, res2, res3, res4, res5, err = funcWithOk(2, 3, 4, 5, 6)
		assert.Nil(t, err)
		assert.Equal(t, 4, res1)
		assert.Equal(t, 6, res2)
		assert.Equal(t, 8, res3)
		assert.Equal(t, 10, res4)
		assert.Equal(t, 12, res5)
	})

}

func TestHandleFunctions(t *testing.T) {

	t.Run("shouldn't recover non exception panics", func(t *testing.T) {
		nonExceptionPanic := func() {
			defer e.HandleErr(nil)

			panic(errors.New("non exception panic"))
		}
		assert.PanicsWithError(t, "non exception panic", nonExceptionPanic)
	})

	t.Run("should recover  exception panics", func(t *testing.T) {
		errlessException := func() {
			defer e.HandleErr(nil)

			e.Throw(errors.New("errless exception"))
		}
		errlessException()
	})

}
