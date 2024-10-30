package internal

/*
Factorial calculates the factorial of a non-negative integer. The factorial 
of a number n is the product of all positive integers less than or equal to n.

Parameters:

	number int: The non-negative integer for which to calculate the factorial.

Returns:

	int: The factorial of the specified number.
*/
func Factorial(number int) int {
	// Base case
	if number <= 1 {
		return 1
	}
	// Recursive case
	return number * Factorial(number-1)
}