package main

// for more information: https://golang.org/pkg/fmt/

import "fmt"

func main() {
	// integers

	// %b	base 2
	// %c	the character represented by the corresponding Unicode code point
	// %d	base 10
	// %o	base 8
	// %O	base 8 with 0o prefix
	// %q	a single-quoted character literal safely escaped with Go syntax.
	// %x	base 16, with lower-case letters for a-f
	// %X	base 16, with upper-case letters for A-F
	// %U	Unicode format: U+1234; same as "U+%04X"
	fmt.Println("Intergers: ");
	fmt.Println("-----------");
	fmt.Printf("Base(2): %b\n", 2);
	fmt.Printf("Character: %c\n", 65);
	fmt.Printf("Base(10): %d\n", 456);
	fmt.Printf("Base(8): %o\n", 32);
	fmt.Printf("Hex Lower Case(x): %x\n", 15);
	fmt.Printf("Hex Upper Case(X): %X\n", 15);

	// boolean

	// %t	the word true or false
	fmt.Printf("Boolean: %t\n", false);


	// any value can be printed as
	fmt.Printf("%v\n", 43);
	fmt.Printf("%v\n", "Hello World");
	fmt.Printf("%v\n", true);

	// type of a variable can be determined by
	fmt.Printf("%v, %T\n", 43, 43);
	fmt.Printf("%v, %T\n", "Hello World", "Hello World");
	fmt.Printf("%v, %T\n", true, true);


	
}