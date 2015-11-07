package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/wcreate/wuc/handler"
	"gopkg.in/macaron.v1"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {

	log.Debug("Starting server...")

	cfg := macaron.Config()
	web, err := cfg.GetSection("web")
	if err != nil {
		panic(err)
	}
	ip := web.Key("ip").MustString("0.0.0.0")
	port := web.Key("port").MustInt(8080)

	m := macaron.Classic()
	m.Use(macaron.Renderer())
	// user
	handler.InitHandles(m)

	m.Run(ip, port)
}
