package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/maker-space-experimenta/printer-kiosk/internal/models"
	"github.com/maker-space-experimenta/printer-kiosk/internal/util"
)

type PrinterProfile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// Color   string `json:"color"`
	Model string `json:"model"`
	// Default bool   `json:"default"`
}

type PrinterProfileResponse struct {
	Profiles map[string]PrinterProfile
}

func GetPrinterProfiles(ctx context.Context, hostname string, apiKey string) (*models.Printer, error) {
	log.Printf("running GetPrinterProfiles for %v \n", hostname)

	apiUrl := fmt.Sprintf("%v://%v/api", "http", hostname)
	urlPrinterprofile, err := url.JoinPath(apiUrl, "printerprofiles")
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("GET", urlPrinterprofile, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("X-Api-Key", apiKey)

	// client := http.DefaultClient
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var jsonRes PrinterProfileResponse
	_ = json.Unmarshal(body, &jsonRes)

	if len(jsonRes.Profiles) > 0 {
		return &models.Printer{
			Name:  jsonRes.Profiles["_default"].Name,
			Model: jsonRes.Profiles["_default"].Model,
		}, nil
	}

	return nil, errors.New("PrinterProfile Response invalid")
}

func GetPrinter(ctx context.Context, hostname string, apiKey string) (*models.Printer, error) {
	log.Printf("running GetPrinter for %v \n", hostname)

	apiUrl := fmt.Sprintf("%v://%v/api", "http", hostname)

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
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var jsonRes map[string]interface{}
	_ = json.Unmarshal(body, &jsonRes)

	s := jsonRes["state"].(map[string]interface{})

	return &models.Printer{
		State: s["text"].(string),
	}, nil
}

func GetPrinterMetaData(hostname string, apiKey string) (*models.Printer, error) {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", hostname, 80), timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	pp, err := GetPrinterProfiles(ctx, hostname, apiKey)
	if err != nil {
		log.Printf("error getting printer profile: %s\n", err)
		return nil, err
	}

	p, err := GetPrinter(ctx, hostname, apiKey)
	if err != nil {
		return nil, err
	}

	printer := models.Printer{
		Hostname: hostname,
		Name:     pp.Name,
		Model:    pp.Model,
		State:    p.State,
	}

	return &printer, nil

}

func GetPrintersMetaData(config util.Config) (*[]models.Printer, error) {

	var printers []models.Printer

	urls := strings.Split(config.Printers, ";")
	for _, printer := range urls {

		apiKey := strings.Split(printer, ",")[1]
		hostname := strings.Split(printer, ",")[0]

		printer, err := GetPrinterMetaData(hostname, apiKey)
		if err != nil {
			log.Printf("error getting printer: %s\n", err)
		} else {
			printers = append(printers, *printer)
		}

	}

	return &printers, nil
}
