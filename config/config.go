// config/config.go
package config

import (
    "encoding/json"
    "io/ioutil"
)

type Config struct {
    LogDir    string
    LogConfig LogConfig
}

type LogConfig struct {
    MaxStoreBytes  uint64
    MaxIndexBytes  uint64
    InitialOffset  uint64
}

func LoadConfig(path string) (*Config, error) {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var config Config
    err = json.Unmarshal(file, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}
