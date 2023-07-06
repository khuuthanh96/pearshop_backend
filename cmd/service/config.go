package main

type Config struct {
	ENV      string `envconfig:"ENV" default:"development" required:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}
