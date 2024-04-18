package leetcode

func finalString(s string) string {
	q := []rune{}
	head := false
	for _, ch := range s {
		if ch != 'i' {
			if head {
				q = append([]rune{ch}, q...)
			} else {
				q = append(q, ch)
			}
		} else {
			head = !head
		}
	}
	if head {
		reverse(q)
	}
	return string(q)
}

func reverse(arr []rune) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
