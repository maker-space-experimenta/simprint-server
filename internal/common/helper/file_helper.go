package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

/*
 */
func SaveFileFromForm(r *http.Request, field string, dir string, filename string) (string, error) {
	logger := logging.NewLogger()

	file, handler, err := r.FormFile(field)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if filename == "" {
		filename = handler.Filename
	}

	filepath := fmt.Sprintf("%v/%v", dir, filename)
	logger.Infof(filepath)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

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
