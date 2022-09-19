package models

type Printer struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Model    string `json:"model"`
	State    string `json:"state"`
}
