package main

import "fmt"

func main() {

  fmt.Println(digitSum(103445))
}


// "einfache" Quersumme
func simpleDigitSum(x int) int {
  
  result := 0
  for ; x != 0 ; {
    result += x % 10
    x = x / 10
  }

  return result
}

// Hilfsfunktion: Liefert eine Liste der Ziffern von x.
func digitList(x int) []int {
  result := make([]int, 0)

  // Schleife, die die Ziffern von x berechnet und mittels append an result anhängt.
  for ; x != 0 ; {
    result = append(result, x % 10)
    x = x / 10
  }

  return result
}

// Hilfsfunktion: Berechnet die Summe der Elemente einer Slice.
func sliceSum(s []int) int {
  result := 0
  
  for _, v := range(s) {
    result += v
  }

  return result
}

// "richtige" Quersumme
func digitSum(x int) int {
  return sliceSum(digitList(x))
}
