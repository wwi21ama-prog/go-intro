package main

import (
	"fmt"
)

// Eine variadische Funktion erwartet eine beliebige Anzahl an Parametern gleichen Typs:
func printAll(numbers ...int) {
	// Innerhalb stehen diese Parameter als Slice zur Verfügung:
	fmt.Printf("Typ von numbers: %T\n", numbers)

	// Alle Parameter ausgeben
	for i, v := range numbers {
		fmt.Printf("%v: %v\n", i, v)
	}
}

func main() {

	// Mehrere Argumente in die Funktion geben.
	printAll(42, 23, 103)

	// Eine Slice erzeugen und dann deren Elemente einzeln ("ausgepackt") in die Funktion geben.
	l1 := []int{100, 200, 300, 400, 500}
	printAll(l1...)
}
