package main

import "fmt"

/* Spielfeld definieren
   Passen Sie die Spielfeld-Funktionen an, so dass:
   - ein 8x8 (oder besser nxn)-Spielfeld korrekt angezeigt wird.

   Schreiben Sie Funktionen, die die Diagonalen des Spielfelds liefern und prüfen können, ob sie schon belegt sind.

   Schreiben Sie eine Funktion, die die i-te Spalte des Spielfelds prüft, ob sie leer ist.

   Schreiben Sie die Funktion 'allowed()', die prüft, ob an einer bestimmten Stelle eine Dame gesetzt werden darf.
*/

func main() {
	board := makeBoard(8, " ")
	printBoard(board)
	fmt.Println()
	solve(board, 0)
	printBoard(board)
}

/*
  Löst das n-Damen-Problem rekursiv.
  Parameter:
  - board: Das Spielfeld.
  - line: Die Nummer der Zeile, in die die nächste
          Dame platziert werden soll

  Rückgabe: true, falls das Spiel gelöst ist, sonst false.

  Sollte initial mit line == 0 aufgerufen werden.
*/
func solve(board [][]string, line int) bool {
	// Wenn das Spiel für eine Zeilennummer gelöst werden soll
	// die es nicht gibt, dann ist es schon gelöst.
	if line >= len(board) {
		return true
	}

	// Eine Dame in Zeile 'line' platzieren.
	// D.h. in einer Schleife nach freien Plätzen suchen.
	for column := 0; column < len(board[line]); column++ {
		// Suche nach freiem Platz in dieser Zeile.
		if allowed(board, line, column) {
			// Dame an Stelle board[line][column] setzen.
			board[line][column] = "D"
			result := solve(board, line+1)
			if result {
				return true
			}
			// Dame wieder entfernen
			board[line][column] = " "
		}
	}
	return false
}

func allowed(board [][]string, line, column int) bool {
	return columnAllowed(board, column) &&
		diag1Allowed(board, line, column) &&
		diag2Allowed(board, line, column)
}

func makeBoard(size int, initChar string) [][]string {
	// Definieren einer 2D-Slice
	var board [][]string

	for i := 0; i < size; i++ {
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
	for i, row := range board {
		printRow(row)
		if i < len(board)-1 {
			fmt.Println("---+---+---+---+---+---+---+---")
		}
	}
}

// Eine Zeile des Spielfelds mit Trennzeichen ausgeben.
func printRow(row []string) {
	for j := 0; j < len(row); j++ {
		fmt.Print(" " + row[j] + " ")
		if j < len(row)-1 {
			fmt.Print("|")
		}
	}
	fmt.Println()
}

// Hilfsfunktionen, können genutzt werden, um zu prüfen, ob eine bestimmte Spalte nur Leerzeichen enthält.
// Das kann genutzt werden, um zu prüfen, ob eine Spalte noch frei ist.

// Prüft eine einzelne Liste, ob sie nur 'char' enthält.
func checkList(row []string, char string) bool {
	for _, element := range row {
		if element != char {
			return false
		}
	}
	return true
}

// Liefert die i-te Spalte des Spielfelds als Liste.
func getColumn(board [][]string, i int) []string {
	var result []string
	for _, row := range board {
		result = append(result, row[i])
	}
	return result
}

func columnAllowed(board [][]string, column int) bool {
	return checkList(getColumn(board, column), " ")
}

func diag1Allowed(board [][]string, line, column int) bool {
	return checkList(getDiag1(board, line, column), " ")
}

func getDiag1(board [][]string, line, column int) []string {
	var result []string
	for l, c := line, column; l < len(board) && c < len(board); l, c = l+1, c+1 {
		result = append(result, board[l][c])
	}
	for l, c := line, column; l >= 0 && c >= 0; l, c = l-1, c-1 {
		result = append(result, board[l][c])
	}
	return result
}

// Was hier fehlt: getDiag2
