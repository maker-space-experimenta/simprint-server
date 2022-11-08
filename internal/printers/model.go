package printers

type PrinterModel struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Model    string `json:"model"`
	State    string `json:"state"`
	Image    string `json:"image"`
}
