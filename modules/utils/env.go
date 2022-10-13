package utils

import "github.com/spf13/viper"

type Config struct {
	ENV     string `mapstructure:"ENV"`
	DB_HOST string `mapstructure:"DB_HOST"`
	DB_PORT int    `mapstructure:"DB_PORT"`
	DB_PASS string `mapstructure:"DB_PASS"`
	DB_NAME string `mapstructure:"DB_NAME"`
	PORT    int    `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
