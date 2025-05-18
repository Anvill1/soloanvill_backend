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
	v.SetEnvPrefix("SOLOANVILL")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Warn(err)
	}

	v.SetDefault("PORT", "8080")
	v.SetDefault("Database.User", v.GetString("DATABASE_USER"))
	v.SetDefault("Database.Password", v.GetString("DATABASE_PASSWORD"))
	v.SetDefault("Database.Host", v.GetString("DATABASE_HOST"))
	v.SetDefault("Database.Port", v.GetString("DATABASE_PORT"))
	v.SetDefault("Database.Name", v.GetString("DATABASE_NAME"))
	v.SetDefault("Database.DefaultTable", "users")
	v.SetDefault("Jenkins.Host", v.GetString("JENKINS_HOST"))
	v.SetDefault("Jenkins.Login", v.GetString("JENKINS_LOGIN"))
	v.SetDefault("Jenkins.Token", v.GetString("JENKINS_TOKEN"))

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
