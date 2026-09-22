package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/joho/godotenv"
)

type Config struct {
	App
	OneApi               `json:"oneApi"`
	Retry                `json:"retry"`
	AllocationTimeout    time.Duration `json:"allocationTimeout" env:"I3D_ALLOCATION_TIMEOUT"`
	ProviderTimeout      time.Duration `json:"providerTimeout" env:"I3D_PROVIDER_TIMEOUT"`
	ReconcileInterval    time.Duration `json:"reconcileInterval" env:"I3D_RECONCILE_INTERVAL"`
	ReconcileTimeout     time.Duration `json:"reconcileTimeout" env:"I3D_RECONCILE_TIMEOUT"`
	ReconcileGracePeriod time.Duration `json:"reconcileGracePeriod" env:"I3D_RECONCILE_GRACE_PERIOD"`
}
type App struct {
	Name    string
	Version string
}
type OneApi struct {
	ApplicationId     string `json:"applicationId" env:"I3D_APPLICATION_ID"`
	FleetId           string `json:"fleetId" env:"I3D_FLEET_ID"`
	BaseUrl           string `json:"baseUrl" env:"I3D_BASE_URL"`
	Token             string `json:"token" env:"I3D_ACCESS_TOKEN"`
	UseBearerAuth     bool   `json:"useBearerAuth" env:"I3D_USE_BEARER_AUTH"`
	ClientId          string `json:"clientId" env:"I3D_CLIENT_ID"`
	ClientSecret      string `json:"clientSecret" env:"I3D_CLIENT_SECRET"`
	Audience          string `json:"audience" env:"I3D_AUDIENCE"`
	AuthenticationUrl string `json:"authenticationUrl" env:"I3D_AUTHENTICATION_URL"`
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
		App:               App{Name: "Nakama one plugin", Version: "1.0.0"},
		OneApi:            OneApi{BaseUrl: "https://api.i3d.net"},
		Retry:             Retry{Attempts: 3, Delay: 1500 * time.Millisecond, MaxDelay: 7500 * time.Millisecond},
		AllocationTimeout: 120 * time.Second, ProviderTimeout: 90 * time.Second,
		ReconcileInterval: time.Minute, ReconcileTimeout: 30 * time.Second, ReconcileGracePeriod: 2 * time.Minute,
	}
}
func configError(err error) *runtime.Error { return runtime.NewError("configuration: "+err.Error(), 3) }

// NewConfigFromRuntime uses only Nakama's runtime environment. Invalid runtime
// input must not silently select unrelated process/file credentials.
func NewConfigFromRuntime(ctx context.Context) (*Config, *runtime.Error) {
	env, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		return nil, configError(fmt.Errorf("runtime environment is unavailable"))
	}
	cfg := defaultConfig()
	if err := applyEnv(cfg, env); err != nil {
		return nil, configError(err)
	}
	if err := validate(cfg); err != nil {
		return nil, configError(err)
	}
	return cfg, nil
}

// NewConfig reads defaults, optional setting.json, optional .env, then process
// environment overrides. Paths are relative to PROJECT_ROOT or the working directory.
func NewConfig() (*Config, *runtime.Error) {
	root := os.Getenv("PROJECT_ROOT")
	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			return nil, configError(err)
		}
	}
	cfg := defaultConfig()
	data, err := os.ReadFile(filepath.Join(root, "setting.json"))
	if err == nil {
		if err = json.Unmarshal(data, cfg); err != nil {
			return nil, configError(fmt.Errorf("invalid setting.json: %w", err))
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, configError(err)
	}
	env := map[string]string{}
	envFile := filepath.Join(root, ".env")
	if _, err = os.Stat(envFile); err == nil {
		env, err = godotenv.Read(envFile)
		if err != nil {
			return nil, configError(fmt.Errorf("invalid .env file"))
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, configError(err)
	}
	for _, entry := range os.Environ() {
		key, value, _ := strings.Cut(entry, "=")
		if strings.HasPrefix(key, "I3D_") {
			env[key] = value
		}
	}
	if err = applyEnv(cfg, env); err != nil {
		return nil, configError(err)
	}
	if err = validate(cfg); err != nil {
		return nil, configError(err)
	}
	return cfg, nil
}
func applyEnv(cfg *Config, env map[string]string) error {
	if _, canonical := env[ENV_BASE_URL]; !canonical {
		if alias, ok := env["I3D_API_URL"]; ok {
			cfg.BaseUrl = alias
		}
	}
	for key, target := range map[string]*string{
		ENV_APPLICATION_ID: &cfg.ApplicationId, "I3D_FLEET_ID": &cfg.FleetId, ENV_BASE_URL: &cfg.BaseUrl,
		ENV_ACCESS_TOKEN: &cfg.Token, ENV_CLIENT_ID: &cfg.ClientId, ENV_CLIENT_SECRET: &cfg.ClientSecret,
		ENV_AUDIENCE: &cfg.Audience, ENV_AUTHENTICATION_URL: &cfg.AuthenticationUrl,
	} {
		if value, ok := env[key]; ok {
			*target = value
		}
	}
	if value, ok := env[ENV_BEARER_AUTH]; ok {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("%s must be a boolean", ENV_BEARER_AUTH)
		}
		cfg.UseBearerAuth = parsed
	}
	if value, ok := env["I3D_RETRY_ATTEMPTS"]; ok {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("I3D_RETRY_ATTEMPTS must be an integer")
		}
		cfg.Attempts = parsed
	}
	for key, target := range map[string]*time.Duration{
		"I3D_RETRY_DELAY": &cfg.Delay, "I3D_RETRY_MAX_DELAY": &cfg.MaxDelay,
		"I3D_ALLOCATION_TIMEOUT": &cfg.AllocationTimeout, "I3D_PROVIDER_TIMEOUT": &cfg.ProviderTimeout,
		"I3D_RECONCILE_INTERVAL": &cfg.ReconcileInterval, "I3D_RECONCILE_TIMEOUT": &cfg.ReconcileTimeout,
		"I3D_RECONCILE_GRACE_PERIOD": &cfg.ReconcileGracePeriod,
	} {
		if value, ok := env[key]; ok {
			duration, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("%s must be a Go duration", key)
			}
			*target = duration
		}
	}
	cfg.BaseUrl = strings.TrimRight(cfg.BaseUrl, "/")
	return nil
}
func validate(cfg *Config) error {
	var problems []error
	required := map[string]string{ENV_APPLICATION_ID: cfg.ApplicationId}
	if cfg.UseBearerAuth {
		required[ENV_CLIENT_ID] = cfg.ClientId
		required[ENV_CLIENT_SECRET] = cfg.ClientSecret
		required[ENV_AUDIENCE] = cfg.Audience
		required[ENV_AUTHENTICATION_URL] = cfg.AuthenticationUrl
	} else {
		required[ENV_ACCESS_TOKEN] = cfg.Token
	}
	for key, value := range required {
		if strings.TrimSpace(value) == "" {
			problems = append(problems, fmt.Errorf("%s is required", key))
		}
	}
	validateURL := func(key, value string) {
		parsed, err := url.Parse(value)
		if err != nil || parsed.Host == "" || (parsed.Scheme != "https" && parsed.Scheme != "http") || parsed.User != nil {
			problems = append(problems, fmt.Errorf("%s must be an HTTP(S) URL without embedded credentials", key))
		}
	}
	validateURL(ENV_BASE_URL, cfg.BaseUrl)
	if cfg.UseBearerAuth {
		validateURL(ENV_AUTHENTICATION_URL, cfg.AuthenticationUrl)
	}
	if cfg.Attempts < 1 || cfg.Attempts > 10 {
		problems = append(problems, fmt.Errorf("I3D_RETRY_ATTEMPTS must be between 1 and 10"))
	}
	if cfg.Delay < 0 || cfg.MaxDelay < cfg.Delay {
		problems = append(problems, fmt.Errorf("retry delays must be nonnegative and max delay must be at least the initial delay"))
	}
	if cfg.AllocationTimeout <= 0 || cfg.ProviderTimeout <= 0 {
		problems = append(problems, fmt.Errorf("allocation and provider timeouts must be positive"))
	}
	if cfg.ReconcileInterval < 0 || cfg.ReconcileTimeout <= 0 || cfg.ReconcileGracePeriod < 0 {
		problems = append(problems, fmt.Errorf("reconciliation interval/grace must be nonnegative and timeout must be positive"))
	}
	return errors.Join(problems...)
}
