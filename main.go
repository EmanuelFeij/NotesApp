package main

import (
	"fmt"
	"log"

	"github.com/EmanuelFeij/NotesApp/dbConnector"
	"github.com/EmanuelFeij/NotesApp/note"
)

func main() {
	n1 := note.Note{Title: "Emanuel", Body: "SEX"}

	n2 := note.Note{Title: "Eman", Body: "SEXting"}
	n3 := note.Note{Title: "Emanue", Body: "SEXy"}
	n4 := note.Note{Title: "Emanuee", Body: "SEXo"}

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
