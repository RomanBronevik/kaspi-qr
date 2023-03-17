package config

import (
	"github.com/spf13/viper"
)

type confSt struct {
	Debug       bool   `mapstructure:"DEBUG"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	HttpListen  string `mapstructure:"HTTP_LISTEN"`
	HttpCors    bool   `mapstructure:"HTTP_CORS"`
	PgDsn       string `mapstructure:"PG_DSN"`
	KaspiApiUrl string `mapstructure:"KASPI_API_URL"`
	StgApiUrl   string `mapstructure:"STG_API_URL"`
}

func Load() *confSt {
	result := &confSt{}

	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("HTTP_LISTEN", ":80")
	viper.SetDefault("KASPI_API_URL", "https://mtokentest.kaspi.kz:8545/r3/v01")
	viper.SetDefault("STG_API_URL", "https://api.mechta.market/stg")

	viper.SetConfigFile("conf.yml")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()

	_ = viper.Unmarshal(&result)

	return result
}
