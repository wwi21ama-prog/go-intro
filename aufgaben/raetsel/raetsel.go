package main

import "fmt"

// Aufgabe 1
// Betrachten Sie die folgende Funktion:
func foo(x, y int) int {
	if y > x {
		return x
	}
	return foo(x-y, y)
}

/*
 * 1. Was berechnet diese Funktion?
 *    Erklären Sie in Worten, was diese Funktion macht.
 *    Geben Sie, falls möglich, eine ähnliche bekannte Funktion an.
 *
 * 2. Bestimmen Sie den Definitionsbereich der Funktion.
 *    D.h. für welche Eingaben funktioniert sie und welche sind
 *    als ungültig anzusehen.
 *
 * 3. Verbessern Sie die Funktion, indem Sie den
 *    Definitionsbereich maximieren.
 */

// Aufgabe 2
// Betrachten Sie die folgende Funktion:
func bar(s string, c1, c2 rune) bool {
	counter := 0
	for _, c := range s {
		if c == c1 {
			counter++
		}
		if c == c2 {
			counter--
		}
		if counter < 0 {
			return false
		}
	}
	return counter == 0
}

/*
 * 1. Was berechnet diese Funktion?
 *    Erklären Sie in Worten, was diese Funktion macht.
 *    Finden Sie ein Beispiel, wo man eine solche Funktion
 *    gebrauchen kann.
 *
 * 2. Schreiben Sie eine rekursive Variante dieser Funktion.
 */

func main() {
	fmt.Println("Dieses Programm macht noch nichts.")
}
