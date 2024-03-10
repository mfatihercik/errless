//go:build test

package errless_test

import (
	"errors"
	"fmt"
	"testing"

	e "github.com/mfatihercik/errless"
	"github.com/stretchr/testify/assert"
)

func TestCheckFunctions(t *testing.T) {

	testcases := []struct {
		name              string
		expectedErrorText string
		errorHandle       func(err error) error
		customHandler     func(err error) error
	}{
		{name: "default handler should return error",
			expectedErrorText: "test error", // error case
			errorHandle: func(err error) error {
				assert.ErrorContains(t, err, "test error")
				return err
			},
			customHandler: nil,
		},
		{name: "default handler should return success",
			expectedErrorText: "", // success case
			errorHandle: func(err error) error {
				t.Fatal("shouldn't be called")
				return err
			},
			customHandler: nil,
		},
		{name: "custom handler should return customized error",
			expectedErrorText: "custom error", // error in the custom handler
			errorHandle: func(err error) error {
				assert.ErrorContains(t, err, "custom error")
				return err
			},
			customHandler: func(err error) error {
				return fmt.Errorf("custom error :%w ", err)
			},
		},
		{name: "custom handler should return success",
			expectedErrorText: "", // success case
			errorHandle: func(err error) error {
				t.Fatal("shouldn't be called")
				return err
			},
			customHandler: func(err error) error {
				t.Fatal("shouldn't be called")
				return err
			},
		},
	}
	for _, tt := range testcases {

		// zero return function
		functionWithZeroReturn(t, tt)
		functionWithOneReturn(t, tt)
		functionWithTwoReturn(t, tt)
		functionWithFourReturn(t, tt)
		functionWithFiveReturn(t, tt)

		// function with one return value and error

	}

}

func functionWithZeroReturn(t *testing.T, tt struct {
	name              string
	expectedErrorText string
	errorHandle       func(err error) error
	customHandler     func(err error) error
}) {
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
			if tt.customHandler != nil {
				e.CheckW(zeroParameterFunc()).With(tt.customHandler)
			} else {
				e.Check(zeroParameterFunc())
			}
			return err
		}

		// call function
		err := funcWithErrorReturn()
		if tt.expectedErrorText == "" {
			assert.Nil(t, err)
		} else {
			assert.ErrorContains(t, err, tt.expectedErrorText)
		}
	})
}

func functionWithOneReturn(t *testing.T, tt struct {
	name              string
	expectedErrorText string
	errorHandle       func(err error) error
	customHandler     func(err error) error
}) {
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

			if tt.customHandler != nil {
				return e.Check1W(multiplyByTwo(a)).With(tt.customHandler), nil
			} else {
				return e.Check1(multiplyByTwo(a)), nil
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
	})

}

func functionWithTwoReturn(t *testing.T, tt struct {
	name              string
	expectedErrorText string
	errorHandle       func(err error) error
	customHandler     func(err error) error
}) {
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

			if tt.customHandler != nil {
				res1, res2 = e.Check2W(multiplyByTwo(a, b)).With(tt.customHandler)
				return res1, res2, nil
			} else {
				res1, res2 = e.Check2(multiplyByTwo(a, b))
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
	})

}

func functionWithThreeReturn(t *testing.T, tt struct {
	name              string
	expectedErrorText string
	errorHandle       func(err error) error
	customHandler     func(err error) error
}) {
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

			if tt.customHandler != nil {
				res1, res2, res3 = e.Check3W(multiplyByTwo(a, b, c)).With(tt.customHandler)
				return res1, res2, res3, nil
			} else {
				res1, res2, res3 = e.Check3(multiplyByTwo(a, b, c))
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
	})
}
func functionWithFourReturn(t *testing.T, tt struct {
	name              string
	expectedErrorText string
	errorHandle       func(err error) error
	customHandler     func(err error) error
}) {
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

			if tt.customHandler != nil {
				res1, res2, res3, res4 = e.Check4W(multiplyByTwo(a, b, c, d)).With(tt.customHandler)
				return res1, res2, res3, res4, nil
			} else {
				res1, res2, res3, res4 = e.Check4(multiplyByTwo(a, b, c, d))
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
	})

}

func functionWithFiveReturn(t *testing.T, tt struct {
	name              string
	expectedErrorText string
	errorHandle       func(err error) error
	customHandler     func(err error) error
}) {
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

			if tt.customHandler != nil {
				res1, res2, res3, res4, res5 = e.Check5W(multiplyByTwo(a, b, c, d, ee)).With(tt.customHandler)
				return res1, res2, res3, res4, res5, nil
			} else {
				res1, res2, res3, res4, res5 = e.Check5(multiplyByTwo(a, b, c, d, ee))
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
	})

}
