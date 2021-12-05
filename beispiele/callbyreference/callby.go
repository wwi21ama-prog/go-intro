package main

import (
	"fmt"
)

func beispiel_referenzen_slices() {

	// Definieren eine Int-Slice.
	l1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(l1)

	// Definiere einen Ausschnit aus l1.
	l2 := l1[2:5]
	fmt.Println(l2)

	// Ändere einen Wert in l2
	// Beobachtung: Der entsprechende Wert in l1 ändert sich mit.
	l2[1] = 99999
	fmt.Println(l2)
	fmt.Println(l1)
	// l2 hat eine "Referenzsemantik".
	// l2 ist keine eigenständige Variable, sonder nur ein
	// weiterer Name für einen anderen Speicherbereich.
}

func beispiel_pointer() {

	// Definieren eine int-Variable.
	x := 42
	fmt.Printf("Wert von x: %v\n", x)

	// Definiere einen "Pointer" ("Zeiger") auf x.
	x_ptr := &x                                  // &x ist die Adresse von x im Speicher
	fmt.Printf("Adresse von x_ptr: %v\n", x_ptr) // Speicheradresse von x wird ausgegeben
	fmt.Printf("Wert von x_ptr: %v\n", *x_ptr)   // Wert von x wird ausgegeben

	// *x_ptr ist der Wert, der sich an der Adresse befindet, auf die x_ptr zeigt ("Dereferenzierung").

	// Benutze x_ptr, um x zu verändern:
	*x_ptr = 99999
	fmt.Printf("Wert von x, nachdem *x_ptr geändert wurde: %v\n", x)

	// x_ptr hat eine Referenzsemantik.
}

/*
// Funktion, die einen Parameter erwartet, und diesen verändert.
func quadriere(x_in_quadriere int) {
  x_in_quadriere = x_in_quadriere*x_in_quadriere
}
Das hier funktioniert noch nicht, weil x als Kopie in die Funktion gegeben wird.
Die Funktion nutzt "Call-By-Value".
*/

// Funktionierende Variante von quadriere():
// Erwartet statt des Werts von x die Adresse einer zu quadrierenden Zahl.
// Intern wird mit * dereferenziert, um die Berechnung auszuführen.
//
// Die Funktion nutzt "Call-By-Reference", sie hat einen "Seiteneffekt".

func quadriere(x_ptr *int) { // x_ptr ist vom Typ "Pointer auf int"
	*x_ptr = *x_ptr * *x_ptr // Der Wert, auf den x_ptr zeigt (*x_ptr), soll *x_ptr Mal *x_ptr sein.
}

func main() {
	//beispiel_referenzen_slices()
	//beispiel_pointer()
	x_main := 12
	quadriere(&x_main) // Übergebe die Adresse von x_main and die Funktio quadriere().
	fmt.Println(x_main)
}
