package leetcode

// 有效的括号
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if v == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else if v == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else if v == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
