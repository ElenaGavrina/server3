package utils

import (
	"errors"
	"net/http"
)

func CheckingFileFormat(w http.ResponseWriter, data []byte) error{
	filetype := http.DetectContentType(data)
	if filetype != "text/plain; charset=utf-8" {
		http.Error(w, "The provided file format is not allowed. Please upload a txt file", http.StatusBadRequest)
		return errors.New("this file can't be encrypted")
	}
	return nil
}