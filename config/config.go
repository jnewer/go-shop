package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}

	JwtSettings struct {
		SecretKey string
	}

	configReader struct {
		configFile string
		v          *viper.Viper
	}

	Configuration struct {
		DatabaseSettings
		JwtSettings
	}
)

var cfgReader *configReader

func newConfigReader(configFile string) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}
func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	newConfigReader(configFile)
	if err = cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("配置文件读取失败：%s", err)
		return nil, err
	}

	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("解析配置文件到结构体失败：%s", err)
		return nil, err
	}

	return configuration, err

}
