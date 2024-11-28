package main

import (
	"github.com/acc-event-manager/accweb/internal/app"
	"github.com/acc-event-manager/accweb/internal/pkg/cfg"
	"github.com/acc-event-manager/accweb/internal/pkg/helper"
	"github.com/acc-event-manager/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

const configFile = "config.yml"

func main() {
	c := cfg.Load(configFile)

	sM := server_manager.New(c.ConfigPath, c.ACC.ServerPath, c.ACC.ServerExe)

	logrus.Info("accweb: checking for secrets...")
	helper.GenerateTokenKeysIfNotPresent(c.Auth.PublicKeyPath, c.Auth.PrivateKeyPath)

	logrus.Info("accweb: initializing...")
	if err := sM.Bootstrap(); err != nil {
		logrus.WithError(err).Fatal("failed to bootstrap accweb")
	}

	logrus.WithField("addr", c.Webserver.Host).Info("initializing web server")
	app.StartServer(c, sM)
}
