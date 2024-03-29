package configs

import (
	"errors"
	"flag"
	"os"
)

const (
	NumbWorkers = 10
	WorkerBuff  = 100
)

type appConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080/"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"storage.dat"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
	SecretKey       []byte
	NumbWorkers     int `env:"NUMBER_OF_WORKERS"`
	WorkerBuff      int `env:"WORKERS_BUFFER"`
}

func NewConfig() (*appConfig, error) {
	serverAddress := getServerAddress()
	baseURL := getBaseURL()
	fileStoragePath := getFileStoragePath()
	secretKey := getSecretKey()
	databaseDSN := getDatabaseDSN()
	flag.Parse()

	if serverAddress == nil {
		return nil, errors.New("server address not specified")
	}

	if baseURL == nil {
		return nil, errors.New("base url not specified")
	}

	if fileStoragePath == nil {
		return nil, errors.New("file storage path not specified")
	}

	if databaseDSN == nil {
		return nil, errors.New("database dsn not specified")
	}

	if secretKey == nil {
		return nil, errors.New("secret key not specified")
	}

	return &appConfig{
		ServerAddress:   *serverAddress,
		BaseURL:         *baseURL,
		FileStoragePath: *fileStoragePath,
		DatabaseDSN:     *databaseDSN,
		SecretKey:       []byte(*secretKey),
		NumbWorkers:     NumbWorkers,
		WorkerBuff:      WorkerBuff,
	}, nil
}

func getServerAddress() *string {
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	return flag.String("a", address, "server address")
}

func getBaseURL() *string {
	url := os.Getenv("BASE_URL")
	if url == "" {
		url = "http://localhost:8080"
	}

	return flag.String("b", url, "base url")
}

func getFileStoragePath() *string {
	path := os.Getenv("FILE_STORAGE_PATH")

	return flag.String("f", path, "file storage path")
}

func getDatabaseDSN() *string {
	databaseDSN := os.Getenv("DATABASE_DSN")

	return flag.String("d", databaseDSN, "database")
}

func getSecretKey() *string {
	url := os.Getenv("SECRET_KEY")
	if url == "" {
		url = "my-secret-key"
	}

	return flag.String("s", url, "secret key")
}
