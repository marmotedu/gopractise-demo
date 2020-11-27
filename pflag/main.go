package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var (
	stringToInt = pflag.StringToInt("stringToInt", map[string]int{"a": 1, "b": 2}, "stringToInt option")
	help        = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	pflag.Usage = func() {
		fmt.Println(`Usage: gentoken [OPTIONS] SECRETID SECRETKEY`)
		pflag.PrintDefaults()
	}
	pflag.Parse()

	if *help {
		pflag.Usage()
		return
	}

	fmt.Println("====> stringToInt exapmpe <====")
	for k, v := range *stringToInt {
		fmt.Println(k, v)
	}

}
