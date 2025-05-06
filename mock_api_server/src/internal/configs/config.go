package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	App
	HTTP
}

type App struct {
	Name    string
	Version string
}

type HTTP struct {
	Port string `env-required:"true" env:"HTTP_PORT"`
}

func NewConfig() *Config {
	cfg := &Config{
		App: App{
			Name:    "mock-api-server",
			Version: "1.0.0",
		},
		HTTP: HTTP{
			Port: "8080",
		},
	}
	cwd := projectRoot()
	envFilePath := cwd + ".env"

	err := readEnv(envFilePath, cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

func readEnv(envFilePath string, cfg *Config) error {
	envFileExists := checkFileExists(envFilePath)

	if envFileExists {
		err := cleanenv.ReadConfig(envFilePath, cfg)
		if err != nil {
			return fmt.Errorf("config error: %w", err)
		}
	} else {
		err := cleanenv.ReadEnv(cfg)
		if err != nil {

			if _, statErr := os.Stat(envFilePath + ".example"); statErr == nil {
				return fmt.Errorf("missing environmentvariables: %w\n\nprovide all required environment variables or rename and update .env.example to .env for convinience", err)
			}

			return err
		}
	}
	return nil
}

func checkFileExists(fileName string) bool {
	envFileExists := false
	if _, err := os.Stat(fileName); err == nil {
		envFileExists = true
	}
	return envFileExists
}

func projectRoot() string {

	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot != "" {
		fmt.Println("projectRoot: ", projectRoot)
		return projectRoot + "/"
	}

	_, b, _, _ := runtime.Caller(0)
	projectRoot = filepath.Dir(b)

	fmt.Println("projectRoot: ", projectRoot)

	return projectRoot + "/../"
}
