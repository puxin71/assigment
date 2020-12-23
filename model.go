package assignment

import (
	"time"

	"github.com/twinj/uuid"
)

type Recording struct {
	// Unique ID of the new video
	ID uuid.UUID
	// Video name
	Title string `json:"title"`
	// Timestamp for the video's creation
	CreatedAt time.Time
	// Timestamp when the user uploaded the video
	UploadedAt time.Time `json:"uploadedAt"`
	// Duration of the video in seconds
	Duration time.Duration `json:"duration"`
	// Videio content
	Data []byte
}

type ByCreatedAt []Recording

func (a ByCreatedAt) len() int      { return len(a) }
func (a ByCreatedAt) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Sort the recording list by the descending order of the CreatedAt time
func (a ByCreatedAt) Less(i, j int) bool { return a[i].CreatedAt.After(a[j].CreatedAt) }
