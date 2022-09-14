package models

type Printer struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Model string `json:"model"`
	State string `json:"state"`
}
