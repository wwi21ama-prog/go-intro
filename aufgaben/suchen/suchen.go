package main

import "fmt"

func main() {
	l1 := []int{1, 3, 5, 2, 4}
	fmt.Println(find(l1, 2)) // Soll 3 ausgeben
	fmt.Println(find(l1, 7)) // Soll -1 ausgeben
}

// Liefert die Position von x in list, falls x darin vorkommt,
// liefert ansonsten -1
func find(list []int, x int) int {
	// TODO
	return -1
}

// Aufgabe 1:
// Schreiben Sie die Funktion find
// * als lineare Suche
// * als binäre Suche
// * jeweils iterativ (d.h. mit Schleife) und rekursiv
//   Hinweis: Die iterative binäre Suche ist
//            verhältnismäßig kompliziert!

// Aufgabe 2:
// * Suchen Sie im Internet nach "error" im Zusammenhang mit Go
//   und ändern Sie die Funktion(en), so dass sie die Standard-
//   Fehlerbehandlung in Go machen.
