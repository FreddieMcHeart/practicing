package main

import "fmt"

func ValidParentheses(parens string) (result bool) {
	parentheses := map[string]int{
		"(": 1,
		")": -1,
	}

	var count = 0

	for _, elem := range parens {
		count += parentheses[string(elem)]
		if count < 0 {
			return false
		}
	}

	if count == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(ValidParentheses("()"))
	fmt.Println(ValidParentheses(")()("))
	fmt.Println(ValidParentheses("()()((()))()()()()()"))
	fmt.Println(ValidParentheses("((()()()))"))
	fmt.Println(ValidParentheses("()((()))())"))
	fmt.Println(ValidParentheses("())(()"))
	fmt.Println(ValidParentheses(")()"))

}
