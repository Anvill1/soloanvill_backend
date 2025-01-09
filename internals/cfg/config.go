package cfg

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cfg struct {
	Port     string
	Database struct {
		Host         string
		Port         int
		Name         string
		User         string
		Password     string
		DefaultTable string
	}
	Jenkins struct {
		Host  string
		Login string
		Token string
	}
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()

	v.SetConfigFile("/etc/soloanvill/config.yml")

	err := v.ReadInConfig()
	if err != nil {
		log.Info(err)
		err = nil
	}

	v.SetDefault("PORT", "8080")
	v.SetDefault("Database.User", "postgres")
	v.SetDefault("Database.Password", "password")
	v.SetDefault("Database.Host", "127.0.0.1")
	v.SetDefault("Database.Port", "5432")
	v.SetDefault("Database.Name", "soloanvill")
	v.SetDefault("Database.DefaultTable", "users")
	v.SetDefault("Jenkins.Host", "127.0.0.1")
	v.SetDefault("Jenkins.Login", "jenkinslogin")
	v.SetDefault("Jenkins.Token", "jenkinstoken")

	var cfg Cfg

	err = v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}

	return cfg
}

func (cfg *Cfg) GetDBString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%v/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
}
