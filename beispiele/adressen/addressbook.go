package main

// Struct für die Einträge eines Adressbuchs.
type Entry struct {
	name      string
	surname   string
	street    string
	town      string
	favourite bool
}

func (e *Entry) makeFavourite() {
	e.favourite = true
}

func (e *Entry) unFavourite() {
	e.favourite = false
}

// Datentyp für das Adressbuch selbst.
type AddressBook []Entry

// (Alias für den Datentyp []Entry)

/* Alternative, kompliziertere Variante
type AddressBook struct {
  entries []Entry
}
*/
