package main

func soma(a, b *int) int {
	*a = 50
	*b = 50

	return *a + *b
}

func main() {
	minhaVar := 10
	minhaVar2 := 20
	soma(&minhaVar, &minhaVar2)

	println(minhaVar)
	println(minhaVar2)
}
