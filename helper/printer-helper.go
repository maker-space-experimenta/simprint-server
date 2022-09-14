package helper

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/maker-space-experimenta/printer-kiosk/models"
	"github.com/maker-space-experimenta/printer-kiosk/util"
)

func GetPrinterProfiles(ctx context.Context, apiUrl string, apiKey string) (models.Printer, error) {
	log.Printf("running GetPrinterProfiles for %v \n", apiUrl)

	urlPrinterprofile, err := url.JoinPath(apiUrl, "printerprofiles")
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("GET", urlPrinterprofile, nil)
	if err != nil {
		log.Fatalln(err)
		return models.Printer{}, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("X-Api-Key", apiKey)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return models.Printer{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return models.Printer{}, err
	}

	var jsonRes map[string]interface{}
	_ = json.Unmarshal(body, &jsonRes)

	p := jsonRes["profiles"].(map[string]interface{})
	d := p["_default"].(map[string]interface{})

	return models.Printer{
		Name:  d["name"].(string),
		Model: d["model"].(string),
	}, nil
}

func GetPrinter(ctx context.Context, apiUrl string, apiKey string) (models.Printer, error) {
	log.Printf("running GetPrinter for %v \n", apiUrl)

	urlPrinterprofile, err := url.JoinPath(apiUrl, "printer")
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("GET", urlPrinterprofile, nil)
	req = req.WithContext(ctx)

	req.Header.Set("X-Api-Key", apiKey)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return models.Printer{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return models.Printer{}, err
	}

	var jsonRes map[string]interface{}
	_ = json.Unmarshal(body, &jsonRes)

	s := jsonRes["state"].(map[string]interface{})

	return models.Printer{
		State: s["text"].(string),
	}, nil
}

func GetPrinterMetaData(config util.Config) ([]models.Printer, error) {

	var printers []models.Printer

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	urls := strings.Split(config.Printers, ";")
	for _, apiUrl := range urls {

		// apiKey := "F84833BA940E49648F606EA561E97783"
		apiKey := strings.Split(apiUrl, ",")[1]
		apiUrl = strings.Split(apiUrl, ",")[0]

		pp, err := GetPrinterProfiles(ctx, apiUrl, apiKey)

		if err != nil {
			log.Fatalln(err)
			return []models.Printer{}, err
		}

		p, err := GetPrinter(ctx, apiUrl, apiKey)

		if err != nil {
			log.Fatalln(err)
			return []models.Printer{}, err
		}

		printer := models.Printer{
			Url:   apiUrl,
			Name:  pp.Name,
			Model: pp.Model,
			State: p.State,
		}

		printers = append(printers, printer)
	}

	return printers, nil
}
