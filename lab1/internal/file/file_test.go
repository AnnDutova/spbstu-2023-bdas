package file_test

import (
	"io/ioutil"
	"testing"

	"github.com/AnnDutova/bdas/lab1/internal/file"
)

func TestReadFile(t *testing.T) {
	content := "Test file"
	testdataFile := "./testdata/tmp.txt"

	// Тестируем функцию ReadFile
	readContent, readErr := file.ReadFile(testdataFile)
	if readErr != nil {
		t.Fatalf("ReadFile returned an error: %v", readErr)
	}

	if readContent != content {
		t.Errorf("Expected content '%s', got '%s'", content, readContent)
	}

	// Тестируем случай ошибки
	nonExistentFile := "non_existent_file.txt"
	_, err := file.ReadFile(nonExistentFile)
	if err == nil {
		t.Errorf("Expected error for non-existent file, but got nil")
	}
}

func TestWriteFile(t *testing.T) {
	content := "Test WriteFile content"
	tmpdir := "./testdata/"
	fileName := "testfile.txt"

	// Тестируем WriteFile
	err := file.WriteFile(content, tmpdir, fileName)
	if err != nil {
		t.Fatal(err)
	}

	// Проверяем, что файл был создан и его содержимое соответствует ожидаемому
	result, err := ioutil.ReadFile(tmpdir + "/" + fileName)
	if err != nil {
		t.Fatal(err)
	}

	if string(result) != content {
		t.Errorf("Expected: %s, Got: %s", content, result)
	}
}
