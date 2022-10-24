package printers

import (
	"context"
	"time"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
	"github.com/maker-space-experimenta/printer-kiosk/internal/octoprint"
)

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
		pp, err := op.GetPrinterProfiles()

		printer := PrinterModel{
			Hostname: printerConfig.Host,
			Name:     pp.Profiles["_default"].Name,
			Model:    pp.Profiles["_default"].Model,
			State:    p.State.Text,
		}

		printers = append(printers, printer)

	}

	return &printers, nil
}
