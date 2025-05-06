package fleetmanager_config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/ilyakaznacheev/cleanenv"
	"io"
	"os"
	"path/filepath"
	runtime2 "runtime"
	"time"
)

type Config struct {
	App
	OneApi `json:"oneApi" env-required:"true"`
	Retry  `json:"retry" env-required:"true"`
}

type App struct {
	Name    string
	Version string
}

type OneApi struct {
	ApplicationId     string `json:"applicationId" env-required:"true" env:"I3D_APPLICATION_ID"`
	BaseUrl           string `json:"baseUrl" env-required:"true" env:"I3D_BASE_URL"`
	Token             string `json:"token" env:"I3D_ACCESS_TOKEN"`
	UseBearerAuth     bool   `json:"useBearerAuth" env:"I3D_USE_BEARER_AUTH"`
	ClientId          string `json:"clientId"  env:"I3D_CLIENT_ID"`
	ClientSecret      string `json:"clientSecret"  env:"I3D_CLIENT_SECRET"`
	Audience          string `json:"audience"  env:"I3D_AUDIENCE"`
	AuthenticationUrl string `json:"authenticationUrl"  env:"I3D_AUTHENTICATION_URL"`
}

type Retry struct {
	Attempts int           `json:"attempts" env:"I3D_RETRY_ATTEMPTS"`
	Delay    time.Duration `json:"delay" env:"I3D_RETRY_DELAY"`
	MaxDelay time.Duration `json:"maxDelay" env:"I3D_RETRY_MAX_DELAY"`
}

const (
	ENV_APPLICATION_ID     = "I3D_APPLICATION_ID"
	ENV_BASE_URL           = "I3D_BASE_URL"
	ENV_ACCESS_TOKEN       = "I3D_ACCESS_TOKEN"
	ENV_BEARER_AUTH        = "I3D_USE_BEARER_AUTH"
	ENV_CLIENT_ID          = "I3D_CLIENT_ID"
	ENV_CLIENT_SECRET      = "I3D_CLIENT_SECRET"
	ENV_AUDIENCE           = "I3D_AUDIENCE"
	ENV_AUTHENTICATION_URL = "I3D_AUTHENTICATION_URL"
)

func defaultConfig() *Config {
	return &Config{
		App: App{
			Name:    "Nakama one plugin",
			Version: "1.0.0",
		},
		OneApi: OneApi{
			BaseUrl:       "https://api.i3d.net",
			UseBearerAuth: false,
		},
		Retry: Retry{
			Attempts: 3,
			Delay:    1500 * time.Millisecond,
			MaxDelay: 7500 * time.Millisecond,
		},
	}
}

func NewConfigFromRuntime(ctx context.Context) (*Config, *runtime.Error) {
	cfg := defaultConfig()
	envValue := ctx.Value(runtime.RUNTIME_CTX_ENV)
	env, ok := envValue.(map[string]string)
	if !ok {
		return nil, runtime.NewError(fmt.Sprintf("unable to cast '%v' to a map", runtime.RUNTIME_CTX_ENV), 3)
	}

	applicationId, exists := env[ENV_APPLICATION_ID]
	if !exists {
		return nil, runtime.NewError(fmt.Sprintf("unable to find '%v'", ENV_APPLICATION_ID), 3)
	}

	cfg.OneApi.ApplicationId = applicationId

	baseUrl, exists := env[ENV_BASE_URL]
	if exists {
		cfg.OneApi.BaseUrl = baseUrl
	}

	useBearerAuth, exists := env[ENV_BEARER_AUTH]
	if !exists {
		useBearerAuth = "false"
		cfg.OneApi.UseBearerAuth = false
	}

	if useBearerAuth == "true" {
		cfg.OneApi.UseBearerAuth = true
		clientId, exists := env[ENV_CLIENT_ID]
		if !exists {
			return nil, runtime.NewError(fmt.Sprintf("unable to find '%v'", ENV_CLIENT_ID), 3)
		}
		cfg.OneApi.ClientId = clientId

		clientSecret, exists := env[ENV_CLIENT_SECRET]
		if !exists {
			return nil, runtime.NewError(fmt.Sprintf("unable to find '%v'", ENV_CLIENT_SECRET), 3)
		}
		cfg.OneApi.ClientSecret = clientSecret

		audience, exists := env[ENV_AUDIENCE]
		if !exists {
			return nil, runtime.NewError(fmt.Sprintf("unable to find '%v'", ENV_AUDIENCE), 3)
		}
		cfg.OneApi.Audience = audience

		authenticationUrl, exists := env[ENV_AUTHENTICATION_URL]
		if !exists {
			return nil, runtime.NewError(fmt.Sprintf("unable to find '%v'", ENV_AUTHENTICATION_URL), 3)
		}

		cfg.OneApi.AuthenticationUrl = authenticationUrl
	} else {
		token, exists := env[ENV_ACCESS_TOKEN]
		if !exists {
			return nil, runtime.NewError(fmt.Sprintf("unable to find '%v'", ENV_ACCESS_TOKEN), 3)
		}

		cfg.OneApi.Token = token

	}

	if err := validate(cfg); err != nil {
		return nil, runtime.NewError(fmt.Sprintf("config validation error: %v", err), 3)
	}

	return cfg, nil
}

func NewConfig() (*Config, *runtime.Error) {
	cfg := defaultConfig()
	cwd := projectRoot()
	envFilePath := cwd + ".env"

	err := readEnv(envFilePath, cfg)

	if err != nil {
		fmt.Println("Environment variables not found, trying to read config from json")
		err = readJson(envFilePath+"setting.json", cfg)

		if err != nil {
			return nil, runtime.NewError(fmt.Sprintf("Environment variables were not found: %v", err), 3)
		}
	}

	if err := validate(cfg); err != nil {
		fmt.Println("Config validation error: ", err)
		return nil, runtime.NewError(fmt.Sprintf("config validation error: %v", err), 3)
	}

	return cfg, nil
}

func validate(cfg *Config) error {
	errs := make([]error, 0)
	if cfg.OneApi.ApplicationId == "" {
		errs = append(errs, fmt.Errorf("missing application id"))
	}
	if cfg.OneApi.BaseUrl == "" {
		errs = append(errs, fmt.Errorf("missing base url"))
	}
	if cfg.OneApi.UseBearerAuth {
		if cfg.OneApi.ClientId == "" {
			errs = append(errs, fmt.Errorf("missing client id"))
		}
		if cfg.OneApi.ClientSecret == "" {
			errs = append(errs, fmt.Errorf("missing client secret"))
		}

		if cfg.OneApi.Audience == "" {
			errs = append(errs, fmt.Errorf("missing audience"))
		}

		if cfg.OneApi.AuthenticationUrl == "" {
			errs = append(errs, fmt.Errorf("missing authentication url"))
		}

	} else {
		if cfg.OneApi.Token == "" {
			errs = append(errs, fmt.Errorf("missing token"))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
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

	bytes, err := io.ReadAll(file)
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

	_, b, _, _ := runtime2.Caller(0)
	projectRoot = filepath.Dir(b)

	fmt.Println("projectRoot: ", projectRoot)

	return projectRoot + "/../"
}
