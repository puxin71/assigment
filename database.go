package assignment

import (
	"log"
	"sort"
	"sync"
	"time"

	"github.com/twinj/uuid"
)

type DB struct {
	// Synchronized in-memory database to store and retrieve the vedio records
	smap *sync.Map
}

// Initialize database
func NewDB() *DB {
	return &DB{smap: new(sync.Map)}
}

// Always insert a video file
func (d *DB) Create(data []byte, title string, uploadedTime time.Time, duration time.Duration) uuid.UUID {
	id := uuid.NewV4()

	record := Recording{
		ID:         id,
		Title:      title,
		CreatedAt:  time.Now(),
		UploadedAt: uploadedTime,
		Duration:   duration,
		Data:       data,
	}
	d.smap.LoadOrStore(id, record)
	return id
}

// Delete the video record by its unique ID
func (d *DB) Delete(id uuid.UUID) {
	d.smap.Delete(id)
}

// Get the video by its unique ID
func (d *DB) Get(id uuid.UUID) (*Recording, error) {
	data, ok := d.smap.Load(id)
	if !ok {
		log.Printf("error to get the video record, %s", id)
		return nil, ErrorMissingRecord
	}

	return data.(*Recording), nil
}

// Return all saved recording in the descening order of CreatedAt time
func (d *DB) GetRecordsSortedByCreatedAt() []Recording {
	recordings := make([]Recording, 0)
	d.smap.Range(func(key, value interface{}) bool {
		recordings = append(recordings, value.(Recording))
		return true
	})

	sort.Slice(recordings, func(i, j int) bool {
		return recordings[i].CreatedAt.After(recordings[j].CreatedAt)
	})

	return recordings
}
