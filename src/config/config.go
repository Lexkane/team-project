package config

import (
	//import of database
	"bytes"
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	_ "../database"
)

// Config is variable for config
var (
	Config   Configuration
	FilePath = "team_project_config.json"
)

// Configuration is a singleton object for application config
type Configuration struct {
	ListenURL   string `json:"ListenURL"`
	LogFilePath string `json:"LogFilePath"`
	Database    sql.DB `json:"Database"`
}

// Load loads config once
func Load() error {
	err := readFromJSON(FilePath)
	if err != nil {
		return errors.New("Setup configurations please")
	}

	return nil
}

// readFromJSON reads config data from JSON-file
func readFromJSON(configFilePath string) error {
	log.Printf("Searching for JSON config file (%s)", configFilePath)

	contents, err := ioutil.ReadFile(configFilePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&Config)
	}
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %s\n", configFilePath, err)
	} else {
		log.Printf("Configuration loaded  from JSON (%s) successfully\n", configFilePath)
	}

	return err
}
