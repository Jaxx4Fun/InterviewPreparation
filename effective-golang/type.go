package effective

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

func aboutType(val interface{}) {
	if str, ok := val.(string); ok {
		fmt.Printf(str)
	} else {
		fmt.Printf("value is not a string")
	}

	//conversion
	type Sequence []int
	s := Sequence{3, 2, 1}
	fmt.Print([]int(s))

	//
	sort.IntSlice(s).Sort()
}

// 简单的计数器服务。
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

func aboutEmbedding() {
	type Job struct {
		Command string
		*log.Logger
	}
	job := Job{"a", log.New(os.Stdout, "stdout", log.LstdFlags)}
	job.Printf("starting now...")
}
