package main

import (
	"fmt"
	"time"
)

func main() {
	t := <-time.After(1 * time.Second)
	fmt.Println(t)
}
