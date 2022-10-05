package octoprint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/*
response from /api/printer
*/
type PrinterState struct {
	Temperature struct {
		Tool0 struct {
			Actual float32 `json:"actual"`
			Target float32 `json:"target"`
			Offset int     `json:"offset"`
		} `json:"tool0"`

		Tool1 struct {
			Actual float32 `json:"actual"`
			Target float32 `json:"target"`
			Offset int     `json:"offset"`
		} `json:"tool1"`

		Bed struct {
			Actual float32 `json:"actual"`
			Target float32 `json:"target"`
			Offset int     `json:"offset"`
		} `json:"bed"`

		History []struct {
			Time  int `json:"time"`
			Tool0 struct {
				Actual float32 `json:"actual"`
				Target float32 `json:"target"`
				Offset int     `json:"offset"`
			} `json:"tool0"`

			Tool1 struct {
				Actual float32 `json:"actual"`
				Target float32 `json:"target"`
				Offset int     `json:"offset"`
			} `json:"tool1"`

			Bed struct {
				Actual float32 `json:"actual"`
				Target float32 `json:"target"`
				Offset int     `json:"offset"`
			} `json:"bed"`
		}
	} `json:"temperature"`

	Sd struct {
		Ready bool `json:"ready"`
	} `json:"sd"`

	State struct {
		Text  string `json:"text"`
		Flags struct {
			Operational   bool `json:"operational"`
			Paused        bool `json:"paused"`
			Printing      bool `json:"printing"`
			Cancelling    bool `json:"cancelling"`
			Pausing       bool `json:"pausing"`
			SdReady       bool `json:"sdReady"`
			Error         bool `json:"error"`
			Ready         bool `json:"ready"`
			ClosedOrError bool `json:"closedOrError"`
		} `json:"flags"`
	} `json:"state"`
}

func (m *Octoprinter) GetPrinter() (*PrinterState, error) {
	log.Printf("running GetPrinter for %v \n", m.hostname)

	apiUrl := fmt.Sprintf("%v://%v/api", "http", m.hostname)

	urlPrinterprofile, err := url.JoinPath(apiUrl, "printer")
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("GET", urlPrinterprofile, nil)
	req = req.WithContext(m.ctx)

	req.Header.Set("X-Api-Key", m.apiKey)

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

	var response PrinterState
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error in GetPrinter: %v", err)
	}

	return &response, nil
}
