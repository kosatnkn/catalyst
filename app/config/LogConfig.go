package config

// LogConfig holds application log configurations.
type LogConfig struct {
	Level     string `yaml:"level"`
	Colors    bool   `yaml:"colors"`
	Console   bool   `yaml:"console"`
	File      bool   `yaml:"file"`
	Directory string `yaml:"directory"`
}
