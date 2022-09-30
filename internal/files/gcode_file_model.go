package files

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
