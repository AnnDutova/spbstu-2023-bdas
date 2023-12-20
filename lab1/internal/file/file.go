package file

import (
	"os"
	"strings"
)

func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteFile(content, filePath, newFileName string) error {
	pathFolders := strings.Split(filePath, "/")
	path := strings.Join(append(pathFolders[:len(pathFolders)-1], newFileName), "/")
	return os.WriteFile(path, []byte(content), 0666)
}
