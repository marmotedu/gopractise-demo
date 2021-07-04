package main

import (
	"fmt"
	"log"
)

func main() {
	if err := funcA(); err != nil {
		log.Fatalf("call func got failed: %v", err)
		return
	}

	log.Println("call func success")
}

func funcA() error {
	if err := funcB(); err != nil {
		return err
	}

	return fmt.Errorf("func called error")
}

func funcB() error {
	return fmt.Errorf("func called error")
}
