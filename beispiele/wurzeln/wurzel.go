package main

import "fmt"

func main() {
  fmt.Println(sqrtInt2(4)) // Soll 2 ausgeben
  fmt.Println(sqrtInt2(9)) // Soll 3 ausgeben
  fmt.Println(sqrtInt2(8)) // Soll 2 ausgeben,
                          // 2*2 = 4 <= 8
                          // 3*3 = 9 > 8

  fmt.Println(sqrt(64))
  fmt.Println(sqrt(25))
  fmt.Println(sqrt(8))

}

// Pseudo-Wurzelberechnung:
// Liefert die größte Zahl y,
// so dass y * y <= x
func sqrtInt(x int) int {
  // Möglichkeit 1:
  // Die echte Wurzel berechnen und abrunden/abschneiden.
  // Frage dabei Wie kommt man an die echte Wurzel.

  // Möglichkeit 2:
  // In einer Schleife hochzählen, bis das Quadrat des Zählers nicht mehr kleiner als x ist.

  result := 0
  for i:= 0; i*i <= x; i++ {
    result = i
  }
  return result
}

func sqrtInt2(x int) int {
  i := 0
  for ; i*i <= x; i++ {
  }
  return i-1
}

// Die echte Wurzel
// Newton-Iteration
func sqrt(x float64) float64 {
  
  var z float64 = 1

  for i:=0; i<10; i++ {
    fmt.Println(z)
    z = z - (z*z-x)/(2*z)
  }

  return z
}