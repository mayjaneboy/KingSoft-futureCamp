package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
}
