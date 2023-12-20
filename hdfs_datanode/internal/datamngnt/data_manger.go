package datamgmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	dataDir     = "data"
	metadataDir = "metadata"
	blockPrefix = "block_"
)

// DataManager handles storage and retrieval of data blocks
type DataManager struct {
	dataPath     string
	metadataPath string
}

// NewDataManager creates a new instance of DataManager
func NewDataManager(basePath string) *DataManager {
	return &DataManager{
		dataPath:     filepath.Join(basePath, dataDir),
		metadataPath: filepath.Join(basePath, metadataDir),
	}
}

func (dm *DataManager) StoreBlock(blockID string, data []byte) error {
	// Ensure the data directory exists
	if err := os.MkdirAll(dm.dataPath, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %v", err)
	}

	blockPath := filepath.Join(dm.dataPath, blockPrefix+blockID)

	// Write the data block in binary format
	if err := ioutil.WriteFile(blockPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write data block: %v", err)
	}

	// Create and store metadata
	metadata := &BlockMetadata{
		ID:        blockID,
		Size:      int64(len(data)),
		Checksum:  calculateChecksum(data),
		CreatedAt: time.Now(),
	}
	if err := dm.SaveMetadata(metadata); err != nil {
		return fmt.Errorf("failed to save metadata: %v", err)
	}

	return nil
}

// RetrieveBlock retrieves the data for the given block ID
func (dm *DataManager) RetrieveBlock(blockID string) ([]byte, error) {
	blockPath := filepath.Join(dm.dataPath, blockPrefix+blockID)
	data, err := ioutil.ReadFile(blockPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read data block: %v", err)
	}
	return data, nil
}

// File: internal/datamgmt/data_manager.go (continued)

// SaveMetadata saves the metadata for a given data block
func (dm *DataManager) SaveMetadata(metadata *BlockMetadata) error {
	metadataPath := filepath.Join(dm.metadataPath, metadata.ID+".metadata")
	data, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(metadataPath, data, 0644)
}

// LoadMetadata loads the metadata for a given block ID
func (dm *DataManager) LoadMetadata(blockID string) (*BlockMetadata, error) {
	metadataPath := filepath.Join(dm.metadataPath, blockID+".metadata")
	data, err := ioutil.ReadFile(metadataPath)
	if err != nil {
		return nil, err
	}
	var metadata BlockMetadata
	err = json.Unmarshal(data, &metadata)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}
