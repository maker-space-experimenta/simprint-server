package helper

import (
	"encoding/json"
	"log"
	"testing"
)

func TestReadValueFromKey(t *testing.T) {

	body := []byte("{ \"key\": \"value\" }")

	var jsonRes map[string]interface{}
	_ = json.Unmarshal(body, &jsonRes)

	val, err := ReadValueFromKey("key", jsonRes)

	log.Print(val)

	if err != nil {
		t.Error(err)
	}

	if val != "value" {
		t.Errorf("wrong value")
		return
	}
}

func TestReadValueFromKeyHirarchic(t *testing.T) {

	body := []byte("{ \"key\": {\"key2\": {\"key3\": \"value\"} } }")

	var jsonRes map[string]interface{}
	_ = json.Unmarshal(body, &jsonRes)

	val, err := ReadValueFromKey("key.key2", jsonRes)

	log.Print(val)

	if err != nil {
		t.Error(err)
	}

	if val != "value" {
		t.Errorf("wrong value")
		return
	}
}
