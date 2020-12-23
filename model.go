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
