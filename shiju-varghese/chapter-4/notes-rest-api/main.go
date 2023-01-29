package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedOn	time.Time	`json:"cretedon"`
}

// strore for the Notes collection
var notesStore = make (map[string]Note)

// var to generae key for collection
var id int = 0

// HTTP post api- /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	notesStore[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

// GET
func GetNoteHandler(w http.ResponseWriter, r * http.Request){
	var notes []Note
	for _,v := range notesStore {
		notes = append(notes, v)
	}

	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}


// delete->   /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	k := vars["id"]
	// remove from strore
	if _, ok := notesStore[k]; ok {
		// delete existing item
		delete(notesStore, k)
	}else{
		log.Printf("Could not find key of Note %s to delete ",k)
	}

	w.WriteHeader(http.StatusNoContent)
}

// put-> /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r * http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpd Note
	// Decode the incoing Note json
	err = json.NewDecoder(r.Body).Decode(&noteToUpd)
	if err != nil {
		panic(err)
	}

	if note, ok := notesStore[k]; ok {
		noteToUpd.CreatedOn = note.CreatedOn
		// delete existing item and add the update item
		delete(notesStore,k)
		notesStore[k] = noteToUpd
	}else {
		log.Panicf("Could not find any key %s to delte ", k)
	}
	w.WriteHeader(http.StatusNoContent)

}

// emtry point 
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}",PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr : ":8080",
		Handler : r,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}