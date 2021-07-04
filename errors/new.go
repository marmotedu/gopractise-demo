package main

import (
	"github.com/marmotedu/errors"
)

func main() {
	errors.New("file not found")
	errors.Errorf("file %s not found", "iam-apiserver")
}
