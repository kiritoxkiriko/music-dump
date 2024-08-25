package config

import "github.com/spf13/viper"

var Conf Config

type Config struct {
	Dumper   string `mapstructure:"dumper"`
	Format   string `mapstructure:"format"`
	DestPath string `mapstructure:"dest_path"`
	SrcPath  string `mapstructure:"src_path"`
	LogLevel string `mapstructure:"log_level"`
}

func InitConfig(configPath string) error {
	// add config search paths
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("~/.music-dumper")

	// add config file name
	viper.SetConfigName("config")

	// set default values
	SetDefaults()

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func SetDefaults() {
	viper.SetDefault("dumper", "auto")
	//viper.SetDefault("Format", "mp3")
	viper.SetDefault("dest_path", "./out")
	viper.SetDefault("src_path", "./")
	viper.SetDefault("log_level", "info")
}
