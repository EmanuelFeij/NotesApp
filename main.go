package main

import (
	"fmt"
	"log"

	"github.com/EmanuelFeij/NotesApp/dbConnector"
	"github.com/EmanuelFeij/NotesApp/note"
)

func main() {
	n1 := note.Note{Title: "Emanuel", Body: "12 pounds"}

	n2 := note.Note{Title: "Eman", Body: "bola"}
	n3 := note.Note{Title: "Emanue", Body: "bola"}
	n4 := note.Note{Title: "Emanuee", Body: "ball"}

	var con dbConnector.DBConnector

	con = dbConnector.NewDB("db.db")
	defer con.Close()
	con.Put(n1)
	con.Put(n2)
	con.Put(n3)
	con.Put(n4)
	notes, err := con.GetAll()
	if err != nil {
		log.Println("here", err)
	}
	fmt.Println(notes)

	not, err := con.Get("Emanuel")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(not)
}
