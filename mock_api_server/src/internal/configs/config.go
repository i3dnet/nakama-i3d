package configs

import (
	"encoding/json"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	App
	HTTP `json:"http" env:"HTTP"`
}

type App struct {
	Name    string
	Version string
}

type HTTP struct {
	Port string `env-required:"true" json:"port" env:"HTTP_PORT"`
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

func readJson(jsonFilePath string, cfg *Config) error {

	// check if file exists
	if _, err := os.Stat(jsonFilePath); err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	file, err := os.Open(jsonFilePath)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file) // In Go 1.16+, you can use os.ReadFile(filename) directly
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
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
