package cfg

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Cfg struct {
	Port         string
	DBUser       string
	DBPass       string
	DBHOST       string
	DBPORT       string
	DBName       string
	DefaultTable string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("SOLOAVNILL")
	v.SetDefault("PORT", "8080")
	v.SetDefault("DBUSER", "soloanvill")
	v.SetDefault("DBPASS", "password")
	v.SetDefault("DBHOST", "127.0.0.1")
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("DBNAME", "soloanvill")
	v.SetDefault("DefaultTable", "users")

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPass, cfg.DBHOST, cfg.DBPORT, cfg.DBName)
}
