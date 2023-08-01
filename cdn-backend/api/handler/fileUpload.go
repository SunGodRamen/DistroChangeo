package handler
//fileUpload_test.go

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFile handles a file upload
func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB file limit
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving the file"))
		return
	}
	defer file.Close()

	path, _ := os.Getwd()
	temp_repo_path := filepath.Join(path, "../../repo-temp/")
	f, err := os.OpenFile(temp_repo_path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error saving the file: " + err.Error()))
			return
	}	
	defer f.Close()

	io.Copy(f, file)
	
	w.Write([]byte("File uploaded successfully"))
}
