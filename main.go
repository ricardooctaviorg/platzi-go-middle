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

}
