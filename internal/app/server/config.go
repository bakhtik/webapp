package server

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbConnString string `json:"db_conn_string"`
	CacheAddr    string `json:"cache_address"`
	ServerAddr   string `json:"server_address"`
}

func (c *Config) FromFile(filename string) (err error) {
	input, err := os.Open(filename)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(input)
	if err = dec.Decode(&c); err != nil {
		return err
	}
	return nil
}
