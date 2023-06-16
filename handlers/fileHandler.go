package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)


func GetAllPaths(dirPath string) ([]string, error) {
	var paths []string

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		paths = append(paths, filePath)

		if file.IsDir() {
			subPaths, err := GetAllPaths(filePath)
			if err != nil {
				log.Println("Error reading subdirectory:", err.Error())
				continue
			}
			paths = append(paths, subPaths...)
		}
	}

	return paths, nil
}

func PathsHandler(w http.ResponseWriter, r *http.Request) {
	dirPath := "./data/" // Replace with your desired directory path

	paths, err := GetAllPaths(dirPath)
	if err != nil {
		log.Println("Error retrieving paths:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the paths to JSON	
	jsonData, err := json.Marshal(paths)
	if err != nil {
		log.Println("Error marshaling JSON:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON data to the response
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("Error writing JSON response:", err.Error())
	}
}