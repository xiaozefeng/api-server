package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Init(filename string) error {

	if err := initConfig(filename); err != nil {
		return err
	}
	watch()

	if err := initLog(); err != nil {
		return err
	}
	log.Infof("init log successfully")
	return nil
}

func watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config file changed: %s", e.Name)
	})
}

func initConfig(filename string) error {
	if filename != "" {
		viper.SetConfigFile(filename)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	return viper.ReadInConfig()
}

func initLog() error {
	log.SetFormatter(&log.JSONFormatter{})
	logPath := viper.GetString("log_path")
	dist, err := os.Create(logPath)
	if err != nil {
		return err
	}
	mw := io.MultiWriter(os.Stdout, dist)
	log.SetOutput(mw)
	log.SetLevel(log.InfoLevel)
	return nil
}
