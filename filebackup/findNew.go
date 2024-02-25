package filebackup

import (
	"path/filepath"
	"regexp"
)

type PathHashPair struct {
	Path string
	Hash string
}

func FindNew(rootDir string, pathHashPairs [][]string) (map[string]string, error) {
	hashToPath := make(map[string]string)

	for _, pair := range pathHashPairs {
		hashToPath[pair[0]] = pair[1]
	}

	pattern := filepath.Join(rootDir, "*.bck")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filename := filepath.Base(file)
		stripped := regexp.MustCompile(`\.bck$`).ReplaceAllString(filename, "")
		delete(hashToPath, stripped)
	}

	return hashToPath, nil
}
