package main

// print using ansi color coding


import (
	"fmt"
)

//   Black   = Color("\033[1;30m%s\033[0m")
//   Red     = Color("\033[1;31m%s\033[0m")
//   Green   = Color("\033[1;32m%s\033[0m")
//   Yellow  = Color("\033[1;33m%s\033[0m")
//   Purple  = Color("\033[1;34m%s\033[0m")
//   Magenta = Color("\033[1;35m%s\033[0m")
//   Teal    = Color("\033[1;36m%s\033[0m")
//   White   = Color("\033[1;37m%s\033[0m")

func main() {
	colorReset := "\033[0m";
	fmt.Println("Hello World");
	colorRed := "\033[31m";
	colorGreen := "\033[32m";
	colorYellow := "\033[33m"; 
	colorBlue := "\033[34m";
	colorPurple := "\033[35m";
	colorCyan := "\033[36m";
	colorWhite := "\033[37m";

	fmt.Println(string(colorRed), " test");
	fmt.Println(string(colorGreen), " test");
	fmt.Println(string(colorYellow), " test");
	fmt.Println(string(colorBlue), " test");
	fmt.Println(string(colorPurple), " test");
	fmt.Println(string(colorCyan), " test");
	fmt.Println(string(colorWhite), " test");
	fmt.Println(string(colorReset), "test")	

	// more complicated example
	var yourName string;
	fmt.Print("Enter Your Name: ");
	fmt.Scanln(&yourName);
	// color only the name
	yourName = fmt.Sprint(string(colorGreen), yourName);
	concat := "Your Name is " + yourName + "!@";

	fmt.Println(concat);

}