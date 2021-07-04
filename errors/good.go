package main

import (
	"log"

	"github.com/marmotedu/errors"
)

func main() {
	if err := funcA(); err != nil {
		log.Fatalf("call func got failed: %s", err)
		return
	}

	log.Println("call func success")
}

func funcA() error {
	if err := funcB(); err != nil {
		return err
	}

	return errors.New("func called error")
}

func funcB() error {
	return errors.New("func called error")
}
