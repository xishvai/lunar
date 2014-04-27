package main

import "github.com/xishvai/lunar"
import "fmt"

func main() {
	s := lunar.NewSolarNow()
	fmt.Printf("%v", s)
}
