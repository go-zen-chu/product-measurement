package main

import (
	"encoding/json"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	port             = flag.String("port", "8080", "Server port")
	jsonDataStoreDir = flag.String("data", "./data", "Json data store root directory path")
)

const (
	dataPath = "/data/"
)

func main() {
	flag.Parse()

	http.HandleFunc(dataPath, func(w http.ResponseWriter, r *http.Request) {
		urlRelPath := r.URL.Path[len(dataPath):]

		if !strings.HasSuffix(urlRelPath, ".json") {
			http.NotFound(w, r)
			return
		}

		fullPath := filepath.Join(*jsonDataStoreDir, urlRelPath)

		if _, err := os.Stat(fullPath); err == nil {
			http.ServeFile(w, r, fullPath)
			return
		}

		basePath := strings.TrimSuffix(urlRelPath, ".json")

		currentPath := basePath
		for currentPath != "" {
			dirPath := filepath.Join(*jsonDataStoreDir, currentPath)

			dirInfo, err := os.Stat(dirPath)
			if err == nil && dirInfo.IsDir() {
				handleMergedJsonData(w, r, dirPath)
				return
			}

			currentPath = filepath.Dir(currentPath)
			if currentPath == "." || currentPath == "/" {
				break
			}
		}

		// check if the root directory itself is a valid JSON data store
		rootInfo, err := os.Stat(*jsonDataStoreDir)
		if err == nil && rootInfo.IsDir() {
			handleMergedJsonData(w, r, *jsonDataStoreDir)
			return
		}

		http.NotFound(w, r)
	})

	slog.Info("Starting server", "url", "http://localhost:"+*port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}

func handleMergedJsonData(w http.ResponseWriter, r *http.Request, dirPath string) {
	slog.Info("Merging JSON files", "directory", dirPath)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		slog.Error("Cannot read directory", "error", err, "path", dirPath)
		http.Error(w, "Cannot read directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var mergedData []interface{}
	foundFiles := false

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			filePath := filepath.Join(dirPath, file.Name())
			slog.Info("Processing JSON file", "path", filePath)

			data, err := os.ReadFile(filePath)
			if err != nil {
				slog.Error("File read error", "path", filePath, "error", err)
				continue
			}

			var jsonArray []interface{}
			if err := json.Unmarshal(data, &jsonArray); err == nil {
				mergedData = append(mergedData, jsonArray...)
				foundFiles = true
				slog.Info("Added items as array", "count", len(jsonArray), "file", filePath)
			} else {
				var jsonObject interface{}
				if err := json.Unmarshal(data, &jsonObject); err == nil {
					mergedData = append(mergedData, jsonObject)
					foundFiles = true
					slog.Info("Added as object", "file", filePath)
				} else {
					slog.Error("JSON parsing error", "path", filePath, "error", err)
				}
			}
		}
	}

	if !foundFiles {
		slog.Warn("No valid JSON files found", "directory", dirPath)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mergedData)
	if err != nil {
		slog.Error("Failed to encode JSON response", "error", err)
		http.Error(w, "Failed to encode JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
