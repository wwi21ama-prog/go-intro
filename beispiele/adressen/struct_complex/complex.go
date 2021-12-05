package main

import (
	"fmt"
)

// Datentyp für komplexe Zahlen.
type Complex struct {
	realpart, imgpart float64
}

func (c *Complex) print() {
	fmt.Printf("x == %v + %vi\n", c.realpart, c.imgpart)
}

func multiply(a, b Complex) Complex {
	// (a + bi)(c + di) = ac + adi + cbi - bd = (ac - bd) + (ad + cb)i
	return Complex{a.realpart*b.realpart - a.imgpart*b.imgpart, a.realpart*b.imgpart + a.imgpart*b.realpart}
}

// Definiere eine "Methode" für den Typ Complex
func (a *Complex) mult(b Complex) {
	// (a + bi)(c + di) = ac + adi + cbi - bd = (ac - bd) + (ad + cb)i
	a.realpart, a.imgpart = a.realpart*b.realpart-a.imgpart*b.imgpart, a.realpart*b.imgpart+a.imgpart*b.realpart
}

func main() {
	//c1 := Complex{2.4, 1.3}
	//c2 := Complex{1.0, 0.0}
	c3 := Complex{0.0, 1.0}
	//c4 := Complex{2.0, 1.0}
	//print(c1)
	//print(multiply(c2,c3))
	//print(multiply(c3,c4))
	//print(multiply(c3,c3))

	c3.mult(c3)
	c3.mult(c3)
	c3.print()
}
