package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(z int) (a int, b int, c int) {
	return z * 1, z * 2, z * 3
}

func getValues2(z int) (valor1 int, valor2 int, valor3 int) {
	valor1 = z * 2
	valor2 = z * 3
	valor3 = z * 4
	return

}

func main() {

	names := []string{"Hola", "Mundo"}

	fmt.Println("suma", sum(1, 2, 3, 4, 5, 6))
	printNames(names...)
	fmt.Println(getValues(10))
	fmt.Println(getValues2(100))
}
