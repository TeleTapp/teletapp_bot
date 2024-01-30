package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var App = app{}

func Load() {
	godotenv.Load()

	if err := env.Parse(&App); err != nil {
		log.Fatalf("parse env error %+v\n", err)
	}

}

type app struct {
	Listen        string `env:"LISTEN" envDefault:":3000"`
	ApiKey        string `env:"API_KEY,required"`
	ApiBaseURL    string `env:"API_BASE_URL,required"`
	Token         string `env:"TOKEN,required"`
	WebAppBaseURL string `env:"WEB_APP_BASE_URL,required"`
}
