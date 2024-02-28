/*
Given an integer n, return a string array answer (1-indexed) where:
    answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
    answer[i] == "Fizz" if i is divisible by 3.
    answer[i] == "Buzz" if i is divisible by 5.
    answer[i] == i (as a string) if none of the above conditions are true.
*/

package FizzBuzz

import "strconv"

func fizzBuzz(n int) []string {

	result := []string{}

	//tried for i := range but found switch/case easier
	for i := 1; i <= n; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
