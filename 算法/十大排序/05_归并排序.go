package main

// 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
//设定两个指针，最初位置分别为两个已经排序序列的起始位置；
//比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
//重复步骤 3 直到某一指针达到序列尾；
//将另一序列剩下的所有元素直接复制到合并序列尾。
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
