package main

import "fmt"

func main() {

  // Eine Slice erstellen und befüllen.
  a1 := make([]string, 5)
  for i:=0; i<len(a1); i++ {
    a1[i] = "Test"
  }

  // Range-Schleife, die über die Slice iteriert
  // und alle Elemente ausgibt.
  for _,v := range(a1) {
    fmt.Println(v)
  }

  // Die selbe Schleife mit einem Zähler.
  // Die range-Schreibweise ist eine Abkürzung für
  // dieses hier:
  for i := 0; i<len(a1); i++ {
    v:= a1[i]   // Hier entsteht das v von oben.
    fmt.Println(i,v)
  }

  fmt.Println(a1)
}