package utils

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
)

type Config struct {
	Application Application
	InfluxDB InfluxDB
}

type Application struct {
	Title string
	Version string
	ReleaseDate string
	Port int
	Key string
	Cert string
	Interval int
}

type InfluxDB struct {
	DBName string
	Username string
	Password string
	Address string
}

func LoadConfig() *Config {
	file, _ := ioutil.ReadFile("config.json")
	var configuration Config
	_ = json.Unmarshal(file, &configuration)
	log.Println("application configuration loaded")
	fmt.Println(configuration)
	return &configuration
}