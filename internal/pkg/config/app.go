package config

type App struct {
	LogSavePath string `mapstructure:"log-save-path"`
	LogFileExt  string `mapstructure:"log-file-ext"`
}
