package config

type Database struct {
	DbDriver           string `mapstructure:"db-driver"`
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	Host               string `mapstructure:"host"`
	DbName             string `mapstructure:"db-name"`
	TablePrefix        string `mapstructure:"table-prefix"`
	Charset            string `mapstructure:"charset"`
	ParseTime          bool   `mapstructure:"parse-time"`
	MaxIdleConnections int    `mapstructure:"max-idle-connections"`
	MaxOpenConnections int    `mapstructure:"max-open-connections"`
	LogLevel           string `mapstructure:"log-level"`
	LogFileName        string `mapstructure:"log-file-name"`
	LogFileExt         string `mapstructure:"log-file-ext"`
	LogMaxAge          int    `mapstructure:"log-max-age"`
}
