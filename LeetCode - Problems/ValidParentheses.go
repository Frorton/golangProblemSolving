/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']'
Determine if the input string is valid.
An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.
*/

package ValidParentheses

func isValid(s string) bool {

	stack := []rune{}

	for _, input := range s {
		switch input {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			l := len(stack)

			if l == 0 {
				return false
			} else {
				population := stack[l-1]
				stack = stack[:l-1]
				if population != input {
					return false
				}
			}
		}
	}
	return len(stack) == 0
}
