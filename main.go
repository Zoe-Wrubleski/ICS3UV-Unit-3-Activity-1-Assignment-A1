/*
 * Author Zoe Wrubleski
 * Version 1.0.0
 * Date 2025-11-17
 * Fileoverview This program calculates the area of a circle
 */

package main

import "fmt"

func main() {
	// set numbers
	const Pi float32 = 3.14
	var r float32 = 5.6
	
	// calculate area
	answer := Pi * (r * r)

	// display answer
	fmt.Println("The area of a circle with a radius of", r, "cm is", answer, "cm².")

	fmt.Println("\nDone.")
}
