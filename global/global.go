package global

import (
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	COM_LOG    *oplogging.Logger
	COM_VP     *viper.Viper
	Log_Config Log
)

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}
