package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func errorCheck(e error) {
	if e != nil {
		panic(e);
	}
}

func main() {

	filePath := "./read";
	// reading a file using ioutil
	dat, err := ioutil.ReadFile(filePath);
	errorCheck(err);
	fmt.Println(string(dat));

	// reading a file using os package
	// getting the address of the file
	f, e := os.Open(filePath)
	errorCheck(e);
	// reading a 5 bytes from a file, and saving it as bytes
	b := make([]byte, 5);
	//  how many bytes are being read
	n1, err  := f.Read(b);
	errorCheck(err);
	fmt.Printf("%d bytes: %s\n", n1, string(b[:n1]))

	// using buffer io
	bufferIo := bufio.NewReader(f);
	b1, err := bufferIo.Peek(6);
	errorCheck(err);
	fmt.Printf("5 bytes : %s\n", string(b1));
	// close file after reading
	f.Close();
}