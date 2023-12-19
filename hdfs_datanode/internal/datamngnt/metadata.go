package datamgmt

import (
	"time"
)

type BlockMetadata struct {
	ID        string    `json:"id"`
	Size      int64     `json:"size"`
	Checksum  string    `json:"checksum"` // You can use hash functions like SHA-256
	CreatedAt time.Time `json:"created_at"`
}
