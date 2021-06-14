package main

import (
	"fmt"
	"os"
)

func main()  {
	argsWithProg := os.Args;
	argsWithoutProg := os.Args[1:];

	fmt.Println("Arguments with Program: ", argsWithProg);
	fmt.Println("Arguments without Program: ", argsWithoutProg);

	// looping through all arguments
	for k, arg := range argsWithoutProg {
		fmt.Printf("Index: %d, Arg: %s\n", k, arg);
	}

}