package main

import (
	"fmt"
	"github.com/xishvai/lunar"
)

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

	y, m, d := lunar.GanZhiYMD(2014, 5, 1) //should be:甲午 戊辰 壬申
	fmt.Println("2014.5.1", y, m, d)

	y, m, d = lunar.GanZhiYMD(2014, 5, 5) //should be:甲午 己巳 丙子
	fmt.Println("2014.5.5", y, m, d)

	a := lunar.AnimalYear(1900)  //should be 鼠
	a2 := lunar.AnimalYear(1988) //should be 龙
	fmt.Println(a, a2)

	z := lunar.ZhiHour(0)
	z1 := lunar.ZhiHour(23)
	z2 := lunar.ZhiHour(1)
	z3 := lunar.ZhiHour(12)
	z4 := lunar.ZhiHour(22)
	fmt.Println(z, z1, z2, z3, z4)
}
