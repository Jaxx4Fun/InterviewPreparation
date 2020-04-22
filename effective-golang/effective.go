package effective

import (
	"fmt"
	"os"
)

func allInit() {
	// f1 := new(os.File)
	// f2 := &os.File{}

	// f3 := &os.File{fd: 1, name: "filename"}
	// f4 := &os.File{1, "filename", nil, 0}
}

func aboutPrint() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
}
