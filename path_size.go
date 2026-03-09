package path_size

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FormatSize(bytes int64, isHumanReadable bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	const base = 1024
	size := bytes
	var mantissa int64
	i := 0
	for isHumanReadable && size >= base && i < len(units)-1 {
		mantissa = size % base
		size /= base
		i++
	}
	tenths := (mantissa * 10) / base
	if tenths == 0 {
		return fmt.Sprintf("%d%s", size, units[i])
	}
	return fmt.Sprintf("%d.%d%s", size, tenths, units[i])
}

func GetDirSize(path string, recursive bool, all bool) (totalSize int64, err error) {
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
