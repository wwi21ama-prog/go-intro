# Aufgaben zu Funktionen, Schleifen und If-Then-Else

Dieses Dokument enthält eine Reihe von Aufgaben, die man nur mittels Variablen, Schleifen und If-Then-Else lösen kann.

Für die ersten paar Aufgaben ist in der Datei `01_Kontrollfluss.go` ein Gerüst als Hilfestellung vorgegeben.
Außerdem sind dort auch Tests formuliert, also Funktionen, die überprüfen, ob die zu schreibenden Funktionen korrekt sind.

Für die hinteren Aufgaben aus dieser Datei sollten Sie sich jeweils zuerst überlegen, wie das Funktionsgerüst aussehen sollte.
Guter Stil ist es auch, analog zu den vorhandenen Tests neue für die weiteren Funktionen zu schreiben.

## Vielfache bestimmen
Schreiben Sie eine Funktion `sumMultiples()`, die zwei ganzzahlige Parameter `x`und `n` erwartet.
Die Funktion soll die Summe aller Vielfachen von `x` liefern, die kleiner als `n` sind.

## Primzahltest
Schreiben Sie eine Funktion `isPrime()`, die einen ganzzahligen Parameter `n` erwartet.
Die Funktion soll genau dann `true` liefern, wenn `n` eine Primzahl ist.

## Größten Primfaktor bestimmen.
Schreiben Sie eine Funktion `largestPrimeFactor()`, die einen ganzzahligen Parameter `n` erwartet.
Die Funktion soll die größte Primzahl liefern, durch die `n` teilbar ist.

## Kleinstes gemeinsames Vielfaches bestimmen.
Schreiben Sie eine Funktion `lcm()`, die zwei ganzzahlige Parameter `m`und `n` erwartet.
Die Funktion soll die kleinste Zahl liefern, die durch `m` und `n` teilbar ist.

## Wurzelberechnung 1
Schreiben Sie eine Funktion `sqrtInt()`, die eine ganze Zahl `n` erwartet.
Die Funktion soll die Wurzel aus `n` berechnen, falls diese existiert und ganzzahlig ist.
Falls die Wurzel nicht existiert, soll für positive `n` die größte Zahl geliefert werden, deren Quadrat kleiner ist als `n`. Für negative `n` muss die Funktion nicht definiert sein.

## Wurzelberechnung 2
Schreiben Sie eine Funktion `sqrt()`, die eine Fließkommazahl `x` erwartet.
Die Funktion soll die Wurzel von `n` mittels der Newton-Iteration näherungsweise berechnen.

*Hinweis:* Verwenden Sie die Iterationsformel $z_{n+1} = z_n - \frac{z_n^2 - x}{2z_n}$
mit einem Startwert $z_0 = 1$.
Sie müssen so lange immer neue Werte f+r $z_n$ berechnen, bis das Ergebnis "genau genug" ist.

## Summe der Zahlen von 1 bis n
Schreiben Sie eine Funktion `summe()`, die eine ganze Zahl `n` als Parameter erwartet und die die Summer alle Zahlen von 1 bis `n` liefert.

## Ziffern auflisten
Schreiben Sie eine Funktion `printZiffern()`, die eine ganze Zahl `n` erwartet und die
alle Dezimalziffern von `n` auf die Konsole schreibt.

## Quersumme berechnen
Schreiben Sie eine Funktion `digitSum()`, die eine ganze Zahl `n` erwartet und die die Quersumme von `n` liefert.
