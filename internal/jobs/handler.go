package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
	"github.com/maker-space-experimenta/printer-kiosk/internal/octoprint"
)

type PostPrintModel struct {
	File    string `json:"file"`
	Printer string `json:"printer"`
}

type PrintHandler struct {
	config configuration.Config
	logger *logging.Logger
}

func NewPrintHandler(config configuration.Config) *PrintHandler {
	return &PrintHandler{
		config: config,
		logger: logging.NewLogger(),
	}
}

func (m *PrintHandler) PostPrint(w http.ResponseWriter, r *http.Request) {

	m.logger.Infof("executing PostPrint Endpoint")

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var bodyData PostPrintModel
	_ = json.Unmarshal(bodyBytes, &bodyData)

	m.logger.Infof(bodyData.File)
	m.logger.Infof(bodyData.Printer)

	m.logger.Infof("sending print job")

	dirName := m.config.Files.TempDir
	filePath := fmt.Sprintf("%v/gcode/%v", dirName, bodyData.File)
	printer := m.config.Printers[bodyData.Printer]

	m.logger.Debugf("filepath is %v", filePath)
	m.logger.Debugf("printer is %v", printer)

	file, _ := os.Open(filePath)
	defer file.Close()

	op, err := octoprint.NewOctoprinter(r.Context(), printer.Host, printer.Key)
	op.PrintFile(file, r.Context())
	// op.SendFile(file, r.Context())

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{ \"result\": \"ok\" }"))
}
