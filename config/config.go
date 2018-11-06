package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/rhasan33/goplate/constants"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// ServerEnv for the app dev/production
const (
	ServerEnv = "dev"
)

var muLock sync.Mutex

// Init loads the initial configuration
func Init() {
	viper.SetEnvPrefix(constants.AppName)
	viper.BindEnv("env")

	// set up consul for app ennvironment
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	consulURL := viper.GetString("consul_url")
	consulPath := viper.GetString("consul_path")

	// if consulURL == "" {
	// 	log.Fatal("CONSUL_URL is missing.")
	// }

	// if consulPath == "" {
	// 	panic("CONSUL_PATH is not defined.")
	// }

	viper.AddRemoteProvider("consul", consulURL, consulPath)
	viper.SetConfigType("yml")

	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("%s named \"%s\"", err.Error(), consulPath))
	}

	// load app db and redis
	LoadApplication()
	LoadDB()
	LoadRedis()
}
