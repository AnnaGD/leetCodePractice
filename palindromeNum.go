package main

func isPalindrome(x int) bool {

	// reversed = 0

	//Grab the last digit (x % 10)

	// Add that digit to the end of reversed by doing:
	// reversed = reversed * 10 + last_digit

	// Cut off the last digit of x using x = x / 10

	// Next steps would be:

	// Grab the next digit: 12 % 10 = 2

	// Update reversed: 3 * 10 + 2 = 32

	// Shrink original: 12 / 10 = 1

	// Edge case: 
	if x < 0 || (x %10 == 0 && x != 0) {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		lastDigit := x % 10
		reversed = reversed * 10 + lastDigit
		x = x/10
	}

	return original == reversed
    
}


/*

9. Palindrome Number
Easy

Given an integer x, return true if x is a palindrome, and false otherwise.

 
Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 

Constraints:

-231 <= x <= 231 - 1
 

Follow up: Could you solve it without converting the integer to a string?

*/
