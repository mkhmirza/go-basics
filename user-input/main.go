package main

import "fmt"

func main() {
	// declaring vars for name
	var fname string;
	var lname string;

	// user input name
	fmt.Print("Enter Your First Name: ");
	fmt.Scanln(&fname);
	fmt.Print("Enter Your Last Name: ");
	fmt.Scanln(&lname)

	// displaying the combined name
	// concating both names for creating a full name
	fullname := fname + " " + lname;
	fmt.Printf("Your Full Name is %v.\n", fullname)
}