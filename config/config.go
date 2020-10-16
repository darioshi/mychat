package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once
var Conf *Config

func init() {
	Init()
}

func Init() {
	once.Do(func() {
		env := GetMode()
		realPath, _ := filepath.Abs("./")
		configFilePath := realPath + "/config/" + env + "/"
		viper.SetConfigType("toml")
		viper.AddConfigPath(configFilePath)
		viper.SetConfigName("/site")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/connect")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		Conf = new(Config)
		viper.Unmarshal(&Conf.Site)
		viper.Unmarshal(&Conf.Connect)
	})
}

func GetMode() string {
	env := os.Getenv("RUN_MODE")
	if env == "" {
		env = "dev"
	}
	return env
}

type Config struct {
	Site    SiteConfig
	Connect ConnectConfig
}

type SiteBase struct {
	ListenPort int `mapstructure:"listenPort"`
}

type SiteConfig struct {
	SiteBase SiteBase `mapstructure:"site-base"`
}

type ConnectWebsocket struct {
	Bind string `mapstructure:"bind"`
}

type ConnectConfig struct {
	ConnectWebsocket ConnectWebsocket `mapstructure:"connect-websocket"`
}
