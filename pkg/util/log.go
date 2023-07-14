package util

import (
	"encoding/json"

	"go.uber.org/zap"
)

var Log *zap.Logger

func SetupLogging() {
	configJson := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)
	var cfg zap.Config
	if err := json.Unmarshal(configJson, &cfg); err != nil {
		panic(err)
	}

	var err error

	Log, err = cfg.Build()
	if err != nil {
		panic(err)
	}
	defer Log.Sync()
}
