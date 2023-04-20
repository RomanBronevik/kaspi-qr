package config

import (
	"github.com/spf13/viper"
)

type ConfSt struct {
	Debug       bool   `mapstructure:"DEBUG"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	HttpListen  string `mapstructure:"HTTP_LISTEN"`
	HttpCors    bool   `mapstructure:"HTTP_CORS"`
	PgDsn       string `mapstructure:"PG_DSN"`
	KaspiApiUrl string `mapstructure:"KASPI_API_URL"`
	CertPath    string `mapstructure:"CERT_PATH"`
	CertPsw     string `mapstructure:"CERT_PSW"`
}

func Load() *ConfSt {
	result := &ConfSt{}

	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("HTTP_LISTEN", ":80")
	viper.SetDefault("KASPI_API_URL", "https://mtokentest.kaspi.kz:8545/r3/v01/")
	viper.SetDefault("CERT_PATH", "")
	viper.SetDefault("CERT_PSW", "")

	viper.SetConfigFile("conf.yml")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()

	_ = viper.Unmarshal(&result)

	return result
}
