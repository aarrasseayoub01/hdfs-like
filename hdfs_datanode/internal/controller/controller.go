package controller

import (
	"encoding/json"
	"net/http"

	mng "github.com/aarrasseayoub01/namenode/datanode/internal/datamngnt" // Replace with your actual project path
	"github.com/gorilla/mux"
)

// BlockRequest represents the request structure for adding a block
type BlockRequest struct {
	BlockID string `json:"blockId"`
	Data    []byte `json:"data"`
}

// Controller holds the dependencies for a HTTP controller.
type Controller struct {
	DataManager *mng.DataManager
}

// NewController creates a new instance of Controller
func NewController(dataManager *mng.DataManager) *Controller {
	return &Controller{
		DataManager: dataManager,
	}
}

// addBlock handles POST requests to add a new block
func (c *Controller) AddBlock(w http.ResponseWriter, r *http.Request) {
	var blockReq BlockRequest
	err := json.NewDecoder(r.Body).Decode(&blockReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store the block using DataManager
	err = c.DataManager.StoreBlock(blockReq.BlockID, blockReq.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// getBlock handles GET requests to retrieve a block
func (c *Controller) GetBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockID := vars["blockId"]

	// Retrieve the block using DataManager
	data, err := c.DataManager.RetrieveBlock(blockID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the block data
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
