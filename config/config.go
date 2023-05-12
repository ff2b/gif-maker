package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

type Config struct {
	GIFRate         int
	GIFLoop         bool
	DefaultSavePath string
	WorkSpace       string
}

func NewConfig() *Config {
	config := &Config{}
	return config
}

func (c *Config) Load() {
	data, err := os.ReadFile(path.Join("config", "config.json"))
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}

func (c *Config) Save() {
	jsonData, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", jsonData)
	os.WriteFile(path.Join("config", "config.json"), jsonData, 0666)
}
