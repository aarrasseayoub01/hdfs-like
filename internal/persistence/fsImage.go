package persistence

import (
	"encoding/gob"
	"os"

	"github.com/aarrasseayoub01/hdfs-mini/internal/fs"
)

func saveFsImage(dir *fs.Directory, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(dir)
	return err
}

func loadFsImage(path string) (*fs.Directory, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var dir fs.Directory
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&dir)
	return &dir, err
}
