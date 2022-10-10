package configuration

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	LogLevel string `xaml:"loglevel"`

	Files struct {
		TempDir        string `yaml:"temp_dir"`
		DeleteDuration int    `yaml:"delete_durations_minutes"`
	} `yaml:"files"`

	Tasks struct {
		Duration int `yaml:"duration"`
	} `yaml:"tasks"`

	Printers map[string]struct {
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
