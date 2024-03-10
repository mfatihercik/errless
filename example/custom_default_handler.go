//go:build exmaple

package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/mfatihercik/errless"
)

type Result struct {
	Success bool
}

func riskyOperation() (Result, error) {
	return Result{Success: false}, errors.New("something went wrong")
}

func callRiskyOperation() (res Result, err error) {
	defer errless.Handle(&err, func(err error) error {
		// You can also modify the error before returning it
		return fmt.Errorf("Main Handler :%w ", err)
	})
	customHandler := func(err error) error {
		return fmt.Errorf("Custom Handler :%w ", err)
	}
	result := errless.Check1W(riskyOperation()).With(customHandler)
	return result, nil
}

func main() {
	res, err := callRiskyOperation()
	if err != nil {
		fmt.Println(res)
		log.Fatal(err)

	}

}
