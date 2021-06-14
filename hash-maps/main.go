package main

import "fmt"

func main() {
	// making a map of string keys and int values
	m := make(map[string]int);

	m["k1"] = 1;
	m["k2"] = 2;

	fmt.Println("Map: ", m);
	fmt.Println("Len(Map): ", len(m));
	// deleting k1 key and its value
	delete(m, "k1");

	fmt.Println("Map: ",m);
	fmt.Println("Len(Map): ", len(m));

	// access value and see if the key exists
	value, isPresent := m["k2"];
	fmt.Println("Value: ", value);
	fmt.Println("isPresent?: ", isPresent);

	// declaring a map without make()
	n := map[string]int {"foo":1, "bar":2};
	// range can be used for looping through the entire map
	for key, value := range n {
		fmt.Printf("Key: %s, Value: %d\n", key, value);
		_, isKeyPresent := n[key];
		fmt.Printf("IsKeyPresent: %t\n", isKeyPresent);
	}

}