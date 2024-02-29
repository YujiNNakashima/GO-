package filebackup

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
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

func BackItUp(srcDir string, dstDir string, timestamp int64) error {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	timestampStr := fmt.Sprintf("%010d", timestamp)
	parsedTimestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return err
	}

	existing, err := HashIt(srcDir)

	if err != nil {
		return err
	}

	needToCopy, err := FindNew(dstDir, existing)
	if err != nil {
		return err
	}

	copyFiles(dstDir, needToCopy)
	saveManifest(dstDir, parsedTimestamp, existing)
	return nil
}

func copyFiles(dst string, needToCopy map[string]string) error {
	for srcPath, hash := range needToCopy {
		dstPath := filepath.Join(dst, hash+".bck")

		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}

		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveManifest(dst string, timestamp int64, pathHash [][]string) error {
	sort.Slice(pathHash, func(i, j int) bool {
		return pathHash[i][0] < pathHash[j][0]
	})

	var content strings.Builder
	for _, pair := range pathHash {
		content.WriteString(fmt.Sprintf("%s,%s\n", pair[0], pair[1]))
	}
	manifestPath := filepath.Join(dst, fmt.Sprintf("%010d.csv", timestamp))

	fmt.Println("manifest")
	fmt.Println(manifestPath)
	file, err := os.Create(manifestPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(content.String()); err != nil {
		return err
	}

	return nil
}
