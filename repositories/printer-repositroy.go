package repositories

import (
	"log"

	"github.com/maker-space-experimenta/printer-kiosk/helper"
	"github.com/maker-space-experimenta/printer-kiosk/models"
	"github.com/maker-space-experimenta/printer-kiosk/util"
)

type PrinterRepository struct {
	config   util.Config
	Printers []models.Printer
}

func NewPrinterRepository(config util.Config) *PrinterRepository {
	return &PrinterRepository{
		config: config,
	}
}

func (m *PrinterRepository) UpdatePrinters() {
	printers, err := helper.GetPrinterMetaData(m.config)

	if err != nil {
		log.Fatalln(err)
		return
	}

	m.Printers = printers
	log.Printf("updated printers - %v printers found and updated", len(printers))
}
