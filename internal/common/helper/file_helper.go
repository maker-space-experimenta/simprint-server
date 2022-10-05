package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
 */
func SaveFileFromForm(r *http.Request, field string, dir string, filename string) (string, error) {
	file, _, err := r.FormFile(field)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	filepath := fmt.Sprintf("%v/%v", dir, filename)
	log.Printf(filepath)

	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return filepath, nil
}
