package config

import (
	"time"

	"github.com/spf13/viper"
)

// Version used to reflect the app version
var Version = "unversioned"

// Application type for the app environment
type Application struct {
	Env            string
	Port           string
	Sentry         string
	Version        string
	ReadTimeout    time.Duration
	RequestTimeout time.Duration
}

var app Application

// App sets the application envs
func App() *Application {
	return &app
}

// LoadApplication loads the application environments
func LoadApplication() {
	// set up mutex lock
	muLock.Lock()
	defer muLock.Unlock()

	// setup environment
	env := ServerEnv
	if e := viper.GetString("app.env"); e != "" {
		env = e
	}

	// setup versioning
	version := Version
	if v := viper.GetString("app.version"); v != "" {
		version = v
	}

	app = Application{
		Env:            env,
		Port:           viper.GetString("app.port"),
		Sentry:         viper.GetString("app.sentry"),
		Version:        version,
		ReadTimeout:    viper.GetDuration("app.read_timeout") * time.Second,
		RequestTimeout: viper.GetDuration("app.request_timeout") * time.Second,
	}
}
