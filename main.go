package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var x int

	y := 10

	fmt.Println("Hola", x, y)

	value, err := strconv.ParseInt("100", 0, 64)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}

	myMap := make(map[string]int)

	myMap["Richi"] = 33
	myMap["Ivon"] = 31
	myMap["Marco"] = 29

	fmt.Println(myMap)

	myslice := []int{100, 200, 300}
	fmt.Println(myslice, len(myslice), cap(myslice))

	for i, value := range myslice {
		fmt.Println(i, value)
	}
	myslice = append(myslice, 50)
	fmt.Println(myslice)

}
