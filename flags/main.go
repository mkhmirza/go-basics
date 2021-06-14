package main

import (
	"fmt"
	"flag"
)

func main() {
	// returns a string pointer
	// flag.String(name, defaultValue, description);
	wordPtr := flag.String("word", "foo", "A String");
	// returns a int pointer, declaring a int flag
	intPtr := flag.Int("numb", 0, "An Integer");
	boolPtr := flag.Bool("bool", false, "True or False");

	// parsing the flags
	flag.Parse()

	fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *intPtr)
    fmt.Println("Bool:", *boolPtr)
	// any other options given which are not declared
	fmt.Println("tail:", flag.Args())
}