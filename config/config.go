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
	Indent          uint64 `yaml:"Indent"`
	PortionDb       uint64 `yaml:"portion_db"`
	PathFileCountId string `yaml:"path_file_count_id"`
	CsvFilePath     string `yaml:"csv_file_path"`
}

type ConfRead struct {
	ConnectDb       string `yaml:"connect_db"`
	StartId         string `yaml:"start_id"`
	Indent          string `yaml:"indent"`
	PortionDb       string `yaml:"portion_db"`
	PathFileCountId string `yaml:"path_file_count_id"`
	CsvFilePath     string `yaml:"csv_file_path"`
}

func NenConfig() (*Config, error) {
	var confRead ConfRead
	var conf Config
	yamlFile, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &confRead)
	if err != nil {
		return nil, err
	}

	intStartId, err := strconv.Atoi(confRead.StartId)
	if err != nil {
		return nil, err
	}

	uintIndent, err := strconv.ParseUint(confRead.Indent, 10, 64)
	if err != nil {
		return nil, err
	}

	uintPortionDb, err := strconv.ParseUint(confRead.PortionDb, 10, 64)
	if err != nil {
		return nil, err
	}

	conf.ConnectDb, conf.StartId = confRead.ConnectDb, intStartId
	conf.Indent, conf.PortionDb, conf.PathFileCountId, conf.CsvFilePath = uintIndent, uintPortionDb, confRead.PathFileCountId, confRead.CsvFilePath

	return &conf, nil
}
