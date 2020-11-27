package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var (
	flagvar = pflag.Int("flagname", 1234, "help message for flagname")
)

func main() {
	pflag.Parse()

	fmt.Printf("argument number is: %v\n", pflag.NArg())
	fmt.Printf("argument list is: %v\n", pflag.Args())
	fmt.Printf("the first argument is: %v\n", pflag.Arg(0))
}
