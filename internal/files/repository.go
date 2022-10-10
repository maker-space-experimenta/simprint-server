package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

type IFileRepository interface {
}

type FileRepository struct {
	config         configuration.Config
	Files          []PrusaSlicerGcodeMetaData
	DeleteDuration int
	logger         *logging.Logger
}

var fileRepoLock = &sync.Mutex{}
var fileRepoInstance *FileRepository

func NewFileRepository(config configuration.Config) *FileRepository {
	if fileRepoInstance == nil {
		fileRepoLock.Lock()

		if fileRepoInstance == nil {
			fileRepoInstance = &FileRepository{
				config:         config,
				DeleteDuration: config.Files.DeleteDuration,
				logger:         logging.NewLogger(),
			}
		}
	}

	return fileRepoInstance
}

func (m *FileRepository) UpdateFiles() {

	var filesList []PrusaSlicerGcodeMetaData
	filesDeleted := 0

	dirName := path.Join(m.config.Files.TempDir, "gcode")

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		m.logger.Errorf("%v", err)
	}

	for _, file := range files {
		m.logger.Debugf("found file %v", file.Name())

		if file.ModTime().Before((time.Now().Add(time.Duration((m.DeleteDuration * -1)) * time.Minute))) {
			filesDeleted++
			m.logger.Debugf("File %v is outdated and will be deleted", file.Name())
			e := os.Remove(fmt.Sprintf("%v/%v", dirName, file.Name()))
			if e != nil {
				m.logger.Errorf("%v", e)
			}
		} else {

			fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", dirName, file.Name()))
			if err != nil {
				m.logger.Errorf("%v", err)
			}

			gcode := string(fileBytes)

			f, err := GCodeToMap(gcode)
			if err != nil {
				m.logger.Errorf("%v", err)
			}

			metadata := PrusaSlicerGcodeMetaData{}

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
				imgFirstLine := strings.Index(img, "")
				img = img[imgFirstLine:]
				img = strings.Replace(img, "", "", -1)
				img = strings.Replace(img, ";", "", -1)
				metadata.Image = strings.TrimSpace(img)
			}

			filesList = append(filesList, metadata)
		}

	}

	m.Files = filesList

	m.logger.Infof("updated files - %v files found and updated - %v files deleted", len(filesList), filesDeleted)
}
