package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/octoprint"
)

type PostPrintModel struct {
	File    string `json:"file"`
	Printer string `json:"printer"`
}

type PrintHandler struct {
	config configuration.Config
}

func NewPrintHandler(config configuration.Config) *PrintHandler {
	return &PrintHandler{
		config: config,
	}
}

func (m *PrintHandler) PostPrint(w http.ResponseWriter, r *http.Request) {

	log.Println("executing PostPrint Endpoint")

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var bodyData PostPrintModel
	_ = json.Unmarshal(bodyBytes, &bodyData)

	log.Println(bodyData.File)
	log.Println(bodyData.Printer)

	log.Println("sending print job")

	dirName := m.config.Files.TempDir
	filePath := fmt.Sprintf("%v/%v", dirName, bodyData.File)
	printer := m.config.Printers[bodyData.Printer]

	file, _ := os.Open(filePath)
	defer file.Close()

	op, err := octoprint.NewOctoprinter(r.Context(), printer.Host, printer.Key)
	op.PostFiles(file)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{ \"result\": \"ok\" }"))
}

// func SliceFile() {

// model := ""
// config_path := ""
// scale := ""
// output := ""

// args := []string{
// 	"-g", model,
// 	"--load", config_path,
// 	"--scale-to-fit", scale,
// 	"--output", output,
// }

// cmd := exec.Command("prusa-slicer", args)
// }
