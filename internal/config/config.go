package config

import (
	"ggate/pkg/config"

	"github.com/sirupsen/logrus"
)

type ProxyConfig struct {
	Addr               string `mapstructure:"addr"`
	HttpsAddr          string `mapstructure:"https_addr"`
	FilterPluginsDir   string `mapstructure:"filter_plugin_dir"`
	PredicatePluginDir string `mapstructure:"predicate_plugin_dir"`
	RoutesFile         string `mapstructure:"routes_file"`
	ServicesFile       string `mapstructure:"services_file"`

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
	err := config.LoadYaml(configFile, &cfg)
	if err != nil {
		logrus.Panic("LoadProxyConfig fail")
		return cfg, err
	}
	return cfg, nil
}
