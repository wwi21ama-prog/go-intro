package main

import "fmt"

func main() {
  l1 := []string{"Auto","Baum","Haus"}
  l2 := []string{"car","tree","house"}

  result, valid := lookup("Auto",l1,l2)

  if valid {
    fmt.Println(result) // Soll "car" ausgeben
  } else {
    fmt.Println("Das war nichts!")
  }
}

// Liefert die Position von x in list, falls x darin vorkommt,
// liefert ansonsten -1
func find(list []string, x string) int {

  // Liste auf Leerheit prüfen.
  if len(list) == 0 {
    return -1
  }
  
  // 1. Mitte finden.
  // Genauer: Position des mittleren Elements bestimmen.
  mitte := len(list) / 2

  // 2. Prüfen, ob das mittlere Element das gesuchte ist.
  // Falls ja, dessen Position zurückliefern.
  if list[mitte] == x {
    return mitte
  }

  // 3. Linke und rechte Teilliste benennen.
  links := list[:mitte] // Vom Anfang bis vor dem mittleren Element.
  rechts := list[mitte+1:] // Von rechts von der Mitte bis zum Ende.

  // 4. Prüfen, ob das gesuchte Element kleiner oder größer als das mittlere ist.
  // Entsprechend links oder rechts weitersuchen.
  if x < list[mitte] {
    return find(links, x)
  } else {
    result := find(rechts, x)
    if result != -1 {
      return result  + mitte + 1
    }
  }
  return -1
}


// Aufgabe 1:
// Schreiben Sie eine Funktion lookup().
// Die Funktion soll 3 Parameter erwarten:
// 1. Einen Wert (String), der gesucht werden soll.
// 2. Eine Liste von Strings, in denen gesucht werden soll.
// 3. Eine Liste von Strings, aus denen das Ergebnis geliefert werden soll.
//
// Die Funktion soll den Suchstring in der ersten Liste suchen und das
// entsprechende Ergebnis aus der zweiten Liste liefern.
//
// Sie können annehmen, dass der Suchbereich sortiert ist.
func lookup(
  suchschluessel string,
  suchbereich, ergebnisbereich []string) (string, bool) {
    result := find(suchbereich, suchschluessel)
    if result >= 0 {
      return ergebnisbereich[result], true
    }
    return "", false
}