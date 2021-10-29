package main

import "fmt"

func main() {
  spielfeld := makeBoard(3, "*")

  printBoard(spielfeld)

  fmt.Println(checkRows(spielfeld)) // Soll false ausgeben

  spielfeld2 := makeBoard(3, "X")
  fmt.Println(checkRows(spielfeld2)) // Soll true ausgeben

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


// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Zeile drei mal X enthält.
func checkRows(board [][]string) bool {
  // Zeile 0 pruefen
  for _, row := range(board) {
    if checkRow(row) {
      return true // Early Out
    }
  }
  return false
}

// Prüft eine einzelne Zeile, ob sie drei X enthält.
func checkRow(row []string) bool {
  return row[0] == row[1] &&
         row[1] == row[2] && 
         row[2] == "X"
}


// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Spalte drei mal X enthält.
func checkColumns(board [][]string) bool  {
  return false // TODO
}

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Diagonale drei mal X enthält.
func checkDiagonals(board [][]string) bool  {
  return false // TODO
}



