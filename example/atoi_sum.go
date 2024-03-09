//go:build exmaple

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mfatihercik/errless"
)

func sum(a, b string) (res int, err error) {
	defer errless.Handle(&err, errless.EmptyHandler)
	x := errless.Check1(strconv.Atoi(a))
	y := errless.Check1(strconv.Atoi(b))
	return x + y, nil
}

func main() {
	res, err := sum("10", "20")
	if err != nil {
		log.Fatalf("Error occurred: %v", err)
	}
	fmt.Println("result:", res)

}
