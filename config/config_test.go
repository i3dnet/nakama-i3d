package config

import (
	"context"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func runtimeConfig(env map[string]string) (*Config, *runtime.Error) {
	return NewConfigFromRuntime(context.WithValue(context.Background(), runtime.RUNTIME_CTX_ENV, env))
}
func staticEnv() map[string]string {
	return map[string]string{"I3D_APPLICATION_ID": "123", "I3D_ACCESS_TOKEN": "test-only"}
}
func isolateProcessConfig(t *testing.T, env map[string]string) {
	t.Helper()
	t.Setenv("PROJECT_ROOT", t.TempDir())
	for _, entry := range os.Environ() {
		key, _, _ := strings.Cut(entry, "=")
		if strings.HasPrefix(key, "I3D_") {
			t.Setenv(key, "")
			require.NoError(t, os.Unsetenv(key))
		}
	}
	for key, value := range env {
		t.Setenv(key, value)
	}
}
func TestRuntimeAndProcessConfigAgreeOnRetryOverrides(t *testing.T) {
	env := staticEnv()
	env["I3D_BASE_URL"] = "https://provider.example"
	env["I3D_RETRY_ATTEMPTS"] = "5"
	env["I3D_RETRY_DELAY"] = "20ms"
	env["I3D_RETRY_MAX_DELAY"] = "100ms"
	isolateProcessConfig(t, env)
	fromRuntime, err := runtimeConfig(env)
	require.Nil(t, err)
	fromProcess, err := NewConfig()
	require.Nil(t, err)
	require.Equal(t, 5, fromRuntime.Attempts)
	require.Equal(t, 20*time.Millisecond, fromRuntime.Delay)
	require.Equal(t, 100*time.Millisecond, fromRuntime.MaxDelay)
	require.Equal(t, fromRuntime, fromProcess)
}
func TestCanonicalBaseURLWinsOverDocumentedAlias(t *testing.T) {
	env := staticEnv()
	env["I3D_API_URL"] = "https://alias.example"
	cfg, err := runtimeConfig(env)
	require.Nil(t, err)
	require.Equal(t, "https://alias.example", cfg.BaseUrl)
	env["I3D_BASE_URL"] = "https://canonical.example"
	cfg, err = runtimeConfig(env)
	require.Nil(t, err)
	require.Equal(t, "https://canonical.example", cfg.BaseUrl)
}
func TestInvalidRuntimeConfigurationFails(t *testing.T) {
	for _, test := range []struct{ key, value string }{
		{"I3D_USE_BEARER_AUTH", "maybe"}, {"I3D_RETRY_ATTEMPTS", "0"}, {"I3D_RETRY_ATTEMPTS", "abc"}, {"I3D_RETRY_DELAY", "later"}, {"I3D_RETRY_DELAY", "-1s"}, {"I3D_RETRY_MAX_DELAY", "1ms"}, {"I3D_BASE_URL", "not-a-url"}, {"I3D_APPLICATION_ID", " "},
	} {
		t.Run(test.key+"/"+test.value, func(t *testing.T) {
			env := staticEnv()
			env[test.key] = test.value
			_, err := runtimeConfig(env)
			require.NotNil(t, err)
		})
	}
}
func TestSettingJSONUsesProjectRootAndEnvironmentOverride(t *testing.T) {
	isolateProcessConfig(t, nil)
	dir := os.Getenv("PROJECT_ROOT")
	require.NoError(t, os.WriteFile(filepath.Join(dir, "setting.json"), []byte(`{"oneApi":{"applicationId":"123","token":"file-test-only","baseUrl":"https://file.example"}}`), 0600))
	cfg, err := NewConfig()
	require.Nil(t, err)
	require.Equal(t, "https://file.example", cfg.BaseUrl)
	t.Setenv("I3D_BASE_URL", "https://override.example")
	cfg, err = NewConfig()
	require.Nil(t, err)
	require.Equal(t, "https://override.example", cfg.BaseUrl)
	t.Setenv("I3D_USE_BEARER_AUTH", "invalid")
	_, err = NewConfig()
	require.NotNil(t, err, "invalid environment must not fall back to file credentials")
}
func TestOAuthRuntimeAndProcessConfigAgree(t *testing.T) {
	env := map[string]string{"I3D_APPLICATION_ID": "123", "I3D_USE_BEARER_AUTH": "true", "I3D_CLIENT_ID": "client", "I3D_CLIENT_SECRET": "test-only", "I3D_AUDIENCE": "audience", "I3D_AUTHENTICATION_URL": "https://auth.example/token"}
	isolateProcessConfig(t, env)
	fromRuntime, err := runtimeConfig(env)
	require.Nil(t, err)
	fromProcess, err := NewConfig()
	require.Nil(t, err)
	require.Equal(t, fromRuntime, fromProcess)
}

func TestTimeoutAndReconciliationSettingsAgreeAcrossSources(t *testing.T) {
	env := staticEnv()
	for key, value := range map[string]string{"I3D_FLEET_ID": "456", "I3D_ALLOCATION_TIMEOUT": "2m", "I3D_PROVIDER_TIMEOUT": "45s", "I3D_RECONCILE_INTERVAL": "10s", "I3D_RECONCILE_TIMEOUT": "5s", "I3D_RECONCILE_GRACE_PERIOD": "20s"} {
		env[key] = value
	}
	isolateProcessConfig(t, env)
	a, err := runtimeConfig(env)
	require.Nil(t, err)
	b, err := NewConfig()
	require.Nil(t, err)
	require.Equal(t, a, b)
	require.Equal(t, "456", a.FleetId)
	require.Equal(t, 45*time.Second, a.ProviderTimeout)
	require.Equal(t, 10*time.Second, a.ReconcileInterval)
	require.Equal(t, 5*time.Second, a.ReconcileTimeout)
	require.Equal(t, 20*time.Second, a.ReconcileGracePeriod)
}
func TestInvalidBoundsAndDisabledReconciliation(t *testing.T) {
	for _, setting := range []struct{ key, value string }{
		{"I3D_RETRY_ATTEMPTS", "11"}, {"I3D_ALLOCATION_TIMEOUT", "0s"}, {"I3D_PROVIDER_TIMEOUT", "0s"}, {"I3D_RECONCILE_INTERVAL", "-1s"}, {"I3D_RECONCILE_TIMEOUT", "0s"}, {"I3D_RECONCILE_GRACE_PERIOD", "-1s"},
	} {
		env := staticEnv()
		env[setting.key] = setting.value
		_, err := runtimeConfig(env)
		require.NotNil(t, err, setting.key)
	}
	env := staticEnv()
	env["I3D_RECONCILE_INTERVAL"] = "0s"
	cfg, err := runtimeConfig(env)
	require.Nil(t, err)
	require.Zero(t, cfg.ReconcileInterval)
}
func TestDotEnvAndWorkingDirectoryResolution(t *testing.T) {
	isolateProcessConfig(t, nil)
	dir := t.TempDir()
	t.Setenv("PROJECT_ROOT", "")
	t.Chdir(dir)
	require.NoError(t, os.WriteFile(filepath.Join(dir, ".env"), []byte("I3D_APPLICATION_ID=123\nI3D_ACCESS_TOKEN=test-only\nI3D_BASE_URL=https://file.example\n"), 0600))
	cfg, err := NewConfig()
	require.Nil(t, err)
	require.Equal(t, "https://file.example", cfg.BaseUrl)
	t.Setenv("I3D_BASE_URL", "https://process.example")
	cfg, err = NewConfig()
	require.Nil(t, err)
	require.Equal(t, "https://process.example", cfg.BaseUrl)
}
