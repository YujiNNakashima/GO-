package filebackup

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"path/filepath"
)

func HashIt(rootDir string) ([][]string, error) {

	currentDirFiles, err := globFiles(filepath.Join(rootDir, "*"))
	if err != nil {
		return nil, err
	}

	subDirFiles, err := globFiles(filepath.Join(rootDir, "**", "*"))
	if err != nil {
		return nil, err
	}

	allFiles := append(currentDirFiles, subDirFiles...)

	var fileHashPairs [][]string
	for _, filePath := range allFiles {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return nil, err
		}
		if fileInfo.Mode().IsRegular() {
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			hashValue, err := hashPath(fileContent)
			if err != nil {
				fileHashPairs = append(fileHashPairs, []string{filePath, ""})
				continue
			}
			fileHashPairs = append(fileHashPairs, []string{filePath, hashValue})
		}
	}

	return fileHashPairs, nil
}

func hashPath(content []byte) (string, error) {
	hasher := sha1.New()
	_, err := hasher.Write(content)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func globFiles(pattern string) ([]string, error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}
