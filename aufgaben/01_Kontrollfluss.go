// Aufgaben zu Funktionen, Schleifen und If-Then-Else
package main

import "fmt"

// Liefert die Summe aller Vielfachen von x, die kleiner als n sind.
func sumMultiples(x, n int) int {
	result := 0
	// TODO: Hier das Ergebnis berechnen.
	return result
}

// Tests für sumMultiples.
func testSumMultiples() {
	expectEqual(sumMultiples(3, 10), 18)
	expectEqual(sumMultiples(2, 10), 20)
	expectEqual(sumMultiples(10, 2), 0)
}

// Liefert genau dann true, wenn n eine Primzahl ist.
func isPrime(n int) bool {
	// TODO: Das pauschale Return durch etwas sinnvolles ersetzen.
	return true
}

// Tests für isPrime.
func testIsPrime() {
	expectEqual(isPrime(1), false)
	expectEqual(isPrime(2), true)
	expectEqual(isPrime(3), true)
	expectEqual(isPrime(4), false)
	expectEqual(isPrime(5), true)
}

// Liefert den größten Primfaktor der Zahl n.
func largestPrimeFactor(n int) int {
	// TODO: Hier das Ergebnis berechnen.
	return 0
}

// Tests für largestPrimeFactor.
func testLargestPrimeFactor() {
	expectEqual(largestPrimeFactor(10), 5)
	expectEqual(largestPrimeFactor(49), 7)
	expectEqual(largestPrimeFactor(60), 5)
	expectEqual(largestPrimeFactor(17), 17)
	expectEqual(largestPrimeFactor(1), 0)
	expectEqual(largestPrimeFactor(0), 0)
}

// Liefert das kleinste gemeinsame Vielfache von m und n.
func lcm(m, n int) int {
	// TODO: Hier das Ergebnis berechnen.
	return 0
}

// Tests für lcm
func testLcm() {
	expectEqual(lcm(3, 5), 15)
	expectEqual(lcm(2, 5), 10)
	expectEqual(lcm(4, 10), 20)
	expectEqual(lcm(20, 5), 20)
	expectEqual(lcm(25, 10), 50)
}

// Hilfsfunktion, wird in den Tests benutzt
func expectEqual(value, expected interface{}) {
	if value != expected {
		fmt.Printf("Fehler: %v bekommen, erwartet war aber %v.\n", value, expected)
	}
}

// Main-Funktion, ruft die Tests auf.
// Kommentieren Sie am Besten die Testfunktionen aus, die Sie gerade nicht benötigen.
func main() {
	testSumMultiples()
	testIsPrime()
	testLargestPrimeFactor()
	testLcm()
}
