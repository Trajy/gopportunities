package main

import (
	"github.com/Trajy/gopportunities/config"
	"github.com/Trajy/gopportunities/router"
)

func main() {
	config.GetLogger().Info("Init Gopportunities application")
	router.Init()
}
