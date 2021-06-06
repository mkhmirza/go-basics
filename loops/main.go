package main

import "fmt"

func main() {
	
	// typical for-loop
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, From A Foor Loop");
	}

	fmt.Println()
	// while loop
	sum := 0;
	counter := 0;
	for counter < 5 {
		fmt.Println("Hello From A While Loop")
		counter += 1;
		sum += counter;
	}
	fmt.Printf("Sum is: %d\n", sum);

}