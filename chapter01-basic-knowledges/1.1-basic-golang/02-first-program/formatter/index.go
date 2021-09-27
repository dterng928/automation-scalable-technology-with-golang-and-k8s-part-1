package formatter

import "fmt"

// MyPrintln will print input string ('M' capital public)
func MyPrintln(input string) {
	myPrintln(input)
}

// myPrintln will not export ('m' private)
func myPrintln(input string) {
	fmt.Println(input)
}
