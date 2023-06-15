package api

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type File struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	ModTime int64  `json:"modTime"`
	Size    int64  `json:"size"`
}

func WalkDirectory(pathToWalk string) ([]byte, error) {
	var files []*File

	err := filepath.Walk(pathToWalk, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file := &File{
				Name:    info.Name(),
				ModTime: info.ModTime().Unix(),
				Path:    path,
				Size:    info.Size(),
			}
			files = append(files, file)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Marshal files into JSON format
	jsonData, err := json.Marshal(files)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
