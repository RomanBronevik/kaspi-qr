package main

import (
	"github.com/rendau/dop/dopTools"
	"github.com/spf13/viper"
)

type ConfSt struct {
	Debug         bool   `mapstructure:"DEBUG"`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	HttpListen    string `mapstructure:"HTTP_LISTEN"`
	HttpCors      bool   `mapstructure:"HTTP_CORS"`
	SwagHost      string `mapstructure:"SWAG_HOST"`
	SwagBasePath  string `mapstructure:"SWAG_BASE_PATH"`
	SwagSchema    string `mapstructure:"SWAG_SCHEMA"`
	PgDsn         string `mapstructure:"PG_DSN"`
	KaspiApiUrl   string `mapstructure:"KASPI_API_URL"`
	CertPath      string `mapstructure:"CERT_PATH"`
	CertPsw       string `mapstructure:"CERT_PSW"`
	QrUrlTemplate string `mapstructure:"QR_URL_TEMPLATE"`
}

func ConfLoad() *ConfSt {
	result := &ConfSt{}

	dopTools.SetViperDefaultsFromObj(result)

	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("HTTP_LISTEN", ":80")
	viper.SetDefault("SWAG_HOST", "example.com")
	viper.SetDefault("SWAG_BASE_PATH", "/")
	viper.SetDefault("SWAG_SCHEMA", "https")
	viper.SetDefault("KASPI_API_URL", "https://mtokentest.kaspi.kz:8545/r3/v01/")
	viper.SetDefault("CERT_PATH", "")
	viper.SetDefault("CERT_PSW", "")
	viper.SetDefault("QR_URL_TEMPLATE", "")

	viper.SetConfigFile("conf.yml")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()

	_ = viper.Unmarshal(&result)

	return result
}
