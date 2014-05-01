package main

import "github.com/xishvai/lunar"
import "fmt"

func main() {
	l0 := lunar.NewLunar(1988, 2, 11, 9, 9, 9)
	fmt.Println(l0)
	s0 := l0.Convert()
	fmt.Println(s0)

	s1 := lunar.NewSolar(1988, 3, 28, 9, 9, 9)
	fmt.Println(s1)
	l1 := s0.Convert()
	fmt.Println(l1)

	s := lunar.NewSolarNow()
	fmt.Println(s)

	l := lunar.NewLunarNow()
	fmt.Println(l)

	s2l := s.Convert()
	fmt.Println(s2l)

	l2s := l.Convert()
	fmt.Println(l2s)
}
