package datamgmt

import (
	"crypto/sha256"
	"encoding/hex"
)

func calculateChecksum(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
