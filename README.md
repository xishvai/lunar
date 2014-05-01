lunar
=====

Chinese Lunar Calendar Package Written by Go.

About
=====

	//Solar structure
	type Solar struct {
		time.Time
	}
	
	//Luanr structure
	type Lunar struct {
		year   int
		month  int
		day    int
		hour   int
		minute int
		second int
	}

`func NewSolar(year, month, day, hour, min, sec int) *Solar`

`func NewSolarNow() *Solar`

`func NewLunar(year, month, day, hour, min, sec int) *Lunar`

`func NewLunarNow() *Lunar`

Lunar or Solar has a method `Convert` to convert itself to the *opposite* one.

`func (s *Solar) Convert() *Lunar`

`func (l *Lunar) Convert() *Solar`

NOTICE
======
This package's year range is `[1900,2050]` and month range is `[1,12]`.

Example
=======
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