package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

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

	// fmt.scanln doesnt accept string with whitespaces
	fmt.Print("Enter a string with whitespaces: ");
	// fmt.scanln doesnt accept string with whitespaces
	scanner := bufio.NewReader(os.Stdin)
    s, _:= scanner.ReadString('\n');

	s = strings.TrimSuffix(s, "\n")
	fmt.Printf("String with Whitespaces: %v\n", s);

}