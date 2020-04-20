package sort

// HeapSort 堆排序
func HeapSort(arr []int) {
	// 构建一个大顶堆
	buildHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		// 将堆顶放在数组末尾，待排序部分则变为len-1
		arr[i], arr[0] = arr[0], arr[i]
		downAdjust(arr, 0, i)
	}
}

// 构建堆
func buildHeap(arr []int) {
	size := len(arr)
	for i := (size - 1 - 1) >> 1; i >= 0; i-- {
		downAdjust(arr, i, size)
	}
}

// 向下调整->如果左右子节点有大于顶的，则与根节点交换
// 因为根节点下沉到左右的某一个，可能影响到了原先的堆结构
// 所以需要继续向下调整
func downAdjust(arr []int, parent int, size int) {
	temp := arr[parent]

	for i := parent*2 + 1; i < size; i = i*2 + 1 {
		// 存在右子树节点，且右子树节点值大于左子树节点值
		if i+1 < size && arr[i] < arr[i+1] {
			i++
		}
		// 左右子树中较大值拿来与根比较
		// 被根替换的位置暂时先不需要填充，因为已经存储在temp变量
		if arr[i] > temp {
			arr[parent] = arr[i]
			parent = i
		} else {
			break
		}
	}

	arr[parent] = temp
}

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// SelectionSort 选择排序排序
func SelectionSort(arr []int) {
	// 进行len-1轮
	for i := 0; i < len(arr)-1; i++ {
		// 每一轮以最前面的数为基准
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[minIdx], arr[i] = arr[i], arr[minIdx]
		}
	}
}

// QuickSort 快速排序
func QuickSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	pivot := partition(arr)
	if pivot != 0 {
		QuickSort(arr[:pivot])
	}
	if pivot+1 < len(arr) {
		QuickSort(arr[pivot+1:])
	}
}
func partition(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	pivot := len(arr) / 2
	end := len(arr) - 1
	// 比pivot数小的计数器，也可以认为是下一个待交换的位置
	smallCount := 0

	swap(&arr[pivot], &arr[end])
	for i := 0; i < end; i++ {
		//因为pivot已经被移动到end了
		if arr[i] < arr[end] {
			// 自己和自己换个锤子啊！
			if i != smallCount {
				swap(&arr[i], &arr[smallCount])
			}
			smallCount++
		}
	}
	swap(&arr[end], &arr[smallCount])
	// 返回现在pivot所在的位置
	return smallCount
}
func swap(a, b *int) {
	*a, *b = *b, *a
}

// CountingSort 计数排序
func CountingSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	//range:[min,max]
	min, max := arr[0], arr[0]
	for _, num := range arr[1:] {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	var count = make([]int, max-min+1)

	//遍历源数组，统计各个数字出现次数
	for _, num := range arr {
		count[num-min]++
	}
	//遍历统计数组，填充原数组
	i := 0
	for num, c := range count {
		for c != 0 {
			arr[i] = num + min
			i++
			c--
		}
	}
}
