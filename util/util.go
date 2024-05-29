package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func List(path string) error {
	fmt.Println("Listing:", path)
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return err
		}
		// Corrected to use the appropriate methods from FileInfo
		fmt.Println(fmt.Printf("Mode\tSize\tModTime\tName"))
		fmt.Println(fmt.Printf("%v\t%12d\t%s\t%s", info.Mode(), info.Size(), info.ModTime().Format(time.RFC822), info.Name()))
	}
	return nil
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error:", err)
		return ""
	}
	return dir
}

func EmptyFolder(folderPath string) error {
	d, err := os.Open(folderPath)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		fullPath := filepath.Join(folderPath, name)
		err := os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
	}
	return nil
}
