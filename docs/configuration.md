# Configuration

NewConfigFromRuntime reads only Nakama's runtime.env map. NewConfig reads defaults, optional setting.json, optional .env and then process environment overrides, in that order. Files resolve from PROJECT_ROOT or the current working directory. Invalid input returns an error; it does not trigger a credential fallback. The example chooses runtime configuration whenever Nakama supplies that context value.

I3D_BASE_URL is canonical; the documented I3D_API_URL is accepted only when BASE_URL is absent. The default is https://api.i3d.net. I3D_APPLICATION_ID is required. I3D_FLEET_ID optionally scopes allocations/listing/reconciliation to one fleet.

Static-token mode uses I3D_ACCESS_TOKEN. Set I3D_USE_BEARER_AUTH=true for OAuth and supply I3D_CLIENT_ID, I3D_CLIENT_SECRET, I3D_AUDIENCE and the complete I3D_AUTHENTICATION_URL endpoint. URLs must use HTTP(S) and cannot embed credentials.

| Setting | Default | Validation |
| --- | --- | --- |
| I3D_RETRY_ATTEMPTS | 3 | 1–10 |
| I3D_RETRY_DELAY | 1.5s | Nonnegative |
| I3D_RETRY_MAX_DELAY | 7.5s | At least the initial delay |
| I3D_ALLOCATION_TIMEOUT | 120s | Positive |
| I3D_PROVIDER_TIMEOUT | 90s | Positive |
| I3D_RECONCILE_INTERVAL | 1m | Nonnegative; 0s disables polling |
| I3D_RECONCILE_TIMEOUT | 30s | Positive |
| I3D_RECONCILE_GRACE_PERIOD | 2m | Nonnegative |

Runtime and process environment values use Go duration strings such as 250ms or 2m. Programmatic Config fields use time.Duration. JSON duration fields use integer nanoseconds; setting.json can omit them to retain defaults.

Example setting.json (keep real credentials outside version control):

~~~json
{
  "oneApi": {
    "applicationId": "YOUR_APPLICATION_ID",
    "baseUrl": "https://api.i3d.net",
    "token": "REPLACE_WITH_SERVER_TOKEN"
  }
}
~~~

These settings configure the adapter, not Nakama's server HTTP key. Configure runtime.http_key separately for trusted headless-server lifecycle RPCs. Never distribute either credential to players.
