package aboutString

func longestCommonPrefix(strs []string) string {
	val := ""
	for i := 0; i < len(strs[0]); i++ {
		temp := strs[0][i]
		// strconv.QuoteRune(rune(temp))
		if string(temp) == "" {
			break
		}
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) {
				return val
			}
			if temp == strs[j][i] {
				continue
			} else {
				return val
			}
		}
		val = strs[0][:i+1]
	}
	return val
}
