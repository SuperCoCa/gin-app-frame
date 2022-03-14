package config

import "time"

type Server struct {
	RunMode      string        `mapstructure:"host"`
	HttpPort     string        `mapstructure:"http-port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout"`
}
