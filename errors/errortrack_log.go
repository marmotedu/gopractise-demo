package main

import (
	"fmt"

	"github.com/marmotedu/errors"
	"github.com/marmotedu/log"

	code "github.com/marmotedu/sample-code"
)

func main() {
	if err := getUser(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func getUser() error {
	if err := queryDatabase(); err != nil {
		return err
	}

	return nil
}

func queryDatabase() error {
	opts := &log.Options{
		Level:            "info",
		Format:           "console",
		EnableColor:      true,
		EnableCaller:     true,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{},
	}

	log.Init(opts)
	defer log.Flush()

	err := errors.WithCode(code.ErrDatabase, "user 'Lingfei Kong' not found.")
	if err != nil {
		log.Errorf("%v", err)
	}
	return err
}
