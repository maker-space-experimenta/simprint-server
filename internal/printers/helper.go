package printers

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
	"github.com/maker-space-experimenta/printer-kiosk/internal/octoprint"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func readFileToBase64(path string) string {

	logger := logging.NewLogger()

	if path == "" || path == "default" {
		logger.Debugf("no path for image, reading default image")
		path = "./images/_default.png"
	}

	logger.Debugf("reading image from path %v", path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Errorf("File not found %v", path)
		return ""
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)

	return base64Encoding
}

func GetPrintersMetaData(config configuration.Config) (*[]PrinterModel, error) {
	logger := logging.NewLogger()

	var printers []PrinterModel

	for _, printerConfig := range config.Printers {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		op, err := octoprint.NewOctoprinter(ctx, printerConfig.Host, printerConfig.Key)
		if err != nil {
			logger.Errorf("error getting printer: %s", err)
		}

		p, err := op.GetPrinter()
		if err != nil {
			return &printers, err
		}

		pp, err := op.GetPrinterProfiles()
		if err != nil {
			return &printers, err
		}

		image := readFileToBase64(printerConfig.Image)

		printer := PrinterModel{
			Hostname: printerConfig.Host,
			Name:     pp.Profiles["_default"].Name,
			Model:    pp.Profiles["_default"].Model,
			State:    p.State.Text,
			Image:    image,
		}

		printers = append(printers, printer)

	}

	return &printers, nil
}
