package main

import "fmt"

func main() {
  result := caesar("vorlesung")
  fmt.Println(result)
}




func caesar(wort string) string {
  result := ""

  for i := 0; i < len(wort); i++ {
    originalBuchstabe := wort[i]
    verschluesselterBuchstabe := 
      (originalBuchstabe + 3 - 'a') % 26 + 'a'
    result = result + string(verschluesselterBuchstabe)
  }

  return result
}