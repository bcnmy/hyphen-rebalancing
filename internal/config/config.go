package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const (
	defaultJobTimeout = 60
)

func FromFile(filename string) (*Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	conf := new(Config)
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}

	postProcess(conf)

	return conf, nil
}

func postProcess(conf *Config) {
	if conf.General.JobTimeout == 0 {
		conf.General.JobTimeout = defaultJobTimeout
	}
}
