package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config is a type for handling information from a configuration file.
type Config struct {
	Path     string
	Database string `json:"database"`
	Gallery  string `json:"gallery"`
}

//SetPath is hardcoded path to configuration file.
func (c *Config) SetPath() {
	c.Path = "./config.json"
}

//ReadConfig sets config-fields.
func (c *Config) ReadConfig() {
	file, _ := ioutil.ReadFile(c.Path)
	err := json.Unmarshal([]byte(file), &c)
	if err != nil {
		panic(err)
	}
}

//DatabasePath returns path to SQLite databasefile.
func (c *Config) DatabasePath() string {
	return c.Database
}

//GalleryPath returns path to gallery.
func (c *Config) GalleryPath() string {
	return c.Gallery
}

//NewConfig constructs the config object.
func NewConfig() *Config {
	Conf := Config{}
	Conf.SetPath()
	Conf.ReadConfig()
	return &Conf
}
