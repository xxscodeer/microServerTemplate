package tools

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	App    AppConfig    `yaml:"app"`
	Etcd   EtcdConfig   `yaml:"etcd"`
	Mysql  MysqlConfig  `yaml:"mysql"`
	Redis  RedisConfig  `yaml:"redis"`
	Jaeger JaegerConfig `yaml:"jaeger"`
}

type AppConfig struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Name    string `yaml:"name"`
	Mode    string `yaml:"mode"`
	Version string `yaml:"version"`
}

type EtcdConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type MysqlConfig struct {
	DbName    string `yaml:"dbName"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DbUser    string `yaml:"dbUser"`
	DbPwd     string `yaml:"dbPwd"`
	Loc       string `yaml:"loc"`
	ParseTime string `yaml:"parseTime"`
	Charset   string `yaml:"charset"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type JaegerConfig struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	MicroName string `yaml:"name"`
}

func ParseConfig(path string) (cfg *Config) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("open file err,", err)
	}
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		log.Fatalln("unmarshal fail err,", err)
	}
	return
}
