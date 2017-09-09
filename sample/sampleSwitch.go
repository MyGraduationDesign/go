package main

import "fmt"

func test(a int) {

	switch (a) {

	case 1:
		fmt.Println("1")

	case 2:
		fmt.Println("2")

	default:
		fmt.Println("error")

	}

}

func main() {

	test(1)
	test(2)
	test(3)
}