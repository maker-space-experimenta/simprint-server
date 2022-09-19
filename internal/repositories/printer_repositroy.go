package repositories

import (
	"log"

	"github.com/maker-space-experimenta/printer-kiosk/internal/helper"
	"github.com/maker-space-experimenta/printer-kiosk/internal/models"
	"github.com/maker-space-experimenta/printer-kiosk/internal/util"
)

type PrinterRepository struct {
	config   util.Config
	Printers *[]models.Printer
}

func NewPrinterRepository(config util.Config) *PrinterRepository {
	return &PrinterRepository{
		config: config,
	}
}

func (m *PrinterRepository) UpdatePrinters() {
	printers, err := helper.GetPrintersMetaData(m.config)

	if err != nil {
		log.Fatalln(err)
		return
	}

	m.Printers = printers
	log.Printf("updated printers - %v printers found and updated", len(*printers))
}
