package main

import "fmt"

func main() {
  l1 := []int{20,3,17,35,42,27,12}
  
  fmt.Println(qsort(l1))
}

func qsort(list []int) []int {

  if len(list) <= 1 {
    return list
  }

  var result []int

  // Alles, was kleiner ist, als das erste Element, kommt nach links.
  var links []int
  // Alles, was größer ist, als das erste Element, kommt nach rechts.
  var rechts[]int

  for _, v := range(list[1:]) {
    if v < list[0] {
      links = append(links, v)
    } else {
      rechts = append(rechts, v)
    }
  }

  // Die beiden Teillisten rekursiv sortieren
  links = qsort(links)
  rechts = qsort(rechts)

  // Die beiden Teillisten mit dem früheren ersten Element in der Mitte zusammenfassen.

  for _,v := range(links) {
    result = append(result, v)
  }
  result = append(result, list[0])
  for _,v := range(rechts) {
    result = append(result, v)
  }

  return result
}