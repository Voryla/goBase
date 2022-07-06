package aboutStack

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { // 如果遇到了右括号
			// 若当前栈没有用于匹配的左括号或栈顶的左括号与当前的右括号不匹配，则错误
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 栈顶左括号匹配成功，出栈
			stack = stack[:len(stack)-1]
		} else { // 如果遇到了左括号，入栈
			stack = append(stack, s[i])
		}
	}
	// 若循环结束后当前栈中没有元素，说明全部匹配成功
	return len(stack) == 0
}
