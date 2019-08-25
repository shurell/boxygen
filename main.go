package main

import "fmt"

type boxElement struct {
	width     float32
	height    float32
	shipLen   float32
	thickness float32
}

type boxTorets struct {
	boxElement
	rr float32
}

func main() {
	fmt.Println("БОКСИГЕН. версия 0.00001")
}
