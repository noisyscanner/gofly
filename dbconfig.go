package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
)

type DBConfig struct {
	Driver string
	Host string
	User string
	Pass string
	Port int
	Db string
}

func (c *DBConfig) DBString() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", c.User, c.Pass, c.Host, c.Port, c.Db)
}

type ConfigService interface {
	GetConfig() *DBConfig
}

type FakeConfigService struct {}

func (_ *FakeConfigService) GetConfig() *DBConfig {
	return &DBConfig{
		Driver: "mysql",
		Host: "localhost",
		User: "root",
		Pass: "ufx366",
		Port: 3306,
		Db: "reed.brad_iVerbs",
	}
}

type FileConfigService struct {
	File string
}

func (s FileConfigService) GetConfig() *DBConfig {
	cnf, err := s.parseConfig(s.File)
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(cnf["Port"])
	if err != nil {
		panic(err)
	}

	return &DBConfig{
		Driver: cnf["Driver"],
		Host: cnf["Host"],
		User: cnf["User"],
		Pass: cnf["Pass"],
		Port: port,
		Db: cnf["Db"],
	}
}

func (s FileConfigService) parseConfig(c string) (map[string]string, error) {
	var out = map[string]string{}
	conf, err := ioutil.ReadFile(s.File)
	if err != nil {
		return out, err
	}


	pairs := strings.Split(string(conf), "\n")
	for _, pair := range pairs {
		split := strings.Split(pair, " ")
		if len(split) == 2 {
			k, v := split[0], split[1]
			out[k] = v
		}
	}

	return out, nil
}