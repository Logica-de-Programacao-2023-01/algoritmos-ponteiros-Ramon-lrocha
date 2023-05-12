package main

import "fmt"

func main() {
	valor := true
	inverter(&valor)

	fmt.Println(valor)

}
func inverter(b *bool) {
	*b = !*b
}
