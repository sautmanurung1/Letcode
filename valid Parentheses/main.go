package main

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top == '(' && c != ')' {
				return false
			}
			if top == '[' && c != ']' {
				return false
			}
			if top == '{' && c != '}' {
				return false
			}
		}
	}
	return len(stack) == 0
}