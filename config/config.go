package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Db       Db
		Jwt      Jwt
		Kafka    Kafka
		Grpc     Grpc
		Paginate Paginate
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}

	Db struct {
		Host     string
		Port     int64
		User     string
		Password string
		Name     string
		SSLMode  string
	}

	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSecretKey     string
		AccessDuration   int64
		RefreshDuration  int64
		ApiDuration      int64
	}

	Kafka struct {
		Url    string
		ApiKey string
		Secret string
	}

	Grpc struct {
		AuthUrl string
		UserUrl string
	}

	Paginate struct {
		UserNextPageBaseUrl string
		AuthNextPageBaseUrl string
	}
)

func LoadConfig(path string) *Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			Stage: os.Getenv("APP_STAGE"),
		},
		Db: Db{
			Host: os.Getenv("DB_HOST"),
			Port: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
				if err != nil {
					log.Fatal("Error loading database port failed")
				}
				return result
			}(),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
		},
		Jwt: Jwt{
			AccessSecretKey:  os.Getenv("JWT_ACCESS_SECRET_KEY"),
			RefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
			ApiSecretKey:     os.Getenv("JWT_API_SECRET_KEY"),
			AccessDuration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error loading access dulation failed")
				}
				return result
			}(),
			RefreshDuration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error loading access dulation failed")
				}
				return result
			}(),
		},
		Kafka: Kafka{
			Url:    os.Getenv("KAFKA_URL"),
			ApiKey: os.Getenv("KAFKA_API_KEY"),
			Secret: os.Getenv("KAFKA_API_SECRET"),
		},
		Grpc: Grpc{
			AuthUrl: os.Getenv("GRPC_AUTH_URL"),
			UserUrl: os.Getenv("GRPC_USER_URL"),
		},
		Paginate: Paginate{
			UserNextPageBaseUrl: os.Getenv("PAGINATE_USER_NEXT_PAGE_BASED_URL"),
			AuthNextPageBaseUrl: os.Getenv("PAGINATE_AUTH_NEXT_PAGE_BASED_URL"),
		},
	}
}
