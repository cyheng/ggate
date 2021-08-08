package config

import (
	"bytes"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ProxyConfig struct {
	Addr               string `mapstructure:"addr"`
	HttpsAddr          string `mapstructure:"https_addr"`
	FilterPluginsDir   string `mapstructure:"filter_plugin_dir"`
	PredicatePluginDir string `mapstructure:"predicate_plugin_dir"`
	ConfigReaderType   string `mapstructure:"config_reader_type"`
	ConfigReaderArg    string `mapstructure:"config_reader_arg"`
	RoutesReaderType   string `mapstructure:"routes_reader_type"`
	RoutesReaderArg    string `mapstructure:"routes_reader_arg"`
	ServicesReaderType string `mapstructure:"services_reader_type"`
	ServicesReaderArg  string `mapstructure:"services_reader_arg"`

	MaxIdleConns        int    `mapstructure:"max_idle_conns"`
	MaxIdleConnsPerHost int    `mapstructure:"max_idle_conns_per_host"`
	AccessLogFile       string `mapstructure:"access_log_file"`
	EnableHttps         bool   `mapstructure:"enable_https"`
	RedirectHttps       bool   `mapstructure:"redirect_https"`

	AdminDomain string `mapstructure:"admin_domain"`
	AdminPort   string `mapstructure:"admin_port"`
	AdminName   string `mapstructure:"admin_name"`
	AdminPass   string `mapstructure:"admin_pass"`
}

func LoadProxyConfig(configFile string) (ProxyConfig, error) {
	cfg := ProxyConfig{}
	rawConfig, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.Panic("read config file error", err)
		return cfg, err
	}

	viper.SetConfigType("yaml")
	if err = viper.ReadConfig(bytes.NewBuffer(rawConfig)); err != nil {
		logrus.Panic("read config error", err)
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		logrus.Panic(err)
		return cfg, err
	}

	return cfg, nil
}
