package main

import (
	"api-server/config"
	"api-server/db"
	"api-server/router"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"net/http"
	"time"
)

var (
	cfg string
)

func main() {
	flag.StringVar(&cfg, "c", "config.yaml", "config file")
	flag.Parse()

	err := config.Init(cfg)
	if err != nil {
		panic(err)
	}
	log.Info("init config successfully")

	// init DB
	err = db.Init()
	if err != nil {
		panic(err)
	}
	log.Info("init db successfully")

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	var middlewares []gin.HandlerFunc
	router.Load(g, middlewares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("the router has no response, or it might took too long to start up.", err)
		}
		log.Info("the router has been deployed successfully")
	}()
	port := viper.GetString("port")
	log.Infof("start to listening the incoming requests on http address: %s", port)
	log.Infof(http.ListenAndServe(port, g).Error())

}

func pingServer() error {
	url := "http://127.0.0.1" + viper.GetString("port") + "/sd/health"
	for i := 0; i < 2; i++ {
		response, err := http.Get(url)
		if err == nil && response.StatusCode == http.StatusOK {
			return nil
		}
		log.Info("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to router")
}
