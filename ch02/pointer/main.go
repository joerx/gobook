package main

import (
	"fmt"
)

func main() {
	i := 1
	p := &i

	fmt.Printf("i:\t\t%v\n", i)  // 1
	fmt.Printf("*p:\t\t%v\n", *p) // 1

	fmt.Printf("p:\t\t%v\n", p)  // address of i (0xc82...)
	fmt.Printf("&i:\t\t%v\n", &i) // address of i (0xc82...)
	fmt.Printf("&p:\t\t%v\n", &p) // address of p

	fmt.Printf("p==&i:\t%v\n", p == &i) // true => same address
	//fmt.Println(&p == &i) // error! => different types (**int vs *int)

	*p++
	fmt.Println("// *p++")
	fmt.Printf("i:\t\t%v\n", i) // 2

	incByVal(i)
	fmt.Println("// incByVal(i)")
	fmt.Printf("i:\t\t%v\n", i) // 2

	incByRef(&i) // pass pointer to i
	fmt.Println("// incByRef(&i)")
	fmt.Printf("i:\t\t%v\n", i) // 3

	p2 := mkPointer()
	p3 := mkPointer()

	fmt.Printf("p2:\t\t%v\n", p2)		// address of v
	fmt.Printf("p3:\t\t%v\n", p3)   // address of another v
	fmt.Printf("p2==p3:\t%v\n", p2 == p3) // false, different addresses

	fmt.Printf("*p2:\t%v\n", *p2) 	// value of v (10)

	incByRef(p2)
	fmt.Println("// incByRef(p2)")
	fmt.Printf("*p2:\t%v\n", *p2) 	// value of v (10)

	incByVal(*p2) // pointer to p2 (int)
	fmt.Println("// incByVal(p2)")
	fmt.Printf("*p2:\t%v\n", *p2)  // still 11

	p4 := &i
	assign(p4, 42)

	fmt.Println("// assign(p4, 42)")

	fmt.Printf("*p4:\t%v\n", *p4) // 42
	fmt.Printf("i:\t\t%v\n", i) // 42
}

func incByVal(val int) {
	val++
}

// func takes a pointer to an integer
func incByRef(ref *int) {
	*ref++ // increment value referenced by ref
}

func assign(ref *int, val int) {
	*ref = val
}

func mkPointer() *int {
	v := 10
	return &v
}
