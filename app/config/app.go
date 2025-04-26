package config

import "time"

type App struct {
	ServiceName    string        `envconfig:"APP_SERVICE_NAME" required:"false"`
	Env            string        `envconfig:"APP_ENV" required:"false"`
	APIKey         string        `envconfig:"APP_API_KEY" required:"false"`
	ContextTimeout time.Duration `envconfig:"CONTEXT_TIMEOUT" required:"false"`
}
