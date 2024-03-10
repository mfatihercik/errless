//go:build exmaple

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"

	e "github.com/mfatihercik/errless"
)

// exmaple from proposal

func SortContents(w io.Writer, files []string) (err error) {

	defer e.HandleReturn(func(err error) {
		// assign the error to the return value with named error
		err = fmt.Errorf("process: %v", err) // handler A
	})

	lines := []string{}
	for _, file := range files {
		handleB := func(err error) error {
			return fmt.Errorf("read %s: %v ", file, err) // handler B
		}
		scan := bufio.NewScanner(e.Check1W(os.Open(file)).With(handleB)) // handler B

		for scan.Scan() {
			lines = append(lines, scan.Text())
		}
		e.CheckW(scan.Err()).With(handleB) // handler B
	}
	sort.Strings(lines)
	for _, line := range lines {
		e.Check1(io.WriteString(w, line)) // check runs A on error
	}
	return nil
}
