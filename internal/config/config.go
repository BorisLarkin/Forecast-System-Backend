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
	Minio       `yaml:"minio"`
	JWT
	Redis
	Postgresql
}
type Minio struct {
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Endpoint string `yaml:"endpoint"`
}

type JWT struct {
	Token         string            `yaml:"token"`
	SigningMethod jwt.SigningMethod `yaml:"signing-method"`
	ExpiresIn     time.Duration     `yaml:"expires-in"`
}

type Redis struct {
	Redis_host     string
	Redis_password string
	Redis_port     int
	Redis_user     string
	DialTimeout    time.Duration
	ReadTimeout    time.Duration
}
type Postgresql struct {
	DB_Host string
	DB_Port string
	DB_Name string
	DB_User string
	DB_Pass string
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

	log.Info("config parsed")
	log.Info(cfg.ServiceHost)
	log.Info(cfg.ServicePort)
	log.Info(cfg.Minio)
	log.Info(cfg.Redis)
	log.Info(cfg.Postgresql)

	return cfg, nil
}
