package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

}

// 比较按字母顺序比较两个字符串返回一个整数。如果 a = b，结果是0; 如果 a < b，结果是-1; 如果 a > b，结果是 + 1
// 只有在与包字节对称时才包含比较。使用内置的字符串比较运算符 = 、 < 、 > 等通常更清晰，而且总是更快
func aboutCompare() {
	fmt.Println(strings.Compare("bb", "b"))
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}

// 包含 substr 是否在 s 内的报告
func aboutContains() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}

// func ContainsRune(s string, r rune) bool
// 报告 Unicode字符 r 是否在 s 内。
func aboutContainsRune() {
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
}

// func ContainsRune(s string, r rune) bool
// chars中的任意一个unicode代码点是否在s内
func aboutContainsAny() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

// func Count(s, substr string) int
// Count 对 s 中不重叠的 substr 实例的数量进行计数。如果 substr 是一个空字符串，Count 返回1 + s 中 Unicode 代码点的数量。
func aboutCount() {
	fmt.Println(strings.Count("cheese", "e")) // out 3
	fmt.Println(strings.Count("five", ""))    // out 5
}

// func EqualFold(s, t string) bool
// EqualFold 报告在 Unicode 大小写折叠下，被解释为 utf-8字符串的 s 和 t 是否相等，这是大小写不敏感的一种更一般的形式
func aboutEqualFold() {
	fmt.Println(strings.EqualFold("Go", "go")) // out: ture
}

// func Fields(s string) []string
// Fields 将字符串 s 分隔为一个或多个连续空白字符的每个实例，如果 s 只包含空白，则返回 s 的子字符串片段或空白片段。
func aboutFields() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//out: Fields are: ["foo" "bar" "baz"]
}

// func FieldsFunc(s string, f func(rune) bool) []string
// FieldsFunc 在每一次 Unicode 代码点 c 满足 f (c)时分割字符串 s，并返回 s 的片数组。
// 如果 s 中的所有代码点都满足 f (c)或字符串为空，则返回一个空片。
func aboutFieldsFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
	// out:Fields are: ["foo1" "bar2" "baz3"]
}

// func Split(s, sep string) []string
// 使用 sep 分割字符串，返回一个string切片
func aboutSplit() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	// out:
	// ["a" "b" "c"]
	// ["" "man " "plan " "canal panama"]
	// [" " "x" "y" "z" " "]
	// [""]
}

// func HasPrefix(s, prefix string) bool
// HasPrefix 测试字符串 s 是否以prefix开头。
func aboutHasPrefix() {
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	// out : true false
}

// func HasSuffix(s, suffix string) bool
// HasSuffix 测试字符串 s 是否以suffix结尾
func aboutHasSuffix() {
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
}

// func IndexAny(s, chars string) int
// IndexAny 返回来自字符的任何 Unicode字符的第一个实例的索引(以 s 表示) ，如果 s 中没有字符的 Unicode字符，返回-1。
func aboutIndexAny() {
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) // out: 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // out: -1
}

// func IndexByte(s string, c byte) int
// IndexByte 返回第一个 c 实例的索引，单位为 s，如果 c 不在 s 中，则返回-1。
func aboutIndexByte() {
	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))
	// out: 0  3 -1
}

// func IndexFunc(s string, f func(rune) bool) int
// IndexFunc 将满足 f (c)的第一个 Unicode字符的索引返回 s，如果没有满足 f (c) ，则返回-1
func aboutIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
	// out : 7 -1
}

// func IndexRune(s string, r rune) int
// IndexRune 返回 Unicode字符 r 的第一个实例的索引，如果 s 中没有 rune，返回-1。如果 r 是 utf8.RuneError，则返回任何无效 utf-8字节序列的第一个实例。
func aboutIndexRune(s string, r rune) {
	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
	// out: 4,-1
}

// func Join(elems []string, sep string) string
// Join 连接其第一个参数的元素，以创建单个字符串。分隔符字符串部分被放置在结果字符串中的元素之间
func aboutJoin() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	// out: foo, bar, baz
}

// func LastIndex(s, substr string) int
// LastIndex 返回最后一个 substr 实例的索引，单位为 s，如果在 s 中不存在 substr，则返回-1
func aboutLastIndex() {
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
}

// func LastIndexAny(s, chars string) int
// LastIndexAny 返回来自字符的任何 Unicode字符的最后一个实例的索引，单位为 s，如果 s 中没有字符的 Unicode字符，则返回-1。
// 如果有多个字符存在，那么范围位置为最后的字符
func aboutLastIndexAny() {
	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
}

// func LastIndexByte(s string, c byte) int
// LastIndexByte 返回 c 的最后一个实例的索引，单位为 s，如果 c 不在 s 中，则返回-1。
func aboutLastIndexByte() {
	fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'x'))
	// out: 10 8 -1
}

// func LastIndexFunc(s string, f func(rune) bool) int
// LastIndexFunc 将满足 f (c)的最后一个 Unicode字符的索引返回 s，如果没有，则返回-1。
func aboutLastIndexFunc() {
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))
}

// func Map(mapping func(rune) rune, s string) string
// Map 返回字符串 s 的一个副本，其所有字符都根据 mapping 函数进行了修改。如果映射返回一个负值，则该字符将从字符串中删除，不进行替换。
func aboutMap() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	// out: 'Gjnf oevyyvt naq gur fyvgul tbcure...
}

// func Repeat(s string, count int) string
// Repeat 返回一个由字符串 s 的计数副本组成的新字符串,如果计数为负值或者计数结果(len (s) * 计数)溢出，就会出现恐慌。
func aboutRepeat() {
	fmt.Println("ba" + strings.Repeat("na", 2))
	// out: banana
}

// func Replace(s, old, new string, n int) string
// 用new替换n次，s中old子字符串，若n<0 则全替换
func aboutReplace() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	// out: oinky oinky oink  moo moo moo
}

// func ReplaceAll(s, old, new string) string
func aboutRelaceAll() {
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
	// out: moo moo moo
}
