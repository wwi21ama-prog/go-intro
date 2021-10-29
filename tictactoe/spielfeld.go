package main

import "fmt"

func main() {
  spielfeld := makeBoard(3, "*")

  printBoard(spielfeld)

}


func makeBoard(size int, initChar string) [][]string {
  // Definieren einer 2D-Slice
  var board [][]string
  
  for i := 0; i<size; i++ {
    var row []string
    for j := 0; j < size; j++ {
      row = append(row, initChar)
    }
    board = append(board, row)
  }
  return board
}

// Das Spielfeld mit Trennlinien ausgeben.
func printBoard(board [][]string) {
  for i,row := range(board){
    printRow(row)
    if i < len(board) - 1 {
      fmt.Println("---+---+---")
    }
  }
}

// Eine Zeile des Spielfelds mit Trennzeichen ausgeben.
func printRow(row []string) {
  for j := 0; j<len(row); j++ {
    fmt.Print(" "+row[j]+" ")
    if j<len(row)-1 {
      fmt.Print("|")
    }
  }
  fmt.Println()
}
