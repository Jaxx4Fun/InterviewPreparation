package effective

import "fmt"

func aboutRange() {
	score := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	// 只要key
	for key := range score {
		fmt.Printf("key:%s ", key)
	}
	fmt.Println()
	// 丢弃key
	for _, val := range score {
		fmt.Printf("val:%d ", val)
	}
	fmt.Println()
	// k，v都要
	for key, val := range score {
		fmt.Printf("map[%s]=%d ", key, val)
	}
	fmt.Println()

	nums := []int{1, 2, 3}
	for idx := range nums {
		fmt.Printf("idx:%d ", idx)
	}
	fmt.Println()

	for _, val := range nums {
		fmt.Printf("val:%d ", val)
	}
	fmt.Println()

	for idx, val := range nums {
		fmt.Printf("nums[%d]=%d\n", idx, val)
	}
}
