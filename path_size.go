// Package path_size provides functions for getting the size
// of files and directories on disk, with the possibility of recursive traversal
// and formatting the size in a human-readable form.
package path_size

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FormatSize returns the size in human-readable format if isHumanReadable=true.
func FormatSize(bytes int64, isHumanReadable bool) string {
	const base = 1024

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	size := bytes
	suffix := units[0]

	var mantissa int64
	for i := 0; isHumanReadable && size >= base && i < len(units)-1; i++ {
		mantissa = size % base
		size /= base
		suffix = units[i+1]
	}

	tenths := (mantissa * 10) / base
	if tenths == 0 {
		return fmt.Sprintf("%d%s", size, suffix)
	}

	return fmt.Sprintf("%d.%d%s", size, tenths, suffix)
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

// GetPathSize returns the size of a file or directory in the "<size>\t<name>" format.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fileIinfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	name := fileIinfo.Name()

	var totalSize int64

	if fileIinfo.IsDir() {
		totalSize, err = GetDirSize(path, recursive, all)
		if err != nil {
			return "", err
		}
	} else {
		totalSize = fileIinfo.Size()
	}

	return fmt.Sprintf("%s\t%s", FormatSize(totalSize, human), name), nil
}
