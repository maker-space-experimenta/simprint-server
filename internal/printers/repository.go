package printers

import (
	"sync"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

type PrinterRepository struct {
	config   configuration.Config
	Printers *[]PrinterModel
	logger   *logging.Logger
}

var printerRepoLock = &sync.Mutex{}
var printerRepoInstance *PrinterRepository

func NewPrinterRepository(config configuration.Config) *PrinterRepository {
	if printerRepoInstance == nil {
		printerRepoLock.Lock()

		if printerRepoInstance == nil {
			printerRepoInstance = &PrinterRepository{
				config: config,
				logger: logging.NewLogger(),
			}
		}
	}

	return printerRepoInstance
}

func (m *PrinterRepository) UpdatePrinters() {
	printers, err := GetPrintersMetaData(m.config)

	if err != nil {
		m.logger.Errorf("%v", err)
		return
	}

	m.Printers = printers
	m.logger.Infof("updated printers - %v printers found and updated", len(*printers))
}
