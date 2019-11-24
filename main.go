package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	mbconfig "github.com/42wim/matterbridge/bridge/config"
	mbeconfig "github.com/jheiselman/matterbridge-editor/config"
	"github.com/sirupsen/logrus"
)

var (
	version = "0.0.1"

	flagConfig  = flag.String("conf", "matterbridge-web.toml", "config file")
	flagDebug   = flag.Bool("debug", false, "enable debug")
	flagVersion = flag.Bool("version", false, "show version and exit")

	logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{},
		Level:     logrus.InfoLevel,
	}
)

func main() {
	flag.Parse()
	if *flagVersion {
		fmt.Printf("version: %s\n", version)
		return
	}
	if *flagDebug {
		logger.SetLevel(logrus.DebugLevel)
		logger.Debug("Debugging output enabled")
	}

	logger.Info("Starting matterbridge-web UI")
	config := mbeconfig.ReadConfig(*flagConfig, logger)
	if config == nil {
		errstring := fmt.Sprintf("Failed to read configuration file: %s", *flagConfig)
		logger.Fatal(errstring)
	}

	if config.Matterbridge.ConfigPath == "" {
		logger.Fatal("config key in matterbridge section must be non-empty")
	}

	mb := readMBConfig(config.Matterbridge.ConfigPath)
	if mb == nil {
		logger.Fatal("Failed to load matterbridge config")
	}

	logger.Printf("%+v", mb)
}

func readMBConfig(path string) *mbconfig.Config {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		errstring := fmt.Sprintf("Failed to read matterbridge configuration file: %s", path)
		logger.Fatal(errstring)
	}

	config := mbconfig.NewConfigFromString(logger, input)

	return &config
}
