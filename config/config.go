// config/config.go
package config

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

// Config represents the configuration structure for the entire logging system.
type Config struct {
    Segment struct {
        MaxStoreBytes  uint64 `json:"maxStoreBytes"`
        MaxIndexBytes  uint64 `json:"maxIndexBytes"`
        InitialOffset  uint64 `json:"initialOffset"`
    } `json:"segment"`
    LogDir string `json:"logDir"`  // Directory where log files will be stored.
}

// LoadConfig reads a JSON file at the given path and decodes it into a Config struct.
func LoadConfig(path string) (Config, error) {
    var config Config
    configFile, err := os.Open(path)
    if err != nil {
        return config, err
    }
    defer configFile.Close()

    bytes, err := ioutil.ReadAll(configFile)
    if err != nil {
        return config, err
    }

    err = json.Unmarshal(bytes, &config)
    if err != nil {
        return config, err
    }

    return config, nil
}
