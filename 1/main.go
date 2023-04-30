package main

// should run "go run ." instead of "go run context_switch_test.go"
// cuz there are multiple go files under package main
// if you run it, it will prompt error "undefined: printaa"

func main() {
	printaa()
	var a uint32 = 0
	println(a)
	a, b := 1, 3
	println(a, b)
}
