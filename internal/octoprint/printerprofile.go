package octoprint

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type PrinterProfile struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Model    string `json:"model"`
	Default  bool   `json:"default"`
	Current  bool   `json:"current"`
	Resource string `json:"resource"` //url
	Volume   struct {
		FormFactor string `json:"formFactor"` // "rectangular"
		Origin     string `json:"origin"`     // "lowerleft"
		Width      int    `json:"width"`
		Depth      int    `json:"depth"`
		Height     int    `json:"height"`
	} `json:"volume"`
	HeatedBed     bool `json:"heatedBed"`
	HeatedChamber bool `json:"heatedChamber"`
	Axes          struct {
		X struct {
			Speed    int  `json:"speed"`
			Inverted bool `json:"inverted"`
		} `json:"x"`
		Y struct {
			Speed    int  `json:"speed"`
			Inverted bool `json:"inverted"`
		} `json:"y"`
		Z struct {
			Speed    int  `json:"speed"`
			Inverted bool `json:"inverted"`
		} `json:"z"`
		E struct {
			Speed    int  `json:"speed"`
			Inverted bool `json:"inverted"`
		} `json:"e"`
	} `json:"axes"`
	Extruder struct {
		Count   int `json:"count"`
		Offsets []struct {
			X float32 `json:"x"`
			Y float32 `json:"y"`
		} `json:"offsets"`
	} `json:"extruder"`
}

type PrinterProfileList struct {
	Profiles map[string]PrinterProfile `json:"profiles"`
}

func (m *Octoprinter) GetPrinterProfiles() (*PrinterProfileList, error) {
	m.logger.Infof("running GetPrinterProfiles for %v ", m.hostname)

	apiUrl := fmt.Sprintf("%v://%v/api", "http", m.hostname)
	urlPrinterprofile, err := url.JoinPath(apiUrl, "printerprofiles")
	if err != nil {
		m.logger.Errorf("%v", err)
	}

	req, err := http.NewRequest("GET", urlPrinterprofile, nil)
	if err != nil {
		m.logger.Errorf("%v", err)
		return nil, err
	}

	req = req.WithContext(m.ctx)
	req.Header.Set("X-Api-Key", m.apiKey)

	// client := http.DefaultClient
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		m.logger.Errorf("FATAL: client: error making http request: %s", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		m.logger.Errorf("%v", err)
		return nil, err
	}

	var result PrinterProfileList
	_ = json.Unmarshal(body, &result)

	if len(result.Profiles) > 0 {
		return &result, nil
	}

	return nil, errors.New("PrinterProfile Response invalid")
}
