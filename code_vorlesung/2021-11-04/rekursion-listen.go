package main

import "fmt"

func main() {
	s1 := "Hallo Welt"
	s2 := ""

	fmt.Println(isEmpty(s1)) // Soll false ausgeben.
	fmt.Println(isEmpty(s2)) // Soll true ausgeben.

	// Das erste Element von s1 ausgeben.
	fmt.Printf("%c\n", s1[0])
	// Den Rest von s1 ausgeben:
	fmt.Println(s1[1:])
	// fmt.Println(s1[1:len(s1)]) äquivalent

	fmt.Println(count(s1, 'x')) // Soll 0 ausgeben.
	fmt.Println(count(s1, 'a')) // Soll 1 ausgeben.
	fmt.Println(count(s1, 'l')) // Soll 3 ausgeben.

	// Eine int-Slice m,it 5 Nullen erstellen.
	a1 := make([]int, 5)

	// a1 ist sortiert:
	fmt.Println(sorted(a1)) // Soll true ausgeben.

	// An der letzten Stelle eine 5 einfügen:
	a1[4] = 5

	// a1 ist sortiert:
	fmt.Println(sorted(a1)) // Soll true ausgeben.

	// An der ersten Stelle eine 5 einfügen:
	a1[0] = 5

	// a1 ist nicht sortiert:
	fmt.Println(sorted(a1)) // Soll false ausgeben.
}

// Soll true liefern, wenn s leer ist.
func isEmpty(s string) bool {
	return len(s) == 0
}

// Liefert die Anzahl der Vorkommen von c in s.
func count(s string, c byte) int {
	//   0, falls s leer ist
	if isEmpty(s) {
		return 0
	}
	//   1 + count(s[1:], c), falls s[0] == c
	if s[0] == c {
		return 1 + count(s[1:], c)
	}
	//   0 + count(s[1:], c), sonst
	return 0 + count(s[1:], c)
}

// Liefert true, wenn c in s vorkommt.
func contains(s string, c byte) bool {
	// Hinweis: Ganz ähnlich wie count, evtl. sogar etwas einfacher.
	if isEmpty(s) {
		return false
	}
	return contains(s[1:], c)
}

// Liefert true, falls die Liste aufsteigend sortiert ist.
func sorted(list []int) bool {
	/* Hinweis:
	 * Eine Liste ist sortiert, wenn sie leer ist,
	 * wenn sie die Länge 1 hat ...
	 */
	if len(list) < 2 {
		return true
	}

	/*
	 * ... oder wenn die beiden ersten Elemente
	 * richtig herum sind und die Liste ab Stelle
	 * 2 auch noch sortiert ist.
	 */
	if list[0] > list[1] {
		return false
	}
	return sorted(list[1:])
}
