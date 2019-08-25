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

func genBoxy(sw, sh, sd, tm, sl, rr float32) {

	w := sw - tm
	h := sh - tm
	d := sd - tm

	bottom := genDno(w, d)

}

func genDno(w, h float32) boxElement {
	//
	theDno := new(boxElement)
	theDno.width = w
	theDno.height = h

	return *theDno
}

func genLongSide(w, h float32) boxElement {

}

func genTorets(w, h, rr float32) boxTorets {

}
