package main

import "fmt"

func main() {
  //foo(10)

  for i:=1; i<10; i++ {
    fmt.Println(i, fib(i))
    // 1 1
    // 2 1
    // 3 2
    // 4 3
    // 5 5
    // 6 8
    // 7 13
    // 8 21
    // 9 34
  }
}

func foo(k int) {
  if k == 0 {
    return
  }
  foo(k-1)
  fmt.Println(k)
}

func fib(n int) int {
  // fib(1) = 1
  // fib(2) = 1
  if n == 1 || n == 2 { // Oder: if n<=2
    return 1
  }

  // fib(n) = fib(n-1) + fib(n-2)
  return fib(n-1) + fib(n-2)
}



