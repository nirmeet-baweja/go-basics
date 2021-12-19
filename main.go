/* To run this file type in the terminal - "go run main.go" */

// every go file needs to be a part of a package.
// this main.go file is a part of a package called "main"
// package main tells the go compiler to compile the file into a standalong executable file at the end
package main

// importing a single package
import "fmt"

// function definition.
// function main is the entry point of our program / application
// there can be only one main function per application
func main() {

	/***************************************************************************/
	// function names starts with capitals when they are exported.
	// Thus when we will use functions from other packages,
	// the function names will start with capitals.
	// print to console
	fmt.Println("Hello, ninjas!")

	/***************************************************************************/
}