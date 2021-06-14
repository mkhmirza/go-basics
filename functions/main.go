package main

import "fmt"

func getSum(array []int) int {
	sum := 0;
	
	for i:=0;i<len(array);i++ {
		sum += array[i];
	}

	return sum;
}


func main() {
	integers := []int{1, 2, 3, 4, 5};
	fmt.Println("Sum: ", getSum(integers));
}