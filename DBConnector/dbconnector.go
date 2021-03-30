package DBConnector

import (
	"github.com/EmanuelFeij/NotesApp/note"
)

type DBConnector interface {
	GetAll() ([]note.Note, error)
	Get(name string) (note.Note, error)
	Put(note.Note) error
	Delete(name string) error
}
