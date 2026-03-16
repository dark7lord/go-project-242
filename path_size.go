// Package code provides functions for getting the size
// of files and directories on disk, with the possibility of recursive traversal
// and formatting the size in a human-readable form.
package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const base = 1024

// FormatSize returns the size in human-readable format if isHumanReadable=true.
func FormatSize(size int64, isHumanReadable bool) string {
	if !isHumanReadable || size < base {
		return fmt.Sprintf("%dB", size)
	}

	resultSize := float64(size)
	index := 0

	for resultSize >= base {
		resultSize /= base
		index++
	}

	return fmt.Sprintf("%.1f%s", resultSize, units[index])
}

// GetDirSize returns the total size of the directory,
// with the option of recursive traversal and hidden files.
func GetDirSize(path string, recursive, all bool) (totalSize int64, err error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, file := range files {
		fileName := file.Name()
		if !all && strings.HasPrefix(fileName, ".") {
			continue
		}

		fullPath := filepath.Join(path, fileName)

		if file.IsDir() {
			if !recursive {
				continue
			}

			var dirSize int64

			dirSize, err = GetDirSize(fullPath, recursive, all)
			if err != nil {
				return totalSize, err
			}

			totalSize += dirSize
		} else {
			fileInfo, err := file.Info()
			if err != nil {
				return totalSize, err
			}

			totalSize += fileInfo.Size()
		}
	}

	return totalSize, err
}

// GetPathSize returns the size of a file or directory.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fileIinfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	var totalSize int64

	if fileIinfo.IsDir() {
		totalSize, err = GetDirSize(path, recursive, all)
		if err != nil {
			return "", err
		}
	} else {
		totalSize = fileIinfo.Size()
	}

	return FormatSize(totalSize, human), nil
}
