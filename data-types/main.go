package main

import "fmt"

func main() {
	// primitive data types supported by go
	// see vars/main.go for declaring a variable

	// int 
	// uint8  : 0 to 255 
	// uint16 : 0 to 65535 
	// uint32 : 0 to 4294967295 
	// uint64 : 0 to 18446744073709551615 
	// int8   : -128 to 127 
	// int16  : -32768 to 32767 
	// int32  : -2147483648 to 2147483647 
	// int64  : -9223372036854775808 to 9223372036854775807
	// see for more information: https://golang.org/ref/spec#Numeric_types

	var u uint = 43;
	var u8 uint8 = 23;
	var u16 uint16 = 44;
	var u32 uint32 = 67;
	var u64 uint64 = 90;
	
	var i int = 69;	
	var i8 int8 = 126;
	var i16 int16 = 255;
	var i32 int32 = 420;
	var i64 int64 = 6969;

	var f32 float32 =  4.56;
	var f64 float64 = 23.445;

	var c64 complex64 = 5 + 2i;
	var c128 complex128 = 6 + 10i;


	fmt.Println()
	fmt.Println("Unsigned Integers (Positive Int): ");
	fmt.Printf("%v, %T\n", u, u);
	fmt.Printf("%v, %T\n", u8, u8);
	fmt.Printf("%v, %T\n", u16, u16);
	fmt.Printf("%v, %T\n", u32, u32);
	fmt.Printf("%v, %T\n", u64, u64);

	fmt.Println()
	fmt.Println("Integers: ");
	fmt.Printf("%v, %T\n", i, i);
	fmt.Printf("%v, %T\n", i8, i8);
	fmt.Printf("%v, %T\n", i16, i16);
	fmt.Printf("%v, %T\n", i32, i32);
	fmt.Printf("%v, %T\n", i64, i64);

	fmt.Println()
	fmt.Println("Floating Point: ");
	fmt.Printf("%v, %T\n", f32, f32);
	fmt.Printf("%v, %T\n", f64, f64);

	fmt.Println()
	fmt.Println("Complex Numbers: ");
	fmt.Printf("%v, %T\n", c64, c64);
	fmt.Printf("%v, %T\n", c128, c128);
	
	// to seprate imginary and real part from a complex number use
	// real() : returns the real part
	// imag() : returns the imginary part
	// complex number 64 = float32 (real part) + float32 (imaginary part)
	// complex number 128 = float64 (real part) + float64 (imaginary part)
	// seperating from the 64 complex64 thats why it returns the type float32
	fmt.Printf("Real Part(64): %v, %T\n", real(c64), real(c64) );
	fmt.Printf("Imaginary Part(64): %v, %T\n", imag(c64), imag(c64));
	// if we seperate from complex128 it'll return the type float64	
	fmt.Printf("Real Part(128): %v, %T\n", real(c128), real(c128) );
	fmt.Printf("Imaginary Part(128): %v, %T\n", imag(c128), imag(c128));

	// string type
	var str string = "Hello World!";

	fmt.Println();
	fmt.Println("Strings: ");
	fmt.Printf("%v, %T\n", str, str);
	// strings can also be represented as collection of bytes 
	// see conversions/main.go to convert string to byte and vice versa
	var by byte = 110;
	fmt.Println();
	fmt.Println("Byte: ");
	// byte can also be represented using uint8
	fmt.Printf("%v, %T\n", by, by);

	// boolean are also present
	var canDrive bool = true;
	fmt.Println();
	fmt.Println("Booleans: ");
	fmt.Printf("%v, %T\n", canDrive, canDrive);

}