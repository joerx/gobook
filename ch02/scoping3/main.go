package main

import (
	"fmt"
	"os"
)

const fname = "~/.bashrc"

func main() {
	doesCompile()
}

func doesCompile() {
	// would not compile: f is scoped inside the if statement but not used
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, 1024)
	f.Read(bs)
	f.Close()
}

// the below code would not compile: f is scoped inside the if statement but not used
func doesNotCompile() {
	// if f, err := os.Open(fname); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // f is not visible here
	// f.ReadByte()
	// f.Close()
}
