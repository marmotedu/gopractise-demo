package main

import (
	"fmt"

	"github.com/marmotedu/errors"
	code "github.com/marmotedu/sample-code"
)

func main() {
	if err := getUser(); err != nil {
		fmt.Printf("%s\n", err)
		// do some business process based on the error type
		if errors.IsCode(err, code.ErrEncodingFailed) {
			fmt.Println("this is a ErrEncodingFailed error")
		}

		// we can also find the cause error
		fmt.Println(errors.Cause(err))
	}
}

func getUser() error {
	if err := queryDatabase(); err != nil {
		// Step3: Wrap the error with a new error message and a new error code if needed.
		return errors.WrapC(err, code.ErrEncodingFailed, "encoding user 'Lingfei Kong' failed.")
	}

	return nil
}

func queryDatabase() error {
	// Step1. Create error with specified error code.
	return errors.WithCode(code.ErrDatabase, "user 'Lingfei Kong' not found.")
}
