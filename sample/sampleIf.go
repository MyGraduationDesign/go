package main

import "fmt"

func compare(a int, b int) {

	if(a > b) {
		fmt.Println("greater")
	}else{
		fmt.Println("smaller")
	}
}

func main() {

	compare(3, 2)
}