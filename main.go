package main

import (
	"fmt"
	"math"
)

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

func (elem *boxElement) buildElem(s string) [][]float32 {
	verticles := [][]float32{{0}, {0}} // стартовая точка
	if s == "bottom" {
		fmt.Println("Дно бокса")
		qtySpykeW := math.Round(float64(elem.width / elem.shipLen))
		qtySpykeH := math.Round(float64(elem.height / elem.shipLen))
		fmt.Println("Доступное кол-во шипов по длинной стороне:", qtySpykeW, ":: по короткой стороне:", qtySpykeH)

	}
	if s == "side" {
		fmt.Println("Стенка бокса")
	}
	fmt.Println("Рассчитываем вершины", verticles)

	return verticles
}

func (elem *boxTorets) buildTorets(s string) {
	if s == "blank" {
		fmt.Println("Глухой торец бокса")
	}
	if s == "thumb" {
		fmt.Println("Торец бокса с вырезом")
	}
	fmt.Println("Рассчитываем вершины")
}

func main() {
	fmt.Println("БОКСИГЕН. версия 0.00001")

	genBoxy(96.0, 50.0, 57.0, 3.0, 15.0, 20.0)
}

func checkSpykeLen(v, s float32) bool {
	if (v / s) <= 2.4 {
		fmt.Println("Слишком большой размер шипа!")
		return true
	}
	return false

}

func genBoxy(sw, sh, sd, tm, sl, rr float32) {

	if checkSpykeLen(sw, sl) || checkSpykeLen(sh, sl) || checkSpykeLen(sd, sl) {
		fmt.Println("ПРОБЛЕМА!")
	}

	w := sw - tm
	h := sh - tm
	d := sd - tm

	bottom := genDno(w, d, sl)
	oneLong := genLongSide(w, h, sl)
	twoLong := genLongSide(w, h, sl)
	oneTor := genTorets(h, d, 0)
	twoTor := genTorets(h, d, rr)

	fmt.Println("СОЗДАНЫ СТОРОНЫ:")
	fmt.Println(bottom, oneLong, twoLong, oneTor, twoTor)

	bottom.buildElem("bottom")

}

func genDno(w, h, sl float32) boxElement {
	//
	theDno := new(boxElement)
	theDno.width = w
	theDno.height = h
	theDno.shipLen = sl

	return *theDno
}

func genLongSide(w, h, sl float32) boxElement {
	theLongSide := new(boxElement)
	theLongSide.width = w
	theLongSide.height = h
	theLongSide.shipLen = sl
	return *theLongSide
}

func genTorets(w, h, rr float32) boxTorets {
	theTorets := new(boxTorets)
	return *theTorets
}
