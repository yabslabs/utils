package config

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yabslabs/utils/config"
	"github.com/yabslabs/utils/config/json"
)

// LoggerApplyToGlobal can be set before the config is read
// if false the configs are not applied to the global logger
var LoggerApplyToGlobal = true

type Log struct {
	Logger *log.Logger
}

type logConfig struct {
	Level     string
	Formatter string
}

func (l *Log) UnmarshalJSON(data []byte) error {
	return l.Unmarshal(data, json.ConfigReader)
}

func (l *Log) UnmarshalText(data []byte) error {
	// cant parse toml since raw values (objects) are not supported
	return l.Unmarshal(data, json.ConfigReader)
}

func (l *Log) Unmarshal(data []byte, reader config.ConfigReader) error {
	conf := &logConfig{}
	if err := reader.Unmarshal(data, conf); err != nil {
		return err
	}

	formatter, err := parseLogFormatter(conf.Formatter)
	if err != nil {
		return err
	}

	if conf.Level == "" {
		conf.Level = "info"
	}

	lvl, err := log.ParseLevel(conf.Level)
	if err != nil {
		return err
	}

	if LoggerApplyToGlobal {
		log.SetLevel(lvl)
		log.SetFormatter(formatter)
	}

	l.Logger = log.New()
	l.Logger.Level = lvl
	l.Logger.Formatter = formatter

	return nil
}

func parseLogFormatter(formatterDescription string) (log.Formatter, error) {
	switch formatterDescription {
	case "json":
		return &log.JSONFormatter{}, nil
	case "text", "":
		return &log.TextFormatter{}, nil
	case "text-color":
		return &log.TextFormatter{ForceColors: true}, nil
	default:
		return nil, status.Error(codes.Internal, formatterDescription+" formatter not supported")
	}
}
