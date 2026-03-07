package path_size

import (
	"fmt"
	"log"
	"os"
)

func FormatSize(size int64, isHumanReadable bool) string {
	if !isHumanReadable {
		return fmt.Sprintf("%dB", size)
	}

	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}

	suffixes := [...]string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	exp := 0
	divisor := int64(unit)

	for size/divisor >= unit && exp < len(suffixes)-1 {
		divisor *= unit
		exp += 1
	}
	
	value := float64(size) / float64(divisor)

	return fmt.Sprintf("%.1f%s", value, suffixes[exp])
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	if recursive || all {
		return "WIP", nil
	}

	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	name := info.Name()
	var totalSize int64

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

	return fmt.Sprintf("%s\t%s", FormatSize(totalSize, human), name), nil
}
