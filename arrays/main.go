package main

import "fmt"

func main() {
	
	// declaring an array 
	var a [2] string;
	a[0] = "Hello";
	a[1] = "World";

	fmt.Println(a[0], a[1]);
	fmt.Println()

	// printing elements of an array using for-loop
	integers := [10]int{0,1,2,3,4,5,6,7,8,9};
	for i := 0; i < 10; i++ {
		fmt.Printf("This is the %d line\n", (integers[i] + 1));
	}


}