//go:build exmaple

package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/mfatihercik/errless"
)

// CustomErrorHandler handles errors by logging and optionally modifying them.
func CustomErrorHandler(err error) error {
	// You can also modify the error before returning it
	return fmt.Errorf("custom error :%w ", err)
}

type Result struct {
	Success bool
}

func riskyOperation() (Result, error) {
	return Result{Success: false}, errors.New("something went wrong")
}

func callRiskyOperation() (res Result, err error) {
	defer errless.Handle(&err, CustomErrorHandler)
	result := errless.Check1(riskyOperation())
	return result, nil
}

func main() {
	res, err := callRiskyOperation()
	if err != nil {
		fmt.Println(res)
		log.Fatal(err)

	}

}
