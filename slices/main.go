package main

import "fmt"

func main() {
	integers := []int{1, 2, 3, 4, 5};
	fmt.Println("Before Appending: ")
	fmt.Println(integers);
	var newElement int;

	// inputting new element
	for i:=0; i < 2;i++ {
		fmt.Print("Enter New Element for Appending: ");
		fmt.Scanln(&newElement);
		integers = append(integers, newElement)
	}

	// printing the whole slice
	fmt.Println("After Appending: ")
	fmt.Println(integers);

}