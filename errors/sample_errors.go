package main

import (
	"fmt"

	"github.com/marmotedu/errors"
	code "github.com/marmotedu/sample-code"
)

func main() {
	if err := getUser(); err != nil {
		fmt.Printf("%+v\n", err)
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
		return errors.WrapC(err, code.ErrEncodingFailed, "get user 'Lingfei Kong' failed.")
	}

	return nil
}

func queryDatabase() error {
	return errors.WithCode(code.ErrDatabase, "user 'Lingfei Kong' not found.")
}
