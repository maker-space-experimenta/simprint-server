package configuration

type Config struct {
	Server struct {
		Port string `yaml:"port" json:"port"`
	} `yaml:"server" json:"server"`

	Logging struct {
		Level string `yaml:"level" json:"level"`
	} `yaml:"logging" json:"logging"`

	Files struct {
		TempDir        string `yaml:"temp_dir" json:"temp_dir"`
		DeleteDuration int    `yaml:"delete_durations_minutes" json:"delete_durations_minutes"`
	} `yaml:"files" json:"files"`

	Tasks struct {
		Duration int `yaml:"duration"`
	} `yaml:"tasks"`

	Printers map[string]struct {
		Host  string `yaml:"host"`
		Key   string `yaml:"key"`
		Image string `default:"" yaml:"image,omitempty"`
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

	Slicer struct {
		Path string `yaml:"path"`
	} `yaml:"slicer"`
}
