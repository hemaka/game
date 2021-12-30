package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var configFile = "./config.json"

type Config struct {
	Port  string
	Redis struct {
		Host string
		Port string
	}
	Mysql struct {
		Host     string
		Port     string
		DBName   string
		User     string
		Password string
	}
}

func (this *Config) Load() {
	data, err := ioutil.ReadFile(configFile)

	if err == nil {
		err = json.Unmarshal(data, this)
	}

	if err != nil {
		log.Printf("Error parsing %s: %v", configFile, err)
	}

	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(configFile); err != nil {
		executablePath, err := filepath.Abs(filepath.Dir(executable))
		if err != nil {
			panic(err)
		}

		log.Printf("Going to folder %v...", executablePath)

		os.Chdir(executablePath)
	}
}
