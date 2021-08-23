package main

import (
	"api-server/config"
	"api-server/router"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
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
	if err!=nil {
		panic(err)
	}


	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	var middlewares []gin.HandlerFunc
	router.Load(g, middlewares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("the router has no response, or it might took too long to start up.", err)
		}
		log.Print("the router has been deployed successfully")
	}()
	port:=viper.GetString("port")
	log.Printf("start to listening the incoming requests on http address: %s", port)
	log.Printf(http.ListenAndServe(port, g).Error())

}

func pingServer() error {
	url := "http://127.0.0.1" +viper.GetString("port") + "/sd/health"
	for i := 0; i < 2; i++ {
		response, err := http.Get(url)
		if err == nil && response.StatusCode == http.StatusOK {
			return nil
		}
		log.Print("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to router")
}
