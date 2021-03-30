package dbConnector

import (
	"database/sql"
	"log"
	"os"

	"github.com/EmanuelFeij/NotesApp/note"
	_ "github.com/mattn/go-sqlite3"
)

type DBConnector interface {
	GetAll() ([]note.Note, error)
	Get(name string) ([]note.Note, error)
	Put(note.Note) error
	Delete(name string) error
	Close()
}

type DB struct {
	DbCon *sql.DB
}

func (db *DB) Close() {
	db.DbCon.Close()
}

func NewDB(path string) *DB {
	os.Remove(path)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal((err.Error()))
	}
	file.Close()

	sqlLiteDatabase, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal((err.Error()))
	}
	createTable(sqlLiteDatabase)

	return &DB{DbCon: sqlLiteDatabase}

}

func createTable(db *sql.DB) {
	createNotesTable := `CREATE TABLE notes (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT ,
		"title" TEXT,
		"body" TEXT
		);`
	log.Println("Creating notes table...")
	statement, err := db.Prepare(createNotesTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("notes table created successfully")
}

func (db *DB) GetAll() ([]note.Note, error) {
	log.Println("Getting everything from database")

	row, err := db.DbCon.Query("SELECT * FROM notes;")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	allNotes := make([]note.Note, 0)
	for row.Next() {
		var id int
		var title string
		var body string
		row.Scan(&id, &title, &body)
		allNotes = append(allNotes, note.Note{Title: title, Body: body})
	}
	log.Println("Sucess getting all notes")
	return allNotes, nil

}

func (db *DB) Get(name string) ([]note.Note, error) {
	log.Println("Getting note from database")
	getStatement := "SELECT * FROM notes WHERE title=(?)"
	rows, err := db.DbCon.Query(getStatement, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]note.Note, 0)
	for rows.Next() {
		var id int
		var title string
		var body string
		rows.Scan(&id, &title, &body)
		results = append(results, note.Note{Title: title, Body: body})
	}
	return results, nil
}

func (db *DB) Put(n note.Note) error {
	log.Println("Inserting the note ", n)
	insertNotesSql := `INSERT INTO notes(title,body) VALUES(?,?);`
	statement, err := db.DbCon.Prepare(insertNotesSql)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = statement.Exec(n.Title, n.Body)
	return err
}

func (db *DB) Delete(name string) error {
	log.Println("Deleting the note with title: ", name)
	insertStudenSql := `"DELETE FROM notes WHERE title=name;"`
	statement, err := db.DbCon.Prepare(insertStudenSql)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(name)

	return err
}
