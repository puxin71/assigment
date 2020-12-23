package assignment

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrInvalidFileName = errors.New("Invalid video filename")
	ErrMissingFile     = errors.New("Missing video file")
	ErrInvalidTitle    = errors.New("Invalid title")
	ErrInvalidID       = errors.New("Invalid record ID")
	ErrorMissingRecord = errors.New("Record ID is not present")
)

// Server REST paths
func NewRouter(handler *Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/upload/{filename}/{title}", handler.CreateRecord).Methods("POST")
	router.HandleFunc("/records", handler.GetAllSortedWithCreatedAt).Methods("GET")
	router.HandleFunc("/records/{id}", handler.GetRecord).Methods("GET")
	router.HandleFunc("/delete/{id}", handler.DeleteRecord).Methods("POST")
	router.HandleFunc("/", Home)
	return router
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home!")
}
