package main

import "fmt"

func main() {
	fmt.Println("IsValid : ", isValid("()[]{}"))
}

func isValid(s string) bool {

	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []rune{}

	for _, v := range s {
		if _, ok := m[v]; ok {

			stack = append(stack, v)
		} else if len(stack) == 0 || m[stack[len(stack)-1]] != v {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}
