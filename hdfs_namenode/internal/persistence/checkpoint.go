package persistence

import (
	"time"
)

var (
	lastCheckpointTime   time.Time
	editLogSizeThreshold int
	checkpointInterval   time.Duration
)

func init() {
	lastCheckpointTime = time.Now()
	editLogSizeThreshold = 2           // for example, trigger after 1000 operations
	checkpointInterval = 1 * time.Hour // for example, trigger at least every hour
}

func shouldTriggerCheckpoint() bool {
	editLogMutex.Lock()
	defer editLogMutex.Unlock()

	editLogSize := len(editLog)
	timeSinceLastCheckpoint := time.Since(lastCheckpointTime)

	// Check if the edit log size exceeds the threshold or if enough time has elapsed
	if editLogSize >= editLogSizeThreshold || timeSinceLastCheckpoint >= checkpointInterval {
		lastCheckpointTime = time.Now() // Reset the checkpoint time
		return true
	}

	return false
}
