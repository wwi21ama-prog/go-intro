package main

import "fmt"
import "bufio" //buffered input/output
import "os"    // operating system

func main() {
  
  // Erste Möglichkeit etwas einzugeben
  scanner := bufio.NewScanner(os.Stdin) //standard input
  fmt.Println("Bitte etwas eingeben: ")
  scanner.Scan()
  fmt.Println("Sie haben", scanner.Text(), "eingegeben.")

// Zweite Möglichkeit etwas einzugeben
  fmt.Println("Bitte noch etwas eingeben: ")
  var input string
  fmt.Scanln(&input)
  fmt.Println("Sie haben", input, "eingegeben.")
}
