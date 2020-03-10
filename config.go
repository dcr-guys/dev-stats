package main

import (
	//	"fmt"
	"github.com/decred/dcrd/dcrutil/v2"
	flags "github.com/jessevdk/go-flags"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const defaultConfigFilename = "devstats.conf"

var (
	defaultDevstatsDir = dcrutil.AppDataDir("devstats", false)
	defaultConfigFile  = filepath.Join(defaultDevstatsDir, defaultConfigFilename)
)

type config struct {
	ConfigFile  string   `short:"C" long:"configfile" description:"Path to config file"`
	GithubToken string   `long:"token" description:"Your github token"`
	Users       []string `long:"user" description:"Map of all repos from that user"`
}

func loadConfig() (*config, error) {
	defaultConfig := config{
		ConfigFile: defaultConfigFile,
	}

	preCfg := defaultConfig
	if _, err := flags.Parse(&preCfg); err != nil {
		return nil, err
	}

	appName := filepath.Base(os.Args[0])
	appName = strings.TrimSuffix(appName, filepath.Ext(appName))
	// usageMessage := fmt.Sprintf("Use %s -h to show usage", appName)

	var configFileError error
	cfg := preCfg
	if err := flags.IniParse(defaultDevstatsDir, &cfg); err != nil {
		if _, ok := err.(*flags.IniError); ok {
			return nil, err
		}
		configFileError = err
	}
	if _, err := flags.Parse(&cfg); err != nil {
		return nil, err
	}

	if configFileError != nil {
		log.Printf("%v", configFileError)
	}

	return &cfg, nil
}
