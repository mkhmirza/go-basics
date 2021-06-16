package main

import "fmt"

func main() {
	// logical operations such as
	// & (and), | (or), ^(exclusive or), nand(&^)

	a := 10; // base(2): 1010
	b := 3; // base(2): 0010

	// printing the binary representation of above numbers
	fmt.Printf("Binary(%v): %b\n", a, a);
	fmt.Printf("Binary(%v): %b\n", b, b);

	//  performing logical opertaions
	and := a & b;
	or := a | b;
	xor := a ^ b;
	nand := a &^ b;

	// decimal here means base(10)
	fmt.Printf("And= Decimal: %v, Binary: %b\n", and, and);
	fmt.Printf("Or = Decimal: %v, Binary: %b\n", or, or);
	fmt.Printf("Xor = Decimal: %v, Binary: %b\n", xor, xor);
	fmt.Printf("Nand = Decimal: %v, Binary: %b\n", nand, nand);

}