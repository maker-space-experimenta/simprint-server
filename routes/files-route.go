package routes

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/maker-space-experiemnta/printer-kiosk/util"
)

type PrusaSlicerGcodeMetaData struct {
	FileName                        string            `json:"filename"`
	Image                           string            `json:"image"`
	FilamentUsedMM                  string            `json:"filamentUsedMM"`
	FilamentUsedCM3                 string            `json:"filamentUsedCM3"`
	FilamentUsedGramm               string            `json:"filamentUsedGramm"`
	FilamentCost                    string            `json:"filamentCost"`
	TotalFilamentUsedGramm          string            `json:"totalFilamentUsedGramm"`
	TotalFilamentCost               string            `json:"totalFilamentCost"`
	EstimatedPrintingTimeNormalMode string            `json:"estimatedPrintingTimeNormalMode"`
	EstimatedPrintingTimeSilentMode string            `json:"estimatedPrintingTimeSilentMode"`
	Options                         map[string]string `json:"options"`
}

type FilesHandler struct {
	config util.Config
}

func NewFilesHandler(config util.Config) *FilesHandler {
	return &FilesHandler{
		config: config,
	}
}

func (m *FilesHandler) GCodeGetValueFromLine(line string) string {
	return strings.TrimSpace(strings.Split(line, "=")[1])
}
func (m *FilesHandler) GCodeSetIfExists(line string, pattern string, metadata *PrusaSlicerGcodeMetaData) {
	if strings.Index(line, "filament used [mm]") != -1 {
		metadata.FilamentUsedMM = m.GCodeGetValueFromLine(line)
	}
}

func (m *FilesHandler) GCodeToMap(gcode string) (map[string]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(gcode))

	l := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(line, ";") < 5 {
			if strings.Index(line, "=") > -1 {
				key := strings.Split(line, "=")[0]
				value := strings.Split(line, "=")[1]

				key = strings.Replace(key, ";", "", -1)
				key = strings.Replace(key, "[", "", -1)
				key = strings.Replace(key, "]", "", -1)
				key = strings.Replace(key, "(", "", -1)
				key = strings.Replace(key, ")", "", -1)
				key = strings.TrimSpace(key)
				key = strings.Replace(key, " ", "_", -1)

				l[key] = value
			}
		}
	}

	return l, nil
}

func (m *FilesHandler) GCodeToImage(path string) {
	dirName := fmt.Sprintf("%v/thumbnails", m.config.TempFileDir)

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	text := string(fileContent)

	start := strings.Index(text, "thumbnail begin")
	end := strings.Index(text, "thumbnail end")

	log.Printf("start %v, end %v", start, end)
	image := ""
	foo := false

	for i := start; i < end; i++ {
		c := string(text[i])
		if foo == false && c == ";" {
			foo = true
		}
		if foo && c != " " && c != ";" {
			image = image + c
		}
	}

	log.Print(image)

	unbased, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		panic("Cannot decode b64")
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		panic("Bad png")
	}

	f, err := os.OpenFile(fmt.Sprintf("%v/thumbnails/test.png", m.config.TempFileDir), os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}

	png.Encode(f, im)
}

/*	Endpoints
 */

func (m *FilesHandler) GetFiles(w http.ResponseWriter, r *http.Request) {
	log.Printf("running GetFiles")

	var filesList []PrusaSlicerGcodeMetaData

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
		log.Printf("found file %v", file.Name())

		if file.ModTime().Before((time.Now().Add(time.Duration((m.config.FileDeleteDurationMinutes * -1)) * time.Minute))) {
			log.Printf("File %v is outdated and will be deleted", file.Name())
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

			f, err := m.GCodeToMap(gcode)
			if err != nil {
				log.Fatal(err)
			}

			log.Print(f)

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

			img := gcode[imgBegin:imgEnd]
			imgFirstLine := strings.Index(img, "\n")
			img = img[imgFirstLine:]
			img = strings.Replace(img, "\n", "", -1)
			img = strings.Replace(img, ";", "", -1)
			metadata.Image = strings.TrimSpace(img)

			filesList = append(filesList, metadata)
		}

	}

	jsonResp, err := json.Marshal(filesList)
	if err != nil {
		log.Fatalf("Error happened in JSOn marshal. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (m *FilesHandler) PostFiles(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	// fileName := r.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dirName := m.config.TempFileDir

	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", dirName, handler.Filename)
	log.Printf(url)

	f, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	_, _ = io.WriteString(w, "File "+url+" Uploaded successfully")
	_, _ = io.Copy(f, file)

	// m.GCodeToImage(url)
}

func (m *FilesHandler) ReadFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("running ReadFile")

	vars := mux.Vars(r)
	log.Print(vars["filename"])

	filePath := fmt.Sprintf("%v/%v", m.config.TempFileDir, vars["filename"])

	err := os.MkdirAll(m.config.TempFileDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not create temp dir. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	fileBytes, err := ioutil.ReadFile(filePath)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/x-gcode")
	w.Write(fileBytes)
}
