package common

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var config map[string]string

var configFilePath = filepath.Join(WORKSPACE, "config.json")

func init() {
	bytes, err := os.ReadFile(configFilePath)
	OnError(err, "failed to read %s", configFilePath)
	err = json.Unmarshal(bytes, &config)
	OnError(err, "failed to init configuration from %s", configFilePath)
}

func TryGetConfig(key string, defaultValue string) string {
	if v, ok := config[key]; ok {
		log.Printf("key='%s', config='%s'(from file)", key, v)
		return v
	} else {
		log.Printf("key='%s', config='%s'(from defaultValue)", key, defaultValue)
		return defaultValue
	}
}
