package main

import "fmt"

func main() {
  spielfeld := makeBoard(3, "*")

  printBoard(spielfeld)

  fmt.Println(checkRows(spielfeld, "X")) // Soll false ausgeben

  spielfeld2 := makeBoard(3, "X")
  fmt.Println(checkRows(spielfeld2, "X")) // Soll true ausgeben

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
// es in einer Zeile drei mal 'char' enthält.
func checkRows(board [][]string, char string) bool {
  // Zeile 0 pruefen
  for _, row := range(board) {
    if checkList(row, char) {
      return true // Early Out
    }
  }
  return false
}

// Prüft eine einzelne Zeile, ob sie drei mal 'char' enthält.
func checkList(row []string, char string) bool {
  for _,element := range(row) {
    if element != char {
      return false
    }
  }
  return true
}

// Liefert die i-te Spalte des Spielfelds als Liste.
func getColumn(board [][]string, i int) []string {
  var result []string
  for _,row := range(board) {
    result = append(result,row[i])
  }
  // TODO
  return result
}

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Spalte drei mal 'char' enthält.
// Funktioniert erstmal nur für quadratische Spielfelder.
func checkColumns(board [][]string, char string) bool  {
  for i,_ := range(board) {
    if checkList(getColumn(board, i), char) {
      return true // Early Out
    }
  }
  return false
}

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Diagonale drei mal 'char' enthält.
func checkDiagonals(board [][]string, char string) bool  {
  return false // TODO
}

