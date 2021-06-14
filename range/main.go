package main

import "fmt"

func main() {
	// usage of range fnction
	array := []int{ 1, 2, 3, 4, 5};

	// looping throught the entire array using range
	for k, i := range array {
		fmt.Printf("Element#[%d]: %d\n",k, i);
	}

	kvs := map[int]string{1:"Joe", 2:"Barbara", 3:"Brad"};
	// looping through the map, accessing only the keys 
	for k, _ := range kvs {
		fmt.Printf("Key: %d\n", k);
	}

	// looping through the map, accessing only the values 
	for _, v := range kvs {
		fmt.Printf("Values: %s\n", v);
	}

}