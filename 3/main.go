package main

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
	a := "teste"
	a = "hello"
	println(a)
	println(f)
}
