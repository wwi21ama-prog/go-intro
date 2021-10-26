ackage main

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


  var a2 [10]int

  for index, value := range(a2[2:6]) {
    fmt.Printf("%v: %v\n", index, value)
  }



  fmt.Println(digitSum(42))
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

// https://www.arndt-bruenner.de/mathe/scripts/pruefziffern.htm 