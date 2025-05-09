package main

import "fmt"

func GetValue() int {
	return 1
}

func main() {
	s1 := []int{3, 4, 5, 6}
	fmt.Println("len s1", len(s1))
	fmt.Println("cap s1", cap(s1))

	s2 := s1[1:3]
	fmt.Println("len s2", len(s2))
	fmt.Println("cap s2", cap(s2))

	fmt.Println("s2", s2)
}
