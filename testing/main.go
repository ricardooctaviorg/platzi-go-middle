package main

func Suma(x, y int) int {
	return x + y
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 0,1,1,2,3,5,8,13
// 1,2,3,4,5,6,7,8
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n > 2 {
		return Fibonacci(n-2) + Fibonacci(n-1)
	}
	return 0
}
