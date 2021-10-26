package main

import "fmt"

func main() {
  
  // Array: Zusammenhängender Speicherbereich mit fester
  //        Größe, alle Elemente haben den selben Typ.

  // Definition eines int-Arrays
  var a1 [5]int
  fmt.Println(a1)

  // Das Element an Stelle 2  verändern:
  a1[2] = 42
  fmt.Println(a1)

  // Das Element an Stelle 3 ausgeben.
  fmt.Println(a1[3])

  // Das Array in einer Schleife ausgeben (V1):
  for i:= 0; i<len(a1); i++ {
    fmt.Println(a1[i])
  }

  // Das Array in einer Schleife ausgeben (V2):
  // Range-For-Scheife:
  // In jedem Durchlauf ist i die Position und v der Wert.
  for i,v := range(a1) {
    fmt.Printf("%v: %v\n", i, v)
  }

  // Wie eben, nur dass uns die Position nicht interessiert.
  for _,v := range(a1) {
    fmt.Printf("%v\n", v)
  }

  // Noche eine Range-Schleife, aber mit anderen Namen.
  for index, value := range(a1) {
    fmt.Printf("%v: %v\n", index, value)
  }

  // Neues, längeres Array
  var a2 [10]int

  // Nur den Teil von Position 2 - 5 durchlaufen.
  for index, value := range(a2[2:6]) {
    fmt.Printf("%v: %v\n", index, value)
  }

  // Eine Slice definieren.
  // Das ist eine Sicht auf einen Ausschnitt eines Arrays.
  s1 := a2[2:6]
  fmt.Println(s1)

  // Slices können verwendet werden wie Arrays.
  // Wenn man eine Slice verändert, verändert sich das Array auch.
  s1[2] = 23
  fmt.Println(s1) 
  fmt.Println(a2)

  // Eine Slice mit Hilfe der Funktion make erzeugen
  s2 := make([]int, 3)
  fmt.Println(s2)
  fmt.Println(len(s2))

  s2 = append(s2, 3,4,5,6,7)
  fmt.Println(s2)



  fmt.Println(digitSum(42))
  fmt.Println(digitList(42))
}

//Quersumme
func digitSum(x int) int {
  
  result := 0
  for ; x != 0 ; {
    result += x % 10
    x = x / 10
  }

  return result
}

// Liefert eine Liste der Ziffern von x.
func digitList(x int) []int {
  result := make([]int, 0)

  // Schleife, die die Ziffern von x berechnet und mittels append an result anhängt.
  for ; x != 0 ; {
    result = append(result, x % 10)
    x = x / 10
  }

  return result
}




// https://www.arndt-bruenner.de/mathe/scripts/pruefziffern.htm 