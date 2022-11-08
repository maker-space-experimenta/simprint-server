package configuration

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type ConfigService struct {
	configPath string
	config     *Config
}

var configServiceLock = &sync.Mutex{}
var configServiceInstance *ConfigService

func NewConfigService() *ConfigService {
	if configServiceInstance == nil {
		configServiceLock.Lock()
		defer configServiceLock.Unlock()
		if configServiceInstance == nil {
			fmt.Println("Creating single instance now.")
			configServiceInstance = &ConfigService{
				configPath: "./config.yml",
			}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return configServiceInstance
}

func (m *ConfigService) LoadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&m.config); err != nil {
		return err
	}

	return nil
}

func (m *ConfigService) GetConfig() (*Config, error) {
	return m.config, nil
}
