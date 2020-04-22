package effective

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"time"
)

func aboutGoroutines() {

	cs := make(chan *os.File, 100) // 指向文件指针的带缓冲信道
	close(cs)

	//使用
	c := make(chan int /*, 0*/) // 整数类型的无缓冲信道
	go func() {
		sort.Sort(sort.IntSlice([]int{3, 2, 1})) //耗时操作
		c <- 1                                   //发送结束信号，发送后直到接收者收到值
	}()
	// 等待完成
	<-c
}

// haha
const (
	MaxOutstanding = 100
	MinOutstanding = 0
)

// 带缓冲的chan
// 直到数据被复制到缓冲中，sender阻塞结束
// 如果缓冲满，则一直阻塞，直到receiver取走数据，空出位置
// 信道缓冲区的容量决定了同时调用 process 的数量上限
var sem = make(chan int, MaxOutstanding)

type Request int

func process(r *Request) {
	time.Sleep(5 * time.Second)
	fmt.Print(r)
}
func handle(r *Request) {
	sem <- 1   // Wait for active queue to drain.
	process(r) // May take a long time.
	<-sem      // Done; enable next request to run.
}
func Serve(queue chan *Request) {
	for {
		req := <-queue
		go handle(req) // Don't wait for handle to finish. 一直会创建goroutine，即使缓冲已经满
	}
}

func Serve2(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			process(req) // Don't wait for handle to finish.
			<-sem
		}(req) // 循环变量需要作为函数的参数传入
	}
}

func Serve3(queue chan *Request) {
	for req := range queue {
		req := req // 用局部变量屏蔽循环变量；每一次循环都有一个单独对应的实例
		sem <- 1
		go func() {
			process(req) // Don't wait for handle to finish.
			<-sem
		}()
	}
}
func handle4(queue chan *Request) {
	for r := range queue {
		process(r)
	}
}
func Serve4(clientRequests chan *Request, quit chan bool) {
	// Start handlers；启动固定数量的协程
	for i := 0; i < MaxOutstanding; i++ {
		go handle4(clientRequests)
	}
	<-quit // Wait to be told to exit.
}

func aboutGoroutines2() {
	type Request struct {
		args       []int
		f          func([]int) int //
		resultChan chan int        // 接收结果的chan
	}
	sum := func(a []int) (s int) {
		for _, v := range a {
			s += v
		}
		return
	}
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	clientRequests := make(chan *Request, MaxOutstanding)
	// Send request
	clientRequests <- request
	// Wait for response.
	fmt.Printf("answer: %d\n", <-request.resultChan)
}

type Vector []float64

// 将此操应用至 v[i], v[i+1] ... 直到 v[n-1]
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] = u[i] * u[i]
	}
	c <- 1 // 发信号表示这一块计算完成。
}

var NCPU = runtime.NumCPU() // number of CPU cores
func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // Buffering optional but sensible.
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// Drain the channel.
	for i := 0; i < NCPU; i++ {
		<-c // wait for one task to complete
	}
	// All done.
}
