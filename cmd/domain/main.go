package main

import (
	"context"
	"runtime"

	"github.com/jmbarzee/dominion/domain"
	"github.com/jmbarzee/dominion/domain/config"
	"github.com/jmbarzee/dominion/system"
)

func main() {
	runtime.GOMAXPROCS(4)

	configFileName := system.RequireEnv("DOMAIN_CONFIG_FILE")

	// Check config
	if err := config.SetupFromTOML(configFileName); err != nil {
		panic(err)
	}

	domain, err := domain.NewDomain(config.GetDomainConfig())
	if err != nil {
		panic(err)
	}

	if err := domain.Run(context.Background()); err != nil {
		panic(err)
	}
}
