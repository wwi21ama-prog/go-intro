package main

import "fmt"

func main() {

	// Beispiel für eine Funktion, die prüft, ob eine Liste sortiert ist.
	// Die Funktion hat dabei nicht hartcodiert, ob auf- oder absteigend, sondern sie
	// erhält diese Information in Form einer anderen Funktion, die ihr als
	// Parameter übergeben wird.

	l1 := []int{1, 3, 5, 7, 9}
	l2 := []int{1, 3, 9, 7, 5}
	l3 := []int{9, 7, 5, 3, 1}

	fmt.Println(sorted(l1, lessThan))
	fmt.Println(sorted(l2, lessThan))
	fmt.Println(sorted(l3, lessThan))

	fmt.Println(sorted(l1, greaterThan))
	fmt.Println(sorted(l2, greaterThan))
	fmt.Println(sorted(l3, greaterThan))

}

// Prüft, ob x kleiner als y ist und liefert dann true, sonst false.
func lessThan(x, y int) bool {
	return x <= y
}

// Prüft, ob x größer als y ist und liefert dann true, sonst false.
func greaterThan(x, y int) bool {
	return x >= y
}

// Diese Funktion erwartet eine Liste von Zahlen
// liefert true, falls die Liste sortiert ist.
// Das Sortierkriterium wird durch die als Parameter
// erwartete Funktion compare() angegeben.
func sorted(list []int, compare func(x, y int) bool) bool {

	if len(list) < 2 {
		return true
	}

	if compare(list[0], list[1]) {
		return sorted(list[1:], compare)
	}
	return false
}
