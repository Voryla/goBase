package aboutString

func main() {

	s := "cc"
	println(firstUniqChar(s))
}
func firstUniqChar(s string) byte {
	cnt := [26]int{}
	for _, runeValue := range s {
		// runeValue - 'a' 计算出对应字母所在的下标
		cnt[runeValue-'a']++
	}
	for index, runeValue := range s {
		if cnt[runeValue-'a'] == 1 {
			return s[index]
		}
	}
	return ' '
}
