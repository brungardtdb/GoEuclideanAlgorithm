/***************************************************************
Author: David Brungardt
Date: 05/13/2018
Program calls a function that uses Euclidean algorithm
to determine the greatest-common factor between two numbers.
The program uses the modulus operator to find the remainder
after dividing the larger number by the smaller number.
The process then continues using the smaller number in place
of the larger number, and the remainder in place of the smaller
number until the remainder is zero and the GCF is found.
***************************************************************/

package main

import (
	"fmt"
)

// Class variables
var firstInt = 3375
var secondInt = 550

func main() {

	// Print results of Euclidean algorithm
	fmt.Printf("Greatest common factor is: %d", euclideanGCF(firstInt, secondInt))

} // end of main

// Function to handle logic of Euclidean algorithm
func euclideanGCF(firstInt int, secondInt int) int {

	// Local fields for function
	var remainder int
	var dividend int

	// Use modulus operator to find remainder
	// of larger number divided by smaller number
	if firstInt > secondInt {
		remainder = (firstInt % secondInt)
		dividend = secondInt
	} else {
		remainder = (secondInt % firstInt)
		dividend = firstInt
	}

	// Return GCF when remainder is 0
	// Call function recursively until 0 is reached
	if remainder == 0 {
		return dividend
	} else {
		return euclideanGCF(dividend, remainder)
	}

} // end of EuclideanGCF
