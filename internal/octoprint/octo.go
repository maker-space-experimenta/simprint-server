package octoprint

import (
	"context"
)

type Octoprinter struct {
	ctx      context.Context
	hostname string
	apiKey   string
}

func NewOctoprinter(ctx context.Context, hostname string, apiKey string) (*Octoprinter, error) {
	return &Octoprinter{
		ctx:      ctx,
		hostname: hostname,
		apiKey:   apiKey,
	}, nil
}
