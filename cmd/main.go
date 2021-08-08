package main

import (
	"flag"
	"ggate/internal/config"
	"ggate/internal/proxy"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)

	configFile := ""
	flag.StringVar(&configFile, "config", "./config.yml", "config file")
	flag.Parse()
	cfg, err := config.LoadProxyConfig(configFile)
	if err != nil {
		logrus.Fatalln(err)
	}
	proxy := proxy.New(cfg)
	proxy.Run()
}
