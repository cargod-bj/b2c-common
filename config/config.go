package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	Logger     *oplogging.Logger
	Log_Config Log
	GVA_VP     *viper.Viper
)

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&Log_Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&Log_Config); err != nil {
		fmt.Println(err)
	}
	GVA_VP = v
}
