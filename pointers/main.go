package main

import "fmt" 

func passByValue(num int) {
	num = 20;
}

func passByReference(num *int) {
	*num = 50;
}


func main() {
	rNum := 0;
	fmt.Println("Initial Value: ", rNum);
	fmt.Println("Address(Initial Value): ", &rNum);

	fmt.Println()
	// passing the var into passbyvalue fnctio
	passByValue(rNum);
	fmt.Println("After Pass By Value: ", rNum);
	fmt.Println("Address(Pass By Value): ", &rNum);

	fmt.Println()
	// passing the var into passbyreference
	passByReference(&rNum);
	fmt.Println("After Pass By Ref: ", rNum);
	fmt.Println("Address(Pass By Ref): ", &rNum);
}