package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	Host     string `env:"DB_HOST"`
	Port     uint16 `env:"DB_PORT"`
	Database string `env:"DB_NAME"`
	SslMode  string `env:"DB_SSLMODE"`
}

type JWTConfig struct {
	Secret string `env:"JWT_SECRET"`
}

type AWSConfig struct {
	Region     string `env:"AWS_REGION"`
	AccessKey  string `env:"AWS_ACCESS_KEY_ID"`
	SecretKey  string `env:"AWS_SECRET_ACCESS_KEY"`
	BucketName string `env:"AWS_S3_BUCKET_NAME"`
}

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
}

var database *gorm.DB
var e error

func Get() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load(".env")
	if err != nil {
		return cfg, err
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func DatabaseInit() {
	cfg, cfgErr := Get()
	if cfgErr != nil {
		log.Fatalf("Failed to load configuration: %v", cfgErr)
	}

	dbCfg := cfg.Database

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=require",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Database,
	)
	fmt.Println(connStr)

	database, e = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}
