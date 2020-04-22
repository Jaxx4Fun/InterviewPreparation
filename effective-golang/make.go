package effective

import "fmt"

func makeFunc() {
	// // make 只能用于创建slice、map、chan
	// // 返回引用，而非*T
	// p1 := make([]int, 0, 10) //类似cpp vector<int> v; v.reserve(10)
	// p2 := make(map[string]int)
	// p3 := make(chan int)

	// p4 := new([]int) //会返回一个指向新分配的，已置零的切片结构， 即一个指向 nil 切片值的指针

	// var p *[]int = new([]int)      // allocates slice structure; *p == nil; rarely useful
	// var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints
	// // Unnecessarily complex:
	// {

	// 	var p1 *[]int = new([]int)
	// 	*p1 = make([]int, 100, 100)
	// 	// Idiomatic:
	// 	v1 := make([]int, 100)
	// }

}

func aboutSlice() [][]uint8 {
	const (
		XSize = 10
		YSize = 10
	)
	// 分配顶层切片，和前面一样。
	picture := make([][]uint8, YSize) // 每 y 个单元一行。
	// 分配一个大的切片来保存所有像素
	pixels := make([]uint8, XSize*YSize) // 拥有类型 []uint8，尽管图片是 [][]uint8.
	// 遍历行，从剩余像素切片的前面切出每行来。
	// 巧妙
	for i := range picture {
		picture[i], pixels = pixels[:XSize], pixels[XSize:]
	}

	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, 4, 5, 6)
	x = append(x, y...)
	fmt.Println(x)

	return picture

}
