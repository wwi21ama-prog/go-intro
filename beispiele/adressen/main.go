package main

import (
	"fmt"
)

func main() {
	e1 := Entry{"Max", "Mustermann", "Musterstr. 42", "12345 Musterhausen", false}
	e2 := Entry{"Monika", "Musterfrau", "Mustergasse. 38", "54321 Musterstadt", false}

	a1 := AddressBook{e1, e2}
	fmt.Println(a1)

	// Versuch, e1 im Adressbuch zum Favoriten zu machen.
	e1.favourite = true
	// Das klappt nicht, e1 wird geändert, aber nicht seine Kopie in a1.
	fmt.Println(a1)
	fmt.Println(e1)

	// Nächster Versuch, dieses Mal richtig:
	//a1[0].favourite = true
	//fmt.Println(a1)

	// Bessere Variante:
	// Dies nimmt dem Programmierer der main-Funktion die Verantwortung ab.
	a1[0].makeFavourite()
	fmt.Println(a1)
}
