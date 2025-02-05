package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ImageGroup represents a group of duplicate images
type ImageGroup []string

// findDuplicateImages finds all duplicate images in a directory
func findDuplicateImages(dir string) ([]ImageGroup, error) {
	// Map to store hash -> filenames
	hashMap := make(map[string][]string)

	// Valid image extensions
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
	}

	// Walk through directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file is an image
		ext := strings.ToLower(filepath.Ext(path))
		if !validExts[ext] {
			return nil
		}

		// Calculate hash
		hash, err := getFileHash(path)
		if err != nil {
			return fmt.Errorf("error hashing file %s: %v", path, err)
		}

		// Add to hash map
		hashMap[hash] = append(hashMap[hash], path)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %v", err)
	}

	// Convert hash map to array of duplicate groups
	var result []ImageGroup
	for _, files := range hashMap {
		if len(files) > 1 {
			result = append(result, files)
		}
	}

	return result, nil
}

// getFileHash calculates SHA-256 hash of a file
func getFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func main() {
	// Define command line flags
	dirPtr := flag.String("d", "", "Directory path to scan for duplicate images")
	flag.Parse()

	// Check if directory is provided
	if *dirPtr == "" {
		fmt.Println("Please provide a directory path using the -d flag")
		fmt.Println("Usage: ./app -d 'path/to/directory'")
		os.Exit(1)
	}

	// Check if directory exists
	if _, err := os.Stat(*dirPtr); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n", *dirPtr)
		os.Exit(1)
	}

	duplicates, err := findDuplicateImages(*dirPtr)
	if err != nil {
		fmt.Printf("Error finding duplicates: %v\n", err)
		os.Exit(1)
	}

	if len(duplicates) == 0 {
		fmt.Println("No duplicate images found")
		return
	}

	// Print results
	fmt.Printf("Found %d groups of duplicate images:\n", len(duplicates))
	for i, group := range duplicates {
		fmt.Printf("\nDuplicate group %d:\n", i+1)
		for _, file := range group {
			fmt.Printf("- %s\n", file)
		}
	}
}
