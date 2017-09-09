package main

import "fmt"

func show(data int) {

	var index int
	index = 0

	for {

		if(index >= data) {

			break
		}

		fmt.Println(index)
		index ++
		continue

	}
}

func main() {

	show(10)
}