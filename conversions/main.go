package main

import (
	"fmt"
)

func main() {
	// go supports many explicit conversions 
	var a int8 = 10;
	var b int32 = 90;
	fmt.Println("Conversion from Int8 to Int32");
	// you cannot add different types together
	// meaning below line will throw an error
	// c := a + b;
	// since 'a' is of type 'int8' and 'b' is of type 'int32' 
	// the compiler will see this as mismatched data types 
	// we can get around this using explicit conversions
	// converting 'a' to a int32
	c := int32(a) + b;
	fmt.Printf("%v, %T\n", c, c); 
	
	// same with floats 
	var f32 float32 = 3.14;
	var f64 float64 = 3.4;
	// this will throw a mismatched error
	// ans := f32 + f64;
	// convert 'f32' to a float64
	ans := float64(f32) + f64;
	fmt.Printf("%v, %T\n", ans, ans); 

	// strings can be converted to an array of byte
	var str string = "Hello World!";
	// replaces each character with its ascii value
	fmt.Printf("%v, %T\n",[]byte(str), []byte(str) );

	// an array of byte can be converted into a string using string()
	var str1 = []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33};
	fmt.Printf("%v, %T\n", string(str1), string(str1));
}