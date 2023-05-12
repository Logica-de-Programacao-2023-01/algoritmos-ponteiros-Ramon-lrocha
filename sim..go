package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Printf("Antes da troca - a: %d, b: %d\n", a, b)

	troca(&a, &b)

	fmt.Printf("depois da troca - a: %d, b: %d\n", a, b)

}
func troca(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp

}
