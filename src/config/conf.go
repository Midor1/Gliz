package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type Conf struct {
	Addresses      AddressConf       `toml:"Addresses"`
	Authenticators AuthenticatorConf `toml:"Authenticators"`
	Keys           KeyConf           `toml:"Keys"`
	//TODO: Make auth info irrelevant of terminals.
}

type AddressConf struct {
	SQLAddr   string `toml:"MySQLAddress"`
	RedisAddr string `toml:"RedisAddress"`
	SrvAddr   string `toml:"ServerAddress"`
}

type AuthenticatorConf struct {
	SQLUserName string `toml:"MySQLUserName"`
	SQLPassword string `toml:"MySQLPassword"`
}

type KeyConf struct {
	SessionAuthKey string `toml:"SessionAuthenticationKey"`
}

func (c *Conf) GetConfig() (*Conf, error) {
	confFile, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return c, err
	}
	err = toml.Unmarshal(confFile, c)
	if err != nil {
		return c, err
	}
	return c, nil
}

var C Conf
