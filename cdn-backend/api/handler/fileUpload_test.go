package handler_test
//handler_test.go

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/ra/cdn-backend/api/handler"
)

func TestUploadFile(t *testing.T) {
	path, _ := os.Getwd()
	test_fileDir := filepath.Join(path, "/")
	test_fileName := "testFile.txt"
	test_filePath := filepath.Join(test_fileDir, test_fileName)

	// Create test file
	test_file, _ := os.Create(test_filePath)
	test_file.Close()

	// Open file for reading
	test_file, err := os.Open(test_filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer test_file.Close()

	// Create multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("myFile", filepath.Base(test_file.Name()))

	_, _ = io.Copy(part, test_file)
	writer.Close()

	// Create http request
	req, _ := http.NewRequest("POST", "/upload", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// Create response recorder
	rr := httptest.NewRecorder()

	// Call UploadFile handler
	handler.UploadFile(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check response body
	expected := "File uploaded successfully"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Cleanup test file
	os.Remove(test_filePath)
}
