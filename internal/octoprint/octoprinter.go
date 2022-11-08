package octoprint

import (
	"context"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

// A representation of a octoprint server
type Octoprinter struct {
	ctx      context.Context
	hostname string
	apiKey   string
	logger   *logging.Logger
}

func NewOctoprinter(ctx context.Context, hostname string, apiKey string) (*Octoprinter, error) {
	return &Octoprinter{
		ctx:      ctx,
		hostname: hostname,
		apiKey:   apiKey,
		logger:   logging.NewLogger(),
	}, nil
}
