package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config Структура конфигурации;
// Содержит все конфигурационные данные о сервисе;
// автоподгружается при изменении исходного файла
type Config struct {
	ServiceHost string
	ServicePort int
	Minio       `toml:"minio"`
	JWT         `toml:"JWT"`
	Redis       `toml:"redis"`
	Postgresql  `toml:"postgresql"`
}
type Minio struct {
	User     string `toml:"user"`
	Pass     string `toml:"pass"`
	Endpoint string `toml:"endpoint"`
}

type JWT struct {
	Key           string `toml:"Key"`
	SigningMethod jwt.SigningMethod
	ExpiresIn     time.Duration `toml:"Expires-in"`
}

type Redis struct {
	Redis_host     string `toml:"REDIS_HOST"`
	Redis_password string `toml:"REDIS_PASSWORD"`
	Redis_port     int    `toml:"REDIS_PORT"`
	Redis_user     string `toml:"REDIS_USER"`
	DialTimeout    time.Duration
	ReadTimeout    time.Duration
}
type Postgresql struct {
	DB_Host string `toml:"DB_HOST"`
	DB_Port string `toml:"DB_PORT"`
	DB_Name string `toml:"DB_NAME"`
	DB_User string `toml:"DB_USER"`
	DB_Pass string `toml:"DB_PASS"`
}

// NewConfig Создаёт новый объект конфигурации, загружая данные из файла конфигурации
func NewConfig(log *log.Logger) (*Config, error) {
	var err error

	configName := "config"
	_ = godotenv.Load()
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	cfg.JWT.SigningMethod = jwt.SigningMethodHS256 //can change inline
	cfg.JWT.ExpiresIn = 12 * time.Hour

	log.Info("config parsed")
	log.Info(cfg.ServiceHost)
	log.Info(cfg.ServicePort)
	log.Info(cfg.Minio)
	log.Info(cfg.Redis)
	log.Info(cfg.Postgresql)
	log.Info(cfg.JWT)

	return cfg, nil
}
