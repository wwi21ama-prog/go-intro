package main

import (
  "fmt"
  "errors"
)

func main() {

  i1, i2 := swap(3,4)

  fmt.Println(i1, i2)
  fmt.Println(swap(15,73))
  fmt.Println(get("Hallo",2))

  r,e := div(3,2)

  if e == nil {
    fmt.Println(r)
  } else {
    fmt.Println(e)
  }

}

// Eine Funktion kann nicht nur einen Wert liefern, sondern auch mehrere.
func swap(x,y int) (int, int) {
  return y,x
}

// Liefert den String ab Stelle i in s und dessen Länge.
func get(s string, i int) (string, int) {
  return s[i:], len(s[i:])
}

// Liefert ein Divisionsergebnis und einen Fehler, falls durch Null geteilt wird.
func div(x,y int) (int, error) {
  if y == 0 {
    return 0, errors.New("Division durch Null")
  }
  return x / y, nil
}

// Führt die komplexe Addition
func addComplex(r1, i1, r2, i2 float32) (float32, float32) {
  return r1 + r2, i1 + i2
}