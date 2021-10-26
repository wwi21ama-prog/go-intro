package main

import "fmt"

func main() {
  
  // Variable für ganze Zahlen (32 oder 64 Bit)
  var i1 int = 42

  // Variable für ganze Zahlen mit 8 Bit (= 1 byte)
  var i2 byte = 15

  // Variable für 32-Bit-Fließkommazahlen
  var f1 float32 = 4.35

  // Variable für 64-Bit-Fließkommazahlen
  var f2 float64 = 5.82

  fmt.Println(i1, i2, f1, f2)


  res1 := f(4)
  fmt.Println(res1)

  // Dieser Aufruf funktioniert nicht, das ist ein Typfehler:
  // fmt.Println(f(i1))

  // Auch das ist ein Typfehler
  //fmt.Println(f(i2))

  // Das hier funktioniert:
  fmt.Println(f(f1))

  // Das hier geht wieder nicht.
  // fmt.Println(f(f2))

  //Abhilfe schaffen explizite Typumwandlungen:
  fmt.Println(f(float32(i1)))


  var res2 int = int(f(float32(i2)))
  fmt.Println(res2)

  fmt.Println(abs(-2))

  fmt.Println(f(float32(abs(-3))))

  var y uint = uint(-2)
  fmt.Println(y)
}


func f(x float32) float32 {
  return x * x
}

func abs(x int) uint {
  if x < 0 {
    return uint(-x)
  }
  return uint(x)
}