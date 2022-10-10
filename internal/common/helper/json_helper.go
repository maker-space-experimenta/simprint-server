package helper

import (
	"errors"
	"strings"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func ReadValueFromKey(key string, json map[string]interface{}) (string, error) {
	logger := logging.NewLogger()

	nodes := strings.Split(key, ".")
	logger.Infof("node[o]: %v", nodes[0])

	if len(nodes) > 1 {
		if json[nodes[0]] != nil {

			if len(nodes[1:]) > 1 {
				k := strings.Join(nodes[1:], ".")
				d := json[nodes[0]].(map[string]interface{})
				return ReadValueFromKey(k, d)
			}

			if json[nodes[1]] != nil {
				return json[nodes[1]].(string), nil
			}

			return "", errors.New("key not found")
		}
	}

	if json[nodes[0]] != nil {
		return json[nodes[0]].(string), nil
	}

	return "", errors.New("key not found")
}
