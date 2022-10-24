package slicer

import (
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/helper"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

type SlicerHandler struct {
	config configuration.Config
	logger *logging.Logger
}

func NewSlicerHandler(config configuration.Config) *SlicerHandler {
	return &SlicerHandler{
		config: config,
	}
}

func (m *SlicerHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	m.logger.Infof("running GetPrinters")

	// jsonResp, err := json.Marshal(m.printerRepo.Printers)
	// if err != nil {
	// 	logger.Errorff("Error happened in JSON marshal. Err %s", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write(nil)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonResp)
}

func (m *SlicerHandler) PostSlicejob(w http.ResponseWriter, r *http.Request) {

	m.logger.Infof("call PostSlicejob")

	stlPath, err := helper.SaveFileFromForm(r, "file", path.Join(m.config.Files.TempDir, "stl"), "model.stl")
	if err != nil {
		m.logger.Infof("FATAL: Could not decode and save stl file from request. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	config_path := "slicer-configs/config_pla_03mm_draft.ini"
	scale := "30,30,30"
	output := path.Join(m.config.Files.TempDir, "gcode", "foo.gcode")

	args := []string{
		"-g", stlPath,
		"--load", config_path,
		"--scale-to-fit", scale,
		"--output", output,
	}

	cmd := exec.Command("prusa-slicer", args...)
	cmd.Run()

	err = os.Remove(stlPath)
	if err != nil {
		m.logger.Infof("FATAL: Could not decode and save stl file from request. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

}
