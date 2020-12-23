package assignment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/twinj/uuid"

	"github.com/gorilla/mux"
)

type Handler struct {
	db *DB
}

// Instantiate REST handler with DB client
func NewHandler(db *DB) *Handler {
	return &Handler{db: db}
}

// Upload the video file and insert it to our database with an auto-generated record ID
func (h *Handler) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var filename string
	var title string
	var err error

	vars := mux.Vars(r)

	// Extract filename from REST URI
	if filename, err = extractFileName(vars["filename"]); err != nil {
		updateHeader(w, err)
		return
	}

	// Extract title from REST URI
	if title, err = extractTitle(vars["title"]); err != nil {
		updateHeader(w, err)
		return
	}

	start := time.Now()
	filebytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error with path %s: %v", filename, err)
		updateHeader(w, ErrMissingFile)
		return
	}

	// Save video to DB and report the database ID for the record
	id := h.db.Create(filebytes, title, time.Now(), time.Since(start))
	updateHeader(w, nil)
	fmt.Fprintf(w, "Record is created and saved into %s", id)
}

// Return list of records sorted by their record creation time
func (h *Handler) GetAllSortedWithCreatedAt(w http.ResponseWriter, r *http.Request) {
	records := h.db.GetRecordsSortedByCreatedAt()
	payload, err := json.Marshal(records)
	if err != nil {
		log.Println("fail to encode data to the JSON format, error:", err)
		updateHeader(w, err)
		return
	}

	updateHeader(w, err)
	w.Write(payload)
}

// Delete record by its database ID
func (h *Handler) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, err := extractUUID(vars["id"])

	if err != nil {
		updateHeader(w, ErrInvalidID)
		return
	}

	h.db.Delete(*uuid)
	updateHeader(w, nil)
}

// Get record by its database ID and post the result in the response as a JSON
func (h *Handler) GetRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, err := extractUUID(vars["id"])

	if err != nil {
		updateHeader(w, ErrInvalidID)
		return
	}

	data, err := h.db.Get(*uuid)
	if err != nil {
		updateHeader(w, err)
		return
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Println("fail to encode data to the JSON format, error:", err)
		updateHeader(w, err)
		return
	}

	updateHeader(w, nil)
	w.Write(payload)
}

func extractUUID(vid string) (*uuid.UUID, error) {
	id := strings.TrimSpace(vid)
	if len(id) == 0 {
		log.Println("empty record ID")
		return nil, ErrInvalidID
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Println("invalid uuid provided:", id)
		return nil, ErrInvalidID
	}

	return uuid, nil
}

func extractFileName(filename string) (string, error) {
	file := strings.TrimSpace(filename)
	if len(file) == 0 {
		return "", ErrInvalidFileName
	}
	return file, nil
}

func extractTitle(title string) (string, error) {
	name := strings.TrimSpace(title)
	if len(name) == 0 {
		return "", ErrInvalidTitle
	}
	return title, nil
}

// Map the internal server error to the proper http.statusCode
func updateHeader(w http.ResponseWriter, err error) {
	var statusCode int

	switch err {
	case nil:
		statusCode = http.StatusOK
	case ErrInvalidFileName, ErrorMissingRecord, ErrInvalidTitle, ErrMissingFile:
		statusCode = http.StatusBadRequest
	case ErrorMissingRecord, ErrInvalidID:
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
