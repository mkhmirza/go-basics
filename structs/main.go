package main

// https://gobyexample.com/structs

import "fmt"

type person struct {
	name string;
	age int;
}

func newPerson(name string) *person {
	p := person{name: name};
	p.age = 42;
	return &p;
}

func main() {
	fmt.Println(person{"Bob", 55});
	fmt.Println(person{name: "Joe", age: 42});
	fmt.Println(person{name: "Joseph"});

	newperson := newPerson("Bob The Builder");
	fmt.Println(newperson.name);
	fmt.Println(newperson.age);

	fmt.Println(newPerson("John"));
}