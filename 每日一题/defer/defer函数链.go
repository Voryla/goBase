package _defer

import "fmt"

//答案是
//
//test 3 1
//test 2 2
//inc 2
//test 1 3
//main 2
// 我的这个测试代码更能够理解defer的逻辑。
//在执行defer t.Inc(3).Inc(2).Inc(1)时，从左到右先计算了3,2这两个,最后的1才压栈。
//根据结果可看到先执行return t.num，因为main 2表示最终返回值是2
//所以return先执行，然后再执行之前压栈的defer t.Inc(1)
func main() {
	fmt.Println("main", inc())
}

func inc() int {
	t := &test{num: 0}
	defer t.Inc(3).Inc(2).Inc(1)
	fmt.Println("inc", t.num)
	return t.num
}

type test struct {
	num int
}

func (t *test) Inc(flag int) *test {
	t.num++
	fmt.Println("test", flag, t.num)
	return t
}
