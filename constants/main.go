package main

import (
	"fmt"
	//"math"
)

func main() {
	// declaring a constant
	const age int = 18;
	fmt.Printf("Age Constant: %v\n", age);
	// constants value cannot be change once declared
	// age = 23;

	// constants values cannot be set as a return value of a function
	// the below line will generate an error if uncommented
	// angle := 27;
	// const sin float32 = int(math.Sin(angle));
	// fmt.Printf("Sin(%v): %f\n", angle, sin ); 

	// constant can be used in arthimatic or logical opertors 
	const a int = 18;
	const b int = 28;

	// arthimatic operations
	const c int = a + b;
	const d int = a - b;
	const e int = a * b;
	const f int = a / b;
	// remainder operator
	const g int = a % b;


	fmt.Printf("c: %v\n", c);
	fmt.Printf("d: %v\n", d);
	fmt.Printf("e: %v\n", e);
	fmt.Printf("f: %v\n", f);
	fmt.Printf("g: %v\n", g);
	
	// logical opertion
	const or int = a | b;
	const and int = a & b;
	const xor int = a ^ b;
	const nand int = a &^ b;

	fmt.Printf("or: %v\n", or);
	fmt.Printf("and: %v\n", and);
	fmt.Printf("xor: %v\n", xor);
	fmt.Printf("nand: %v\n", nand);

	// can also be declared without any data type and the compiler will figure it out
	const canDrive = age == 18;
	fmt.Printf("canDrive: %v\n", canDrive); 

}