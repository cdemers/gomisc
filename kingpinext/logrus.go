package kingpinext

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"strings"
)

var (
	// Flags parsed with kingpin.
	logFlag = kingpin.Flag("log", "Coma separated logging modes, for example \"--log debug,json\". (debug, info, warn, error, fatal, color, nocolor and json)").String()
)

func LogrusFlagsParse() {
	if *logFlag != "" {
		modes := strings.Split(*logFlag, ",")

		for _, mode := range modes {
			logrusLevel, err := log.ParseLevel(mode)
			if err == nil {
				log.SetLevel(logrusLevel)
			}
			if mode == "json" {
				log.SetFormatter(&log.JSONFormatter{})
			}
			if mode == "color" {
				log.SetFormatter(&log.TextFormatter{ForceColors: true})
			}
			if mode == "nocolor" {
				log.SetFormatter(&log.TextFormatter{DisableColors: true})
			}
		}
	}
}
