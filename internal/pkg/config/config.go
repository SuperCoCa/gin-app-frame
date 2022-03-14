package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-app-frame/internal/pkg/global/variable"
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	App      App      `mapstructure:"app"`
	Database Database `mapstructure:"database"`
	Log      Log      `mapstructure:"log"`
	Redis    Redis    `mapstructure:"redis"`
}

var cfg = new(Config)

func InitConfig() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(variable.CONFIG_PATH)
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config failed: %s", err))
	}

	if err = vp.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("read config failed: %s", err))
	}

	// 监听并重载配置
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err = vp.Unmarshal(&cfg); err != nil {
			fmt.Println(err)
		}
	})

	return vp
}

// @description 返回配置
func GetConfig() *Config {
	return cfg
}
