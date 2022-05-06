package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

const yamlFile = "./config.yaml"

type Config struct {
	ConnectDb       string `yaml:"connect_db"`
	StartId         int    `yaml:"start_id"`
	Indent          uint64 `yaml:"indent"`
	PortionDb       uint64 `yaml:"portion_db"`
	PathFileCountId string `yaml:"path_file_count_id"`
}

type confRead struct {
	connectDb       string `yaml:"connect_db"`
	startId         string `yaml:"start_id"`
	indent          string `yaml:"indent"`
	portionDb       string `yaml:"portion_db"`
	pathFileCountId string `yaml:"path_file_count_id"`
}

func NenConfig() (*Config, error) {
	var confRead confRead
	var conf Config
	yamlFile, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &confRead)
	if err != nil {
		return nil, err
	}

	intStartId, err := strconv.Atoi(confRead.startId)
	if err != nil {
		return nil, err
	}

	uintIndent, err := strconv.ParseUint(confRead.indent, 10, 64)
	if err != nil {
		return nil, err
	}

	uintPortionDb, err := strconv.ParseUint(confRead.portionDb, 10, 64)
	if err != nil {
		return nil, err
	}

	conf.ConnectDb, conf.StartId = confRead.connectDb, intStartId
	conf.Indent, conf.PortionDb, conf.ConnectDb = uintIndent, uintPortionDb, confRead.pathFileCountId

	return &conf, nil
}
