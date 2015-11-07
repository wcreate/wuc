package main

import (
	"github.com/wcreate/wuc/handler"
	"gopkg.in/macaron.v1"
)

func main() {
	cfg := macaron.Config()
	web, err := cfg.GetSection("web")
	if err != nil {
		panic(err)
	}
	ip := web.Key("ip").MustString("0.0.0.0")
	port := web.Key("port").MustInt(8080)

	m := macaron.Classic()
	// user
	handler.InitHandles(m)

	m.Run(ip, port)
}
