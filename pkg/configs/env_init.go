package configs

import (
	"github.com/popeskul/houser/pkg/env"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func EnvConfigs() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}

	err = env.InitEnv()
	if err != nil {
		logrus.Error(err)
	}
}
