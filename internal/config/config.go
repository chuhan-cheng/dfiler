// internal/config/config.go
package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    ServerPort string `json:"server_port"`
    DBURL      string `json:"db_url"`
}

func Load(path string) (*Config, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    cfg := &Config{}
    err = decoder.Decode(cfg)
    if err != nil {
        return nil, err
    }

    return cfg, nil
}
