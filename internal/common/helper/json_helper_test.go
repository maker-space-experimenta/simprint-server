package helper

import (
	"encoding/json"
	"testing"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func TestReadValueFromKey(t *testing.T) {
	logger := logging.NewLogger()

	body := []byte("{ \"key\": \"value\" }")

	var jsonRes map[string]interface{}
	_ = json.Unmarshal(body, &jsonRes)

	val, err := ReadValueFromKey("key", jsonRes)

	logger.Infof(val)

	if err != nil {
		t.Error(err)
	}

	if val != "value" {
		t.Errorf("wrong value")
		return
	}
}

func TestReadValueFromKeyHirarchic(t *testing.T) {
	logger := logging.NewLogger()

	body := []byte("{ \"key\": {\"key2\": {\"key3\": \"value\"} } }")

	var jsonRes map[string]interface{}
	_ = json.Unmarshal(body, &jsonRes)

	val, err := ReadValueFromKey("key.key2", jsonRes)

	logger.Infof(val)

	if err != nil {
		t.Error(err)
	}

	if val != "value" {
		t.Errorf("wrong value")
		return
	}
}
