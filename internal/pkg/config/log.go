package config

type Log struct {
	FileName string `mapstructure:"file-name"`
	FileExt  string `mapstructure:"file-ext"`
	Level    string `mapstructure:"level"`
	ShowLine bool   `mapstructure:"show-line"`
	MaxAge   int    `mapstructure:"max-age"`
}
