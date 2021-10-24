package main

import "fmt"

func main() {
	text := "Hello World"
	fmt.Println(text)

	// Aufruf der Funktion plus2
	plus2(77) // Ergebnis wird nicht verwendet, Aufruf hat keinen Effekt.

	// Richtig ist:
	x := plus2(3) // Funktion aufrufen und Ergebnis in x speichern
	fmt.Println(x)

	testPlus2()

	//  fmt.Printf("Ergebnis: %v", x)
}

// Definition einer Funktion
// Name: "plus2"
// Parameter: Eine Zahl x vom Typ int
// - wird beim Aufruf der Funktion eingesetzt
// Rückgabetyp: Eine Zahl
func plus2(x int) int { // "Signatur"
	result := x + 2
	fmt.Println("func plus2 wurde ausgeführt")
	return result // "return": Die Funktion ist zu Ende, das Ergebnis ist x+2
}

func checkPlus2(x int, expected int) {
  actual := plus2(x)
	if actual != expected {
		fmt.Printf("Fehler, plus2(%v) sollte %v liefern, tatsächlich wurde aber %v geliefert.\n", x, expected, actual)
	}
}

func testPlus2() {
	checkPlus2(3, 5)
	checkPlus2(-15, -13)

}

// "Fehler, plus(40) sollte 42 liefern, es wurde aber 43 geliefert."
