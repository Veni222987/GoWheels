package leetcode

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := []string{}
	numStack := []int{}
	for i := range s {
		if s[i] >= 0 && s[i] <= 9 {
			last := len(numStack) - 1
			numStack[last] = numStack[last]*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			numStack = append(numStack, 0)
		} else if s[i] == ']' {

		} else {

		}
	}
	sb := strings.Builder{}
	for i := range stack {
		sb.WriteString(stack[i])
	}
	return sb.String()
}

func decodeStringOfficial(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
