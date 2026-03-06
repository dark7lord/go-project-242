package sizegetters

import (
	"fmt"
	"log"
	"os"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	if recursive || human || all {
		return "WIP", nil
	}
	
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	
	name := info.Name()
	var totalSize int64 = 0
	
	if !info.IsDir() {
		totalSize = info.Size()
	} else {
		files, err := os.ReadDir(path)

		if err != nil {
			log.Fatal(err)
		}
		
		for _, file := range files {
			if !file.IsDir() {
				fileinfo, err := file.Info()
				if err != nil {
					return "", err
				}
				totalSize += fileinfo.Size()
			}
		}
	}
	
	return fmt.Sprintf("%dB\t%s", totalSize, name), nil
}