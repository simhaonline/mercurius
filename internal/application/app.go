// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package application

import (
	sql2 "database/sql"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	srv "github.com/golangee/http"
	"github.com/golangee/log"
	"github.com/golangee/sql"

	"github.com/worldiety/mercurius/internal/config"
	"github.com/worldiety/mercurius/internal/service/setup"
	"os"
)

type Server struct {
	cfgFile             string
	logger              log.Logger
	settings            config.Settings
	db                  *sql2.DB
	configurationErrors []error
}

func (a *Server) ReloadStatus() []error {
	return a.configurationErrors
}

func (a *Server) Reload() {
	a.Configure(a.cfgFile)
}

func NewServer() *Server {
	a := &Server{logger: log.New("server")}
	setup.NewRestController(a)
	return a
}

func (a *Server) Configure(cfgFile string) {
	a.cfgFile = cfgFile
	a.settings = config.Default()

	a.logger.Info("configure from settings file", log.Obj("file", cfgFile))
	a.configurationErrors = nil
	if a.db != nil {
		err := a.db.Close()
		a.logger.Error("failed to close database", log.Obj("err", err))
		a.db = nil
	}

	// if cfg file does not exist, note and return
	if _, err := os.Stat(cfgFile); err != nil {
		a.logger.Error("no configuration file", log.Obj("file", cfgFile), log.Obj("err", err))
		a.configurationErrors = append(a.configurationErrors, config.FirstTimeSetupError{})
		return
	}

	// try to load it, actually something is there
	cfg, err := config.LoadFile(cfgFile)
	if err != nil {
		a.logger.Error("unable to load config", log.Obj("err", err))
		a.configurationErrors = append(a.configurationErrors, config.InvalidConfigurationError{Cause: err})
	}
	a.settings = cfg

	db, err := sql.Open(a.settings.Database)
	if err != nil {
		a.logger.Error("unable to open database", log.Obj("err", err))
		a.configurationErrors = append(a.configurationErrors, config.NoDatabaseError{Cause: err})
	}
	a.db = db
}

func (a *Server) StartDev(frontendDir string) {
	a.startSrv(frontendDir, a.settings.Server.Port)
}

func (a *Server) initControllers(server *srv.Server) {
	srv.MustNewController(server, setup.NewRestController(a))
}
