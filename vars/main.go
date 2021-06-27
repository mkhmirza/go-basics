package main

import "fmt"

// variables define out of a function can be used in any functions
// see variable scopes

// var occupation string = "Manager";

// a block of vars can be declared as

var (
	name = "Joe";
	age = 31;
	salary = 32.43;
	occupation = "Manager"
)

func main() {
	
	// in scope variable declaration
	// name := "Joe";
	// var age = 31;
	// salary := 32.42;
	
	fmt.Print("My name is " + name + ".");
	fmt.Printf(" My age is %d.",age);
	fmt.Printf(" My Salary is $%f.",salary);
	fmt.Printf(" My occupation is %v.\n", occupation);
}