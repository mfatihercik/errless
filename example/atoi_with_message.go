//go:build exmaple

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mfatihercik/errless"
)

func sumWithMessage(a, b string) (res int, err error) {
	defer errless.Handle(&err, errless.EmptyHandler)
	x := errless.Try1(strconv.Atoi(a)).ErrMessage("cannot convert to int")
	y := errless.Try1(strconv.Atoi(b)).ErrMessage("cannot convert to int")
	return x + y, nil
}

func main() {
	res, err := sum("10", "20t")
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
	}
	fmt.Println("result:", res)

	res, err = sumWithCatch("10", "20t")
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
	}
	fmt.Println("result:", res)

}
