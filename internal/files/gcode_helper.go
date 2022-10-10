package files

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func GCodeToMap(gcode string) (map[string]string, error) {
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

func GCodeToImage(path string) {
	logger := logging.NewLogger()

	configService := configuration.NewConfigService()
	config, err := configService.GetConfig()
	if err != nil {
		logger.Errorf("cannot load config:", err)
	}

	dirName := fmt.Sprintf("%v/thumbnails", config.Files.TempDir)

	err = os.MkdirAll(dirName, os.ModePerm)
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

	if start == -1 {

	}

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

	f, err := os.OpenFile(fmt.Sprintf("%v/thumbnails/test.png", config.Files.TempDir), os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}

	png.Encode(f, im)
}
