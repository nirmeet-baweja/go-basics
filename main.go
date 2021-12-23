/* To run this file type in the terminal - "go run main.go" */

// every go file needs to be a part of a package.
// this main.go file is a part of a package called "main"
// package main tells the go compiler to compile the file into a standalone executable file at the end
package main

// importing a single package
import "fmt"

// function definition.
// function main is the entry point of our program / application
// there can be only one main function per application
func main() {

	/***************************************************************************/
	/* Beginning */
	/***************************************************************************/
	// function names starts with capitals when they are exported.
	// Thus when we will use functions from other packages,
	// the function names will start with capitals.
	// print to console
	fmt.Println("Section 1")
	fmt.Println("Hello, ninjas!")

	/***************************************************************************/
	/* Variables */

	/* go is strongly typed language,
	 * and variable data type cannot be changed after declaration
	 */
	/***************************************************************************/
	fmt.Println("Section 2")

	/* string variables */
	var nameOne string = "mario"
	// when data type is not specified, go infers it from the assigned value
	var nameTwo = "luigi"
	// when a variable is declared but not initialised, data type needs to specified
	var nameThree string

	fmt.Print(nameOne, nameTwo, nameThree)

	// variables can be reassigned new values
	nameOne = "peach"
	nameThree = "bowser"

	fmt.Print(nameOne, nameTwo, nameThree)

	// the following is allowed inside functions only
	// A variable declared as below can be used only inside the function it is declared in
	nameFour := "yoshi"
	fmt.Print(nameFour)

	/* int variables */
	var ageOne int = 20
	var ageTwo = 30 // type inferred directly from the assigned value
	ageThree := 40  // short-hand

	fmt.Print(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	// var numTwo int8 = 128   // too large a number for 8-bit
	// var numThree uint = -25 // unsigned ints cannot be negative

	fmt.Print(numOne) //, numTwo, numThree)

	/* float variables */
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1965385877.5
	var scoreThree = 1.5 // inferred as float64

	fmt.Print(scoreOne, scoreTwo, scoreThree)
	fmt.Print("\n")

	/***************************************************************************/
	/* some fmt package functions available */
	/***************************************************************************/
	fmt.Println("Section 3")

	noun := "Dog"
	verb := "running"

	// print the exact characters that are within the double quotes
	fmt.Print("hello     !")
	fmt.Print(noun, "is", verb)

	// print the characters as well as add a new line at the end
	fmt.Println("hello there!")
	// can be used to concatenate variables and constants & then print them
	// and a space is added between each function parameter while printing
	fmt.Println(noun, "is", verb)

	age := 10
	name := "Maria"

	// Printf (print formatted string), %_ = format specifier
	fmt.Printf("my name is %v and my age is %v \n", name, age) // %v = value in default format
	fmt.Printf("my name is %q and my age is %q \n", name, age) // %q = quotes
	fmt.Printf("age is of type %T \n", age)                    // %T is the type
	fmt.Printf("you scored %f points! \n", 225.55)             // %f = float format
	fmt.Printf("you scored %0.1f points! \n", 225.55)          // %0.2f = float with 2 decimal points

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	fmt.Println("the saved string is:", str)

}
