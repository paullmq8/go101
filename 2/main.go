package main

import "go101/use"

func init() {
	println("in 2 pkg")
}

func main() {
	println("running main()")
	use.Pr()
}
