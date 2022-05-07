package main

// 创建一个堆 H[0……n-1]；
//把堆首（最大值）和堆尾互换；
//把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
//重复步骤 2，直到堆的尺寸为 1。
func heapSort(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap1(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap1(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swap1(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
