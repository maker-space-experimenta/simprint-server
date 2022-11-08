package configuration

import (
	"testing"
)

func TestShouldCreateConfigService(t *testing.T) {
	cfgService := NewConfigService()

	if cfgService == nil {
		t.Errorf("Return Value of NewConfigService() is nil - should be ConfigService")
	}
}
