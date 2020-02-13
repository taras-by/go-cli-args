package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	stringValue string
	intValue    int
	boolValue   bool
	helpText = `
Usage: foo [--<optional parameters>=value]
  Simple command that prints argument
Options:
  --string-value=<Just string value>
  --int-value=<Just int value>
  --bool-value=<Just bool value>
`
)

func main() {

	flags := flag.NewFlagSet("foo commands", flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Println(helpText)
	}

	flags.StringVar(&stringValue,  "string-value", "string", "Just string value")
	flags.IntVar(&intValue, "int-value", 60, "Just int value")
	flags.BoolVar(&boolValue, "bool-value", true, "Just bool value")

	args := os.Args[1:]
	if err := flags.Parse(args); err != nil {
		os.Exit(1)
	}

	fmt.Println("stringValue: ", stringValue)
	fmt.Println("intValue: ", intValue)
	fmt.Println("boolValue: ", boolValue)

	os.Exit(0)
}
