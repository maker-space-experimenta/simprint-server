package configuration

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Files struct {
		TempDir        string `yaml:"temp_dir"`
		DeleteDuration int    `yaml:"delete_durations_minutes"`
	} `yaml:"files"`

	Tasks struct {
		Duration int `yaml:"duration"`
	} `yaml:"tasks"`

	Printers []struct {
		Host string `yaml:"host"`
		Key  string `yaml:"key"`
	} `yaml:"printers"`

	Database struct {
		DBDriver      string `yaml:"db_driver"`
		DBSource      string `yaml:"db_source"`
		ServerAddress string `yaml:"server_address"`
	} `yaml:"database"`

	OAuth struct {
		ClientId     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
	} `yaml:"oauth"`
}

func LoadConfig(configPath string) (*Config, error) {

	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
