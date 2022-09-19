package repositories

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/maker-space-experimenta/printer-kiosk/internal/helper"
	"github.com/maker-space-experimenta/printer-kiosk/internal/models"
	"github.com/maker-space-experimenta/printer-kiosk/internal/util"
)

type FileRepository struct {
	config util.Config
	Files  []models.PrusaSlicerGcodeMetaData
}

var lock = &sync.Mutex{}
var fileRepoInstance *FileRepository

func NewFileRepository(config util.Config) *FileRepository {
	if fileRepoInstance == nil {
		lock.Lock()

		if fileRepoInstance == nil {
			fileRepoInstance = &FileRepository{
				config: config,
			}
		}
	}

	return fileRepoInstance
}

func (m *FileRepository) UpdateFiles() {

	var filesList []models.PrusaSlicerGcodeMetaData
	filesDeleted := 0

	dirName := m.config.TempFileDir

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// log.Printf("found file %v", file.Name())

		if file.ModTime().Before((time.Now().Add(time.Duration((m.config.FileDeleteDurationMinutes * -1)) * time.Minute))) {
			filesDeleted++
			// log.Printf("File %v is outdated and will be deleted", file.Name())
			e := os.Remove(fmt.Sprintf("%v/%v", dirName, file.Name()))
			if e != nil {
				log.Fatal(e)
			}
		} else {

			fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", dirName, file.Name()))
			if err != nil {
				log.Fatal(err)
			}

			gcode := string(fileBytes)

			f, err := helper.GCodeToMap(gcode)
			if err != nil {
				log.Fatal(err)
			}

			metadata := models.PrusaSlicerGcodeMetaData{}

			metadata.FileName = file.Name()
			metadata.Options = f
			metadata.EstimatedPrintingTimeNormalMode = f["estimated_printing_time_normal_mode"]
			metadata.EstimatedPrintingTimeSilentMode = f["estimated_printing_time_silent_mode"]
			metadata.FilamentCost = f["filament_cost"]
			metadata.FilamentUsedCM3 = f["filament_used_cm3"]
			metadata.FilamentUsedGramm = f["filament_used_g"]
			metadata.FilamentUsedMM = f["filament_used_mm"]
			metadata.TotalFilamentCost = f["total_filament_cost"]
			metadata.TotalFilamentUsedGramm = f["total_filament_used_g"]

			imgBegin := strings.Index(gcode, "thumbnail begin")
			imgEnd := strings.Index(gcode, "thumbnail end")

			if imgBegin > -1 {
				img := gcode[imgBegin:imgEnd]
				imgFirstLine := strings.Index(img, "\n")
				img = img[imgFirstLine:]
				img = strings.Replace(img, "\n", "", -1)
				img = strings.Replace(img, ";", "", -1)
				metadata.Image = strings.TrimSpace(img)
			}

			filesList = append(filesList, metadata)
		}

	}

	m.Files = filesList

	log.Printf("updated files - %v files found and updated - %v files deleted", len(filesList), filesDeleted)
}
