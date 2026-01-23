package config

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var cfg zap.Config

var rawJSON = []byte(`{
	"level": "debug",
	"encoding": "console",
	"outputPaths": ["stdout"],
	"errorOutputPaths": ["stderr"],
	"encoderConfig": {
		"messageKey": "msg",
		"levelKey": "level",
		"timeKey": "ts",
		"callerKey": "caller",
		"levelEncoder": "capitalColor",
		"timeEncoder": "iso8601",
		"durationEncoder": "ms",
		"callerEncoder": "short",
		"consoleSeparator": "  "
	}
}`)

func StartLogger() {
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("02/01/2006 15:04:05")

	Logger = zap.Must(cfg.Build())

	Logger.Info("Logger initialized")
}
