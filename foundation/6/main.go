package main

import "fmt"

const a = "Hello"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Matheus"
	e float64 = 3.4
	f ID      = 1
)

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 12)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
