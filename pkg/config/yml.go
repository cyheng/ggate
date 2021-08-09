package config

import (
	"bytes"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadYaml(conf string, val interface{}) error {
	rawConfig, err := ioutil.ReadFile(conf)
	if err != nil {
		return err
	}
	viper.SetConfigType("yaml")
	if err = viper.ReadConfig(bytes.NewBuffer(rawConfig)); err != nil {
		return err
	}

	if err := viper.Unmarshal(&val); err != nil {
		logrus.Panic(err)
		return err
	}
	return nil
}
