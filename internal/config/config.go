// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package config

import (
	"fmt"
	"github.com/golangee/sql"
	"github.com/worldiety/mercurius/internal/errors"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Filename is the name of the base settings file
const Filename = "mercurius.yml"

// Yaml contains a default example configuration file
const Yaml = `# mercurius base configuration

# the relational database
database:
 driver: mysql # currently only "mysql" is supported
 host: localhost # the host of the database
 port: 3306 # port of the db
 user: # the database user
 databaseName: # the database name
 sslMode: preferred # 0=preferred, 1=disabled, 2=required, 3=verify_ca, 4=verify_identity

# the blob store to use
blobStore:
 driver: filesystem # currently only "filesystem" is supported
 path: ./workspace # absolute or relative path for the "filesystem" driver

# the http server settings
server:
 port: 8080
 address: localhost

# logging setting
logging:
 development: false # if true, enables colorized non-json output
`

// BlobStore contains the blob store settings
type BlobStore struct {
	Driver string `yaml:"driver"`
	Path   string `yaml:"path"`
}

// Settings contains the minimal bootstrapping configuration to reach all data stores which itself contains
// all other configurations
type Settings struct {
	BlobStore   BlobStore `yaml:"blobStore"`
	Database    sql.Opts  `yaml:"database"`
	Server      Server    `yaml:"server"`
	Development bool      `yaml:"development"`
}

// Server contains the http server specific settings
type Server struct {
	Port    int    `yaml:"port"`
	Address string `yaml:"address"`
}

// Logging settings
type Logging struct {
	Development bool `yaml:"development"`
}

// Default returns the default configuration, which only makes sense for a developer machine, if anything at all.
func Default() Settings {
	wd, _ := os.UserHomeDir()
	return Settings{
		BlobStore: BlobStore{
			Driver: "filesystem",
			Path:   filepath.Join(wd, "workspace"),
		},
		Database: sql.Opts{
			Driver:       "mysql",
			Host:         "localhost",
			Port:         3306,
			User:         "root",
			Password:     "",
			DatabaseName: "mercurius",
			SSLMode:      sql.SSLPreferred,
		},
		Server: Server{
			Port:    8080,
			Address: "localhost",
		},
		Development: false,
	}
}

// LoadFile decodes a yaml configuration from a file
func LoadFile(filename string) (s Settings, err error) {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return Default(), fmt.Errorf("config file '%s'cannot be opened: %w", filename, err)
	}

	defer errors.Try(file.Close, &err)

	r, err := Load(file)
	if err != nil {
		return Default(), fmt.Errorf("config file '%s' cannot be parsed: %w", filename, err)
	}

	return r, nil
}

// Load tries to decode the config from yaml format
func Load(reader io.Reader) (Settings, error) {
	r := Default()

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return r, err
	}

	err = yaml.Unmarshal(b, &r)

	return r, err
}

// SaveFile tries to serialize the settings into the writer
func SaveFile(filename string, settings Settings) (err error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("cannot write config '%s': %w", filename, err)
	}
	defer errors.Try(f.Close, &err)

	err = Save(f, settings)

	return err
}

// Save serializes the settings into the writer
func Save(writer io.Writer, settings Settings) error {
	b, err := yaml.Marshal(settings)
	if err != nil {
		panic(err)
	}

	_, err = writer.Write(b)

	return err
}
