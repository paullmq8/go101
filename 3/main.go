package main

import "fmt"

func main() {
	println("hi!")
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()
	defer func() {
		println("exit normally")
	}()
	panic("bye!")
	println("Unreachable")
}
