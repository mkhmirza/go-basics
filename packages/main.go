package main
import (
	"fmt"
	"os"
	"bufio"
	// to import your defined package, write the full path
	"github.com/mkhmirza/basics/packages/str"
	"strings"
)

func main() {

	fmt.Print("Enter a string to reverse: ");
	// fmt.scanln doesnt accept string with whitespaces
	scanner := bufio.NewReader(os.Stdin)
    s, _:= scanner.ReadString('\n');

	// removing '\n' from the string end
	s = strings.TrimSuffix(s, "\n")
	
	fmt.Printf("Original String: %v\n", s);
	fmt.Printf("Reverse String: %v\n", str.ReverseString(s));
}