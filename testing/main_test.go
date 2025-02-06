package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	/* total := Suma(10, 20)

	if total != 30 {
		t.Errorf("Suma es incorrecta %d debe de ser %d", total, 30)
	} */
	cases := []struct {
		a int
		b int
		n int
	}{
		{10, 20, 30},
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, item := range cases {
		total := Suma(item.a, item.b)
		if total != item.n {
			t.Errorf("La suma de %d + %d debe de ser %d", item.a, item.b, item.n)
		}
	}

}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{5, 5, 5},
		{1, 0, 1},
	}
	for _, item := range tables {
		valueOfTest := GetMax(item.a, item.b)
		if valueOfTest != item.c {
			t.Errorf("Los valores %d y %d el mayor es: %d", item.a, item.b, item.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	cases := []struct {
		a int
		b int
	}{
		{1, 0},
		{2, 1},
		{50, 1},
	}

	for _, value := range cases {
		fiboValue := Fibonacci(value.a)
		if fiboValue != value.b {
			t.Errorf("El valor fibonacci de %d no es %d debe de ser %d", value.a, fiboValue, value.b)
		}
	}
}
