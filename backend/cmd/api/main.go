package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	panic("runtime errorだよ")
	fmt.Println("処理されない")
}
