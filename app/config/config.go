package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Init command line arg and Paser runtime configuration
func Init() {
	file := pflag.StringP("config", "c", "config.yaml", "the configuration file path,default is config.yaml")
	pflag.Parse()
	// f := strings.Split(*file, ".")
	// viper.AddConfigPath(".")
	viper.SetConfigFile(*file)
	// viper.SetConfigName(f[0])
	// viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file: %s" + err.Error())
	}
}
