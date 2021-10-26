package main

import "fmt"

func main() {
  // Definieren eines 2D-Arrays
  var spielfeld [3][3]string

 // Jede Zeile einzeln ausgeben
  for zeile := 0; zeile < 3; zeile++ {
    fmt.Println(spielfeld[zeile])
  } 

  // Überall ins Spielfeld Leerzeichen schreiben.
  // Schleife, die über die Zeilen iteriert
  // Innerhalb: Schleife, die über jedes Element der Zeile iteriert. Diese Schleife muss Leerzeichen schreiben.
  spielfeld = fillBoard(spielfeld, "A")


  // Ein X in Zeile 2, Spalte 0 setzen.
  spielfeld[2][0] = "X"

  // Das Spielfeld ausgeben:
  fmt.Println(spielfeld)

  // Die Länge des Elements an Stelle (0,0) bestimmen:
  fmt.Println(len(spielfeld[0][0]))
}

func fillBoard(spielfeld [3][3]string, s string) [3][3]string {
  for zeile := 0; zeile < 3; zeile++ {
    for spalte := 0; spalte < 3; spalte++ {
      spielfeld[zeile][spalte] = s
    }
  }
  return spielfeld
}