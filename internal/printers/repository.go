package printers

import (
	"log"
	"sync"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
)

type PrinterRepository struct {
	config   configuration.Config
	Printers *[]PrinterModel
}

var printerRepoLock = &sync.Mutex{}
var printerRepoInstance *PrinterRepository

func NewPrinterRepository(config configuration.Config) *PrinterRepository {
	if printerRepoInstance == nil {
		printerRepoLock.Lock()

		if printerRepoInstance == nil {
			printerRepoInstance = &PrinterRepository{
				config: config,
			}
		}
	}

	return printerRepoInstance
}

func (m *PrinterRepository) UpdatePrinters() {
	printers, err := GetPrintersMetaData(m.config)

	if err != nil {
		log.Fatalln(err)
		return
	}

	m.Printers = printers
	log.Printf("updated printers - %v printers found and updated", len(*printers))
}
