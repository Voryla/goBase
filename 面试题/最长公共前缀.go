package main

/*
	最长公共前缀
	description:
	时间复杂度：O(mn)其中 mm 是字符串数组中的字符串的平均长度，nn 是字符串的数量。最坏情况下，字符串数组中的每个字符串的每个字符都会被比较一次。
	空间复杂度：O(1)使用的额外空间复杂度为常数。
	author: zwk
	date: 2022/2/22
*/
func getPrefix(strings []string) string {
	// 如果切片长度为0退出
	if len(strings) == 0 {
		return "输入不存在公共前缀。"
	}
	// 根据第一个字符串长度，遍历每个字符进行对比
	for i := 0; i < len(strings[0]); i++ {
		// 从第二个字符串开始依次对比第i个字符
		for j := 1; j < len(strings); j++ {
			// 若当前下标已超出要对比的字符串长度，或者字符对比失败
			if i == len(strings[j]) || strings[j][i] != strings[0][i] {
				// 返回第1个字符串，的0~i-1 字符串
				return strings[0][:i]
			}
		}
	}
	// 所有字符对比成功
	return strings[0]
}
func main() {
	s := []string{"flower", "flow", "flight"}
	println(getPrefix(s))
}
