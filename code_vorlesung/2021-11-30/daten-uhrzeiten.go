package main

import (
	"fmt"
)

type Date struct {
	day   int
	month int
	year  int
}

func main() {

	d1 := Date{25, 12, 2021}
	d2 := Date{13, 1, 2022}

	printDate(d1)
	printDate(d2)
	printDate(nextDay(d2))

	fmt.Println(daysInMonth(Date{31, 5, 2022}))
	fmt.Println(daysInMonth(Date{20, 2, 2024}))

	fmt.Println(isLeapYear(2024) == true)
	fmt.Println(isLeapYear(2023) == false)
	fmt.Println(isLeapYear(1900) == false)
	fmt.Println(isLeapYear(1600) == true)

	fmt.Println(nextDay(Date{31, 8, 2022}))
	fmt.Println(nextDay(Date{30, 9, 2022}))
	fmt.Println(nextDay(Date{28, 2, 2024}))
	fmt.Println(nextDay(Date{31, 12, 2024}))

	fmt.Println(weekday(Date{30, 11, 2021}))

}

// Erwartet ein Datum und liefert das darauf folgenden Datum.
func nextDay(d Date) Date {

	// Wird true oder false, je nachdem ob wir den letzten Tag haben.
	isLastDay := d.day == daysInMonth(d)

	if isLastDay {
		if d.month == 12 {
			return Date{1, 1, d.year + 1}
		}
		return Date{1, d.month + 1, d.year}
	}
	return Date{d.day + 1, d.month, d.year}
}

// Gibt ein Datum auf die Konsole aus.
func printDate(d Date) {
	fmt.Printf("%v. %v. %v\n", d.day, d.month, d.year)
}

// Erwartet eine Monatsnummer und liefert die Anzahl der Tage dieses Monats.
func daysInMonth(d Date) int {
	if d.month == 2 && isLeapYear(d.year) {
		return 29
	}

	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	return days[d.month-1]
}

// Erwartet eine Jahreszahl und liefert true, falls dieses Jahr ein Schaltjahr ist.
func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

// Erwartet zwei Daten und gibt die Anzahl der Tage zwischen diesen beiden Daten zurück.
func days(d1, d2 Date) int {
	if d1 == d2 {
		return 0
	}
	return days(nextDay(d1), d2) + 1

	/* Als Schleife:
	   result := 0
	   for d1 != d2 {
	     d1 = nextDay(d1)
	     result += 1
	   }
	   return result
	*/

}

// Liefert den Wochentag eines Datums.
func weekday(d Date) int {
	monday := Date{1, 1, 1}
	dist := days(monday, d)
	return dist%7 + 1
}
