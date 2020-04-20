# 排序算法

还是从最简单的说起可能比较好。

## 参考资料

- [十大经典排序算法（动图演示）](https://www.cnblogs.com/onepixel/articles/7674659.html)

## 正文

特殊的算法会提及一下应用场景

### 冒泡排序

对于我个人而言，冒泡排序不难，只是我每次思考 i、j 循环结束的条件需要花点时间。不妨代入循环看看：

第一轮循环：

- arr[0]与 arr[1]比较
- arr[1]与 arr[2]比较
- ...
- arr[N-2] 与 arr[N-1]比较

所以第一轮循环一共进行了 N-1 次比较;

第二轮循环：

- arr[0]与 arr[1]比较
- arr[1]与 arr[2]比较
- ...
- arr[N-3] 与 arr[N-2]比较

所以第二轮循环一共进行了 N-2 次比较

实例是排序后呈一个递增序列，每轮冒泡将最大值冒到数组末端，即第 1 轮将 N 中最大的数冒出（需要比较 N-1 次）、第 2 次将 N-1 个数中的最大数冒出（需要比较 N-2 次）……当进行到 N-1 轮时，将剩余的 2 个中最大值冒出，则最小值自然也出来了。

一共进行了 N-1 轮冒泡；而每一轮中比较次数分别是 N-1、N-2……、1，等差数列求和可以知道总比较次数 S=(N-1)(N-1+1)/2;所以时间复杂度 O(n^2);

#### 代码

```go
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
```

### 选择排序

第 1 轮循环：

在 nums[0,N-1]中，找到最小的数，交换到 0 的位置上；

第 2 轮循环：
在 nums[1,N-1]中，找到最小的数，交换到 1 的位置上；

第 N-1 轮循环：
在 nums[N-2,N-1]中，两个数，找到最小的数，交换到 N-2 位置上;

一共 N-1 轮循环， 比较的次数依次是 N-1,N-2,……,1，总比较次数 S=(N-1)(N-1+1)/2，时间复杂度 O(N^2)

#### 代码

```go
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
```

### 快速排序

#### 代码

```go
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
```

### 计数排序

假设数据范围是[m,n], m<n,即对于任意的 0<=i<N-1,有 n<=arr[i]<=m；记 k=m-n+1，创建一个 count[k]，在遍历数组时用于记录每个数字出现的次数；而后重新扫描 count 数组，众神归位。遍历 N 个元素的数组+遍历 K 个元素的计数数组，时间复杂度 O(N+k)；

创建了 K 个元素的计数数组（额外）空间复杂度 O(k)；而如果排序结果放在新的数组中，则还需要 N 个额外空间，那么空间复杂度就是 O(N+k)。

#### 代码

```go
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
```

#### 使用场景

已知所有数据的范围，而且范围较小，如：

- 一个公司员工年龄的数组
- 0、1 组成的数组

### 桶排序

桶排序是计数排序的升级版本，计数排序可以认为是每个桶只能放任意个某数字，每个桶内不需要排序，一共需要 max-min+1 个桶；

而一般的桶排序，一个桶里可以放不同的数字；所以每个桶里的数字还需要额外排序（递归桶排序或者其他排序算法）

### 堆排序HeapSort
堆的结构就是，堆中任意的节点都大于/小于其左右子节点。所以堆顶元素一定是最大值或最小值。这个结构很适合去存储一批数据里最大/最小的N个数

#### 代码

```go
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

```



#### 使用场景

- 找出数据量为1TB的数组中最大的100个数：一批一批地去遍历这1TB，比如每次内存中只加载1GB，然后用一个大小为100的小顶堆，也就是说堆顶是这100个数里最小的。那么如果输入一个数，大于堆顶，则把堆顶元素挤出Top100咯，此时需要对堆做一定的调整。