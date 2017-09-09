package main

import "fmt"

type node struct {

	data int

}

func(p* node)set(val int)() {

	p.data = val
}

func(p* node)get() int {

	return p.data;
}

func main() {

	n := node{data: 10}

	m := &n
	m.set(12)
	fmt.Println(m.get())
}