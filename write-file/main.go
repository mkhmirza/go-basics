package main

import (
    "fmt"
    "io/ioutil"
	"os"
	"bufio"
)

func errorCheck(e error) {
	if e != nil {
		panic(e);
	}
}

func main() {
	filePath := "./write";
	// writting to a file using ioutil
	data := []byte("This is a sample bro");
	err := ioutil.WriteFile(filePath, data, 0644);
	errorCheck(err);

	fmt.Printf("Written '%s' in the file\n", string(data));

	// creating a new file
	f, err := os.Create(filePath);
    errorCheck(err);

	// writting using buffer writer
	writeToFile := "Using buffer writer";
	w := bufio.NewWriter(f);
	// number of bytes written
	n, err := w.WriteString(writeToFile);
	errorCheck(err);
	fmt.Printf("Wrote to a file '%s', bytes wrote %d\n", writeToFile, n);
	w.Flush();

}