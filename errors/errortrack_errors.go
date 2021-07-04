package main

import (
	"fmt"

	"github.com/marmotedu/errors"

	code "github.com/marmotedu/sample-code"
)

func main() {
	if err := getUser(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func getUser() error {
	if err := queryDatabase(); err != nil {
		return errors.Wrap(err, "get user failed.")
	}

	return nil
}

func queryDatabase() error {
	return errors.WithCode(code.ErrDatabase, "user 'Lingfei Kong' not found.")
}
