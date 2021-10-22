package main

import "fmt"
import "math"

func main() {

  // Einen String definieren.
  s1 := "Hallo"
  fmt.Println(s1)

  // s1 und seinen Typ mit Printf ausgeben.
  fmt.Printf("%v: %T\n", s1, s1)

  // Einen zweiten String definieren.
  // Gleich einen Wert festlegen.
  var s2 string = "Strings"

  // Einen dritten String deklarieren...
  var s3 string

  // ... einen Wert im Nachhinein festlegen:
  s3 = s1 + " " + s2
  fmt.Println(s3)

  // Ein einzelnes Zeichen ("Character") definieren
  c1 := 97
  fmt.Println(c1)

  fmt.Printf("%c: %T\n", c1, c1)
  
  // Das 'a' an s1 anhängen und ausgeben.
  s1 = s1 + string(c1)
  fmt.Println(s1)

  // Versuch, einen String aus Ziffern in eine Zahl zu konvertieren:
  //i1 := int("123")

  fmt.Printf("%c\n", s1[3])
  fmt.Println(string(s1[4]))

  for i := 0; i < len(s1); i++ {
    fmt.Println(string(s1[i]))
  }

  fmt.Println(stringToInt("123"))

  fmt.Println('1')

  //stringToIntFull("-1234")
}

// Funktion, die einen String erwartet, der nur Ziffern enthält, und die die entsprechende Zahl als int liefert.
func stringToInt(s string) int {
  result := 0

  for i := 0; i < len(s); i++ {
    //result = result + ziffer * 10^(len(s) - i - 1)
    ziffer := int(s[i] - '0')
    result = result + ziffer * myPow(10, len(s) - i - 1)
  }

  return result
}

func myPow(base int, exponent int) int {
  return int(math.Pow(float64(base), float64(exponent)))
}












func stringToIntFull(s string) int {
	isNegative := false
	number := 0

	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		isNegative = true
	}

	lastElement := len(s) - 1

	for position := lastElement; position >= 0; position-- {
		currentChar := s[position]

		if currentChar < '0' || currentChar > '9' {
			panic("Cannot parse string with non numeric characters")
		}

		power := lastElement - position
		digit := int(currentChar - '0')

		calculatedPower := int(math.Pow(10, float64(power)))

		numberBefore := number
		number += digit * calculatedPower

		fmt.Printf("number: %v \t->\t %v;\t digit: %v;\t power: %v\n", numberBefore, number, digit, power)
	}


	if isNegative {
		return number * -1
	} else {
		return number
	}
}
