// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package main

import (
	"flag"
	"github.com/golangee/log"
	zap "github.com/golangee/log-zap"
	_ "github.com/worldiety/mercurius" // reflectplus metadata
	"github.com/worldiety/mercurius/build"
	"github.com/worldiety/mercurius/internal/application"
	"github.com/worldiety/mercurius/internal/config"
	"os"
	"path/filepath"
)

func main() {
	zap.Configure()
	logger := log.New("")
	logger.Info("mercurius", log.Obj("buildCommit", build.Commit), log.Obj("buildTime", build.Time))

	dir, err := os.UserConfigDir()
	if err != nil {
		dir, err = os.UserHomeDir()
		if err != nil {
			dir, err = os.Getwd()
			if err != nil {
				panic(err)
			}
		}
	}

	cfgFile := flag.String("cfg", filepath.Join(dir, config.Filename), "the config file to use")
	help := flag.Bool("help", false, "shows this help")
	frontendDir := flag.String("devFrontend", "", "dev-only: absolute path to a directory with index.html and wasm file")

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	app := application.NewServer()
	app.Configure(*cfgFile)
	if *frontendDir != "" {
		logger.Warn("serving development frontend", log.Obj("dir", *frontendDir))
		app.StartDev(*frontendDir)
	}
}
