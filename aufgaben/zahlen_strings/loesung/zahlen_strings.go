package main

import (
	"fmt"
)

// Erwartet eine Zahl n und liefert die entsprechende natürlichsprachliche Darstellung.
// z.B. 42 -> "zweiundvierzig"
//     123 -> "einhundertdreiundzwanzig"
func numToString(n int) string {

	switch n {
	case 0:
		return "null"
	case 11:
		return "elf"
	case 12:
		return "zwölf"
	}

	digitList := split(numToList(n))
	return numToString3(digitList[0])
}

// Erwartet eine Liste von 3 Ziffern und liefert
// die String-Darstellung dieser 3 Ziffern.
func numToString3(digits []int) string {

	digitstrings := []string{"", "ein", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}
	digitstrings10 := []string{"", "zehn", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"}

	d0, d1, d2 := digits[0], digits[1], digits[2]
	digit100 := digitstrings[d0]
	if d0 > 0 {
		digit100 += "hundert"
	}
	digit1 := digitstrings[d2]
	if d0 == 0 && d1 == 0 && d2 == 1 {
		digit1 += "s"
	}
	digit10 := digitstrings10[d1]

	return digit100 + digit1 + "und" + digit10
}

func main() {
	// Standardfälle
	assertEqual(42, "zweiundvierzig")
	assertEqual(153, "einhundertdreiundfünfzig")
	assertEqual(87, "siebenundachtzig")
	assertEqual(45, "fünfundvierzig")
	assertEqual(98, "achtundneunzig")
	assertEqual(24, "vierundzwanzig")
	assertEqual(38, "achtunddreißig")

	// Sonderfälle
	assertEqual(0, "null")
	assertEqual(11, "elf")
	assertEqual(12, "zwölf")
}

// Erwartet eine einzelne Zahl und liefert eine Liste der Ziffern dieser Zahl.
func numToList(n int) []int {
	result := make([]int, 0)
	for ; n != 0; n /= 10 {
		result = append(result, n%10)
	}

	return reverse(result)
}

// Erwartet eine Liste von Zahlen und dreht sie um.
// Liefert die umgedrehte Liste.
func reverse(list []int) []int {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// Erwartet eine Liste von Zahlen und füllt sie mit Nullen bis zur angegebenen Länge auf. Die Nullen werden vorne angehängt. Liefert die aufgefüllte Liste.
func fill(list []int, size int) []int {
	for i := len(list); i < size; i++ {
		list = append([]int{0}, list...)
	}
	return list
}

// Erwartet eine Liste von Zahlen und splittet sie in Blöcke der angegebenen Länge auf. Liefert eine (zweidimensionale) Liste von Blöcken, die selbst jeweils size Zahlen enthalten.
func split(list []int) [][]int {
	for len(list)%3 != 0 {
		list = fill(list, len(list)+1)
	}

	result := make([][]int, 0)
	for i := 0; i < len(list); i += 3 {
		result = append(result, []int{list[i], list[i+1], list[i+2]})
	}
	return result
}

// Hilfsfunktion zum Testen von numToString.
// Gibt eine Fehlermeldung aus, wenn numToString for input nicht das Ergebnis expected hat.
func assertEqual(input int, expected string) {
	result := numToString(input)

	if result != expected {
		fmt.Printf("Fehler: numToString(%v) == \"%v\", erwartet war aber \"%v\".\n", input, result, expected)
	}
}
