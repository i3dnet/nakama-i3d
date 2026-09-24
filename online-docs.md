# Integrate Nakama with i3D.net

> **Draft for Heroic Labs review — 22 September 2026**
>
> This is a replacement draft for the [current i3D integration guide](https://heroiclabs.com/docs/nakama/guides/concepts/i3d-integration/). The implementation candidate targets Nakama 3.41.0 and has passed local unit/race checks, real plugin loading and the full lifecycle smoke test. It is in review; no release tag or production deployment has been published.
>
> The legacy main revision, `f26ac9d`, uses Nakama 3.26.0 / nakama-common 1.36.0 and a nested GitLab module path. This guide uses the implementation candidate commit `3ce03390a9db2b6573a4213ede9c330f79b43f2d`. Replace that pin with the approved release version after merging and staging validation. The final section lists the remaining publication gates.

Use Nakama to authenticate players and find matches, then allocate a dedicated game server through i3D.net. Players receive the server's connection details through a Nakama notification and connect directly using your game's networking transport.

The [Nakama–i3D integration](https://github.com/i3dnet/nakama-i3d) implements Nakama's [FleetManager interface](https://heroiclabs.com/docs/nakama/concepts/multiplayer/session-based/). It allocates an existing online i3D application instance. Provisioning hosts, deploying builds, and maintaining available server capacity are configured in i3D.net.

## Prerequisites and compatibility

You need:

- A running Nakama deployment with a Go runtime project. See the [Go runtime guide](https://heroiclabs.com/docs/nakama/server-framework/go-runtime/).
- An i3D.net account, a deployed game-server application, and available online instances.
- Credentials that allow Nakama to perform the required One API operations.
- Arcus integration in your headless server to receive allocation metadata and participate in the allocation lifecycle.
- A game client connected to Nakama's realtime socket, with a notification handler and a transport for connecting to your dedicated server.

The candidate has been tested with this exact combination:

| Component | Target |
| --- | --- |
| Nakama | 3.41.0 |
| Nakama Go plugin builder | 3.41.0, matching the runtime image |
| nakama-common | 1.48.0 |
| Go | 1.27.1 |
| Shared protobuf dependency | 1.36.12 |
| i3D integration | Candidate 3ce0339; release tag pending |

The dependency versions come from [Nakama 3.41.0's go.mod](https://github.com/heroiclabs/nakama/blob/v3.41.0/go.mod). The matching plugin loads in Nakama on Linux ARM64 locally; CI also builds and loads it on Linux AMD64. This does not establish compatibility with other runtime versions.

Build the Go plugin using the builder for your Nakama version and deployment architecture. Keep shared dependencies aligned with the runtime; see [dependency pinning](https://heroiclabs.com/docs/nakama/server-framework/go-runtime/go-dependencies/). Rebuild and test the plugin when upgrading Nakama.

## How allocation works

1. Nakama's matchmaker groups players.
2. Your server hook calls `FleetManager.Create` with the matched user IDs, capacity, and game metadata.
3. The adapter requests an available instance from the i3D One API, applying any allocation filters.
4. i3D coordinates allocation with the host agent and game server. Arcus delivers the allocation metadata.
5. After a successful allocation and storage update, the adapter invokes your callback with the instance's connection details.
6. Your callback notifies the players. Each client connects directly to the dedicated game server.

Arcus communicates with i3D's infrastructure. The game server separately calls Nakama's lifecycle RPCs to keep Nakama's session information current.

The relevant i3D application-instance states are:

| Status | Meaning |
| --- | --- |
| 2 — OFFLINE | Deployed but not running |
| 3 — STARTING | Starting up |
| 4 — ONLINE | Initialized and available for allocation |
| 6 — ALLOCATING | Allocation is in progress |
| 5 — ALLOCATED | Assigned to a game session |

The documented allocation flow moves an instance through `ONLINE → ALLOCATING → ALLOCATED` before returning it to the matchmaker. See [i3D's allocation lifecycle](https://docs.i3d.net/game-hosting/game-integration/matchmaker-allocation) and the [One API reference](https://docs.i3d.net/api/api_one).

## Prepare the headless server

Follow the integration instructions for your engine:

- [Unity Arcus SDK](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/sdk-unity-plugin)
- [Unreal Arcus plugin](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/sdk-unreal-plugin)
- [C++ integration guide](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/integration-guide)

Configure the application to use Arcus and implement its allocation handlers. Your server should:

1. Initialize its gameplay transport and Arcus integration.
2. Become ONLINE when it is available for allocation.
3. Handle metadata delivered during allocation, such as map and game mode, and complete the required allocation handshake.
4. Authenticate connecting players and maintain its authoritative player count.
5. Report player-count changes to Nakama and follow the end-of-session policy described below.

Metadata can arrive on each allocation, including when an instance is reused. Handle it in the allocation lifecycle rather than only reading it at process startup. Use the actual API for your Arcus SDK version; the engine-specific guides cover those callbacks.

## Install the Go module

The public module is `github.com/i3dnet/nakama-i3d`. The reusable package retains the name `fleetmanager`.

To evaluate the candidate, run these commands in your Go runtime project:

~~~sh
# Only needed for a new project:
go mod init example.com/nakama-i3d-game

go get github.com/i3dnet/nakama-i3d@3ce03390a9db2b6573a4213ede9c330f79b43f2d
go get github.com/heroiclabs/nakama-common@v1.48.0 google.golang.org/protobuf@v1.36.12
~~~

Use the approved release tag in place of the candidate commit when it is available, and commit the resulting `go.mod` and `go.sum`. The candidate is available through the public Go resolver; `@latest` on the legacy default branch is not the candidate installation path.

Import the library and configuration package as:

~~~go
import (
	fleetmanager "github.com/i3dnet/nakama-i3d"
	fleetconfig "github.com/i3dnet/nakama-i3d/config"
)
~~~

External consumers should not need to clone the repository or add a local `replace` directive. The previous nested-module workaround belongs to the legacy source layout.

## Configure Nakama

The examples below load settings from Nakama's `runtime.env` through `NewConfigFromRuntime`.

A minimal configuration using an i3D API token is:

~~~yaml
shutdown_grace_sec: 15
runtime:
  http_key: "REPLACE_WITH_A_RANDOM_SERVER_ONLY_HTTP_KEY"
  env:
    - "I3D_APPLICATION_ID=YOUR_APPLICATION_ID"
    - "I3D_BASE_URL=https://api.i3d.net"
    - "I3D_ACCESS_TOKEN=YOUR_I3D_API_TOKEN"
~~~

Merge this into your existing Nakama configuration. Keep database, socket, and other deployment settings in place. Supply real credentials through your deployment's secret configuration; keep them out of client builds and source control.

Nakama reads `local.yml` and exposes `runtime.env` to the plugin. Setting an operating-system environment variable alone does not add it to this runtime map. The library's `NewConfig` loader is a separate option for applications intentionally using process-environment configuration.

### Provider settings

| Setting | Requirement | Purpose |
| --- | --- | --- |
| `I3D_APPLICATION_ID` | Required | i3D application ID |
| `I3D_FLEET_ID` | Optional | Scope allocation, provider listing and reconciliation to one fleet |
| `I3D_BASE_URL` | Optional | One API URL; defaults to `https://api.i3d.net` |
| `I3D_ACCESS_TOKEN` | Required for API-token authentication | Token sent in the `PRIVATE-TOKEN` header |
| `I3D_USE_BEARER_AUTH` | Optional; defaults to `false` | Select OAuth client-credentials authentication |
| `I3D_CLIENT_ID` | Required for OAuth | Client ID |
| `I3D_CLIENT_SECRET` | Required for OAuth | Client secret |
| `I3D_AUDIENCE` | Required for OAuth | Audience assigned to your credentials |
| `I3D_AUTHENTICATION_URL` | Required for OAuth | Token endpoint assigned to your credentials |

Use `I3D_BASE_URL` in new configurations. The older `I3D_API_URL` alias is accepted only when `I3D_BASE_URL` is absent.

API-token authentication uses the `PRIVATE-TOKEN` header. For OAuth, set `I3D_USE_BEARER_AUTH=true` and supply the four OAuth settings; an API access token is not additionally required.

### Timeouts and retries

These settings use the same validation in runtime and process configuration:

| Setting | Default | Meaning |
| --- | --- | --- |
| `I3D_RETRY_ATTEMPTS` | `3` | Maximum read attempts, including the initial request (1–10) |
| `I3D_RETRY_DELAY` | `1500ms` | Initial delay before a retry |
| `I3D_RETRY_MAX_DELAY` | `7500ms` | Maximum backoff delay |
| `I3D_ALLOCATION_TIMEOUT` | `120s` | Provider allocation stage |
| `I3D_ALLOCATION_FINALIZE_TIMEOUT` | `30s` | Fresh persistence budget and fresh budget for each cleanup attempt |
| `I3D_PROVIDER_TIMEOUT` | `90s` | Timeout for each provider HTTP request |
| `I3D_RECONCILE_INTERVAL` | `1m` | Poll interval; `0s` disables reconciliation |
| `I3D_RECONCILE_TIMEOUT` | `30s` | Maximum duration of one reconciliation pass |
| `I3D_RECONCILE_GRACE_PERIOD` | `2m` | Minimum allocation age before absence cleanup |
| `I3D_RECONCILE_CLOCK_SKEW` | `5s` | Safety window covering relative node/database clock skew and timestamp precision; minimum 1s |

Only reads retry transient network failures and HTTP 408, 429, 500, 502, 503 or 504, with cancellable backoff. Allocation, restart and writes execute once. A connection failure after sending an allocation can leave its outcome uncertain; inspect provider state before retrying. OAuth uses the complete configured token endpoint with a 30-second HTTP bound and shared refreshes. The timeout defaults are conservative bounds and still need validation against your real Arcus/fleet timings.

### Keep the three authentication paths separate

| Connection | Credential or mechanism |
| --- | --- |
| Nakama → i3D One API | i3D API token or OAuth credentials |
| Headless game server → Nakama lifecycle RPCs | Nakama `runtime.http_key` |
| Player → dedicated game server | Your game's player-authentication protocol |

Arcus coordinates server allocation and metadata. It does not replace player authentication.

## Register the fleet manager

Create `main.go` in your Go runtime project:

~~~go
package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	fleetmanager "github.com/i3dnet/nakama-i3d"
	fleetconfig "github.com/i3dnet/nakama-i3d/config"
)

func InitModule(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	cfg, configErr := fleetconfig.NewConfigFromRuntime(ctx)
	if configErr != nil {
		return configErr
	}

	fm, err := fleetmanager.NewI3dFleetManager(
		ctx, logger, initializer, nk, cfg,
	)
	if err != nil {
		return err
	}
	if err := initializer.RegisterFleetManager(fm); err != nil {
		return err
	}

	return initializer.RegisterMatchmakerMatched(MatchmakerMatched)
}
~~~

The fleet manager registers the `update_instance_info` and `delete_instance_info` lifecycle RPCs. These reject player-session callers and require the trusted server HTTP-key path.

## Allocate a server when players match

Create `matchmaking.go` alongside `main.go`:

~~~go
package main

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	notificationConnectionInfo = 9000
	notificationCreateTimeout  = 9001
	notificationCreateFailed   = 9002
)

func MatchmakerMatched(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	entries []runtime.MatchmakerEntry,
) (string, error) {
	if len(entries) == 0 {
		return "", runtime.NewError("no matched players", 3)
	}

	fm := nk.GetFleetManager()
	if fm == nil {
		return "", runtime.NewError("fleet manager unavailable", 13)
	}

	userIDs := make([]string, 0, len(entries))
	seen := make(map[string]bool, len(entries))
	for _, entry := range entries {
		id := entry.GetPresence().GetUserId()
		if !seen[id] {
			userIDs = append(userIDs, id)
			seen[id] = true
		}
	}

	callback := func(
		status runtime.FmCreateStatus,
		instance *runtime.InstanceInfo,
		sessions []*runtime.SessionInfo,
		metadata map[string]any,
		createErr error,
	) {
		code := notificationCreateFailed
		subject := "Game session unavailable"
		content := map[string]any{"Reason": "allocation_failed"}

		if status == runtime.CreateTimeout || errors.Is(createErr, context.DeadlineExceeded) {
			code = notificationCreateTimeout
			content["Reason"] = "allocation_timeout"
		} else if status == runtime.CreateSuccess &&
			createErr == nil &&
			instance != nil &&
			instance.ConnectionInfo != nil {
			connection := instance.ConnectionInfo
			code = notificationConnectionInfo
			subject = "Game session ready"
			content = map[string]any{
				"InstanceId": instance.Id,
				"IpAddress":  connection.IpAddress,
				"DnsName":    connection.DnsName,
				"Port":       connection.Port,
			}
		}

		if code != notificationConnectionInfo {
			logger.Error("i3D allocation failed: status=%d error=%v",
				status, createErr)
		}

		// The matchmaker hook's context can end before allocation completes.
		// Give notification delivery its own bounded context.
		notifyCtx, cancel := context.WithTimeout(
			context.Background(), 10*time.Second,
		)
		defer cancel()

		notifications := make([]*runtime.NotificationSend, 0, len(userIDs))
		for _, userID := range userIDs {
			notifications = append(notifications, &runtime.NotificationSend{
				UserID:     userID,
				Subject:    subject,
				Content:    content,
				Code:       code,
				Persistent: false,
			})
		}
		if err := nk.NotificationsSend(notifyCtx, notifications); err != nil {
			logger.Error("failed to notify matched players: %v", err)
		}
	}

	// Choose metadata on the server and use string values for game settings.
	metadata := map[string]any{
		"map":       "desert_arena",
		"game_mode": "team_deathmatch",
	}

	_, err := fm.Create(
		ctx, len(userIDs), userIDs, nil, metadata, callback,
	)
	if err != nil {
		logger.Error("i3D allocation request rejected: %v", err)
		return "", err
	}

	// Connection details are delivered by the notification callback.
	return "", nil
}
~~~

The first return from `Create` is provider-specific request metadata; this example does not need it. A successful return means the asynchronous request was accepted. The callback supplies its final result.

Accepted allocation work survives normal hook-context cancellation. Already-canceled calls are rejected before acceptance. Provider allocation and storage persistence have separate deadlines and observe manager shutdown. Storage is committed before success, and a published result wins over simultaneous deadline expiry. A failed or late result with a confirmed allocation ID triggers one best-effort restart with a fresh bounded context. Cleanup failures are logged; ambiguous responses without an allocated instance ID require operator reconciliation. A timed-out dependency that ignores cancellation is reclaimed when it eventually returns, while the process remains alive. Cleanup never changes an already delivered callback outcome or blindly deletes a newer cached session. A programmatic Config with a zero AllocationTimeout logs the 120s fallback once per manager. Callbacks are local to their owning process and do not survive a crash or route between nodes. The separate context in the example handles notification delivery only.

On the client, register your notification handler before submitting a matchmaking ticket:

- `9000`: read `InstanceId`, `IpAddress` or `DnsName`, and `Port`, then connect through your game transport.
- `9001`: allocation timed out; show that outcome and let the player retry matchmaking.
- `9002`: allocation failed; show that outcome and offer a controlled retry.

The notification codes and payload fields are conventions chosen by this example. These notifications are not persistent. Give the client a bounded waiting period and a recovery path for disconnects or missed notifications.

This adapter does not issue i3D player-session credentials. The example therefore does not send a `SessionId`. Validate player identity on the dedicated server using your game's authentication flow.

## Apply allocation filters

To select a fleet and region, replace the metadata construction in `MatchmakerMatched` with the following, and add the `fleetmanager` import to `matchmaking.go`:

~~~go
filters := fleetmanager.NewFilterBuilder()
filters.Add(fleetmanager.FleetId, "YOUR_FLEET_ID")
filters.Add(fleetmanager.RegionId, "YOUR_REGION_ID")

metadata := filters.AddFiltersToMetaData(map[string]any{
	"map":       "desert_arena",
	"game_mode": "team_deathmatch",
})
~~~

Use actual i3D IDs, not an AWS-style region name in an ID field. The builder also exposes name-based filters when appropriate.

Supported filter keys include deployment environment, fleet, host, application build, data-center location, and region. Use the corresponding exported constants, such as `ApplicationBuildId`, `DcLocationId`, or `RegionName`.

Metadata must be JSON-encodable. Accepted metadata retains exact Go scalar types and independent nested maps/slices for asynchronous allocation and callbacks. Structs containing mutable unexported state are rejected before allocation because the adapter cannot safely copy that state.

Allocation filters are separate from the metadata delivered to the game server. The adapter extracts its routing fields before sending game metadata. Reserve `i3dFilters`, `overwriteApplicationId` and the entire `i3d_` namespace for adapter use. `FilterBuilder.Query` returns a raw expression; do not URL-encode it yourself.

The `latencies` argument to `Create` is not used for automatic placement by this adapter. Choose an appropriate region or fleet in your matchmaking logic and pass the filter.

## Keep Nakama's session data current

The dedicated server is responsible for reporting player-count changes. Arcus communication with i3D does not automatically call Nakama's RPCs.

Use Nakama's HTTP-key authentication from trusted server code, following the [headless server authentication guide](https://heroiclabs.com/docs/nakama/guides/concepts/headless-server-auth/). Use HTTPS for remote calls and keep the HTTP key server-side. Lifecycle RPCs reject ordinary player-session callers before decoding their payloads.

### Update the active player count

Call `update_instance_info` after players connect or disconnect. Supply the current count as an absolute value:

~~~http
POST /v2/rpc/update_instance_info?http_key=YOUR_SERVER_HTTP_KEY&unwrap
Content-Type: application/json

{
  "id": "APPLICATION_INSTANCE_ID",
  "player_count": 2,
  "metadata": {
    "map": "desert_arena",
    "game_mode": "team_deathmatch"
  }
}
~~~

Use the i3D application-instance ID and the exact `player_count` field name. Send nonnegative counts and game-owned metadata only. Treat the submitted metadata as the current game metadata you intend to retain; do not rely on an undocumented partial-merge behavior.

The `unwrap` option allows the JSON body to be sent directly. Your Nakama SDK's HTTP-key RPC overload can also make the call.

### End a session

Choose an end-of-session policy that matches your application deployment:

- **Request a restart through the adapter:** after the match finishes, call `delete_instance_info`. This requests an i3D application-instance restart and removes its Nakama session record.
- **Exit the process:** use this when your i3D deployment is configured to restart exited servers.
- **Reuse the running process:** reset game state and return the instance to ONLINE through the supported i3D/Arcus lifecycle.

The restart RPC payload is:

~~~http
POST /v2/rpc/delete_instance_info?http_key=YOUR_SERVER_HTTP_KEY&unwrap
Content-Type: application/json

{
  "id": "APPLICATION_INSTANCE_ID"
}
~~~

Coordinate these policies so the same session does not trigger multiple restart actions. Do not blindly repeat a restart after an ambiguous network failure; check the instance state first.

Always use the connection details returned for the next allocation. Address or port assignments may change, and an application-instance ID can be reused for a later game session.

The adapter takes a complete storage snapshot, then scans every provider page for the configured application/fleet. Failed or partial scans cannot establish absence. Records allocated or updated since the pass start minus I3D_RECONCILE_CLOCK_SKEW are excluded from that pass, including records encountered on later storage pages. Configure this window to cover the maximum relative clock skew between Nakama nodes and storage, plus timestamp precision; it does not protect against unbounded clock drift. Old records with known scope are removed only after two complete observations with unchanged storage versions; concurrent allocations and joins win conflicts. Legacy records with unknown scope are retained when absent. Provider-visible records can establish scope when refreshed. Reconciliation repairs storage only and never restarts an instance.

Absence remains an eventual-consistency assumption: choose a grace period longer than observed provider propagation delays, or disable reconciliation if successful listings cannot reliably establish absence. Continue sending lifecycle reports promptly. Application overrides need a separately scoped reconciliation worker. Graceful shutdown stops background work; clients still need recovery after server crashes.

## Get, list, and join sessions

Retrieve the registered manager with `nk.GetFleetManager()`.

| Method | Purpose |
| --- | --- |
| `Get(ctx, instanceID)` | Retrieve current provider information for an instance |
| `List(ctx, query, limit, cursor)` | List instances from the provider or query Nakama's stored session index |
| `Join(ctx, instanceID, userIDs, metadata)` | Request admission according to the adapter's documented capacity/reservation contract |

An empty `List` query selects the provider-backed path for allocated instances in the configured application and optional fleet. A provider page may contain no ALLOCATED entries but still have a continuation cursor; keep following it. A nonempty query searches Nakama's stored session index using [Nakama query syntax](https://heroiclabs.com/docs/nakama/concepts/multiplayer/query-syntax/).

For example, `+value.player_count:>=2` matches stored sessions with at least two players. It does not test how many seats are available. Use the returned continuation cursor with the same query, limit and listing mode to fetch the next page. Stored listings sort by `value.player_count` then descending `value.create_time`.

An ALLOCATED instance can have zero connected players. Avoid treating allocation state and occupancy as equivalent.

### Join and reservations

The adapter implements local capacity accounting with Nakama versioned writes and bounded conflict retries across nodes. A group can be partially admitted; inspect the returned users. Duplicate users do not increase the count again within the current admission epoch.

Returned SessionId values are empty. Admissions do not expire and are not game-server authentication tokens. An authoritative player_count update replaces the local estimate and resets deduplication because that payload does not identify connected users. Handle abandoned admissions, reconnection and player identity in your game. The adapter does not claim the generic FleetManager interface's complete expiring-reservation contract.

Provider refreshes preserve plugin-owned capacity and local admission estimates. New allocations reset them when an instance ID is reused. Provider-discovered sessions without local capacity cannot admit joins until the application establishes capacity.

## Build and load the plugin

After creating the two Go files, resolve their dependencies:

~~~sh
go mod tidy
~~~

A minimal image build for this guide's target is:

~~~dockerfile
FROM heroiclabs/nakama-pluginbuilder:3.41.0@sha256:eb506121ec2e67f39febee3a4252ae05598a0b39d3517b713ec1577714a9e5b5 AS builder

WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build --trimpath --mod=readonly --buildmode=plugin -o /backend/i3dplugin.so .

FROM heroiclabs/nakama:3.41.0@sha256:ca9f3fe65c90f56860fe5d7cd024868a395cb1a001abb9912bf87ef47d605349

COPY --from=builder /backend/i3dplugin.so /nakama/data/modules/
~~~

This layout assumes the two example files are at your consumer project's root. Add any additional source directories if your project needs them.

Build the image, then use it in your existing Nakama deployment with its database connection, migration/startup command, and mounted configuration:

~~~sh
docker build -t nakama-i3d-example .
~~~

Keep credentials in runtime configuration. Verify that Nakama loads the plugin successfully before testing matchmaking; a successful Go build alone does not establish plugin compatibility.

## Troubleshooting

**The module downloads but I cannot import it.**
Check that you are using the published root-module release at `github.com/i3dnet/nakama-i3d`. The legacy nested layout requires a local replacement and is not the installation path described here.

**The FleetManager interface does not compile, or Nakama rejects the plugin.**
Check the exact Nakama, pluginbuilder, Go, nakama-common and shared dependency versions. Current `Create` returns two values. Build and load against the same supported runtime combination.

**The server does not receive metadata.**
Check that the application uses Arcus and handles metadata during allocation. Confirm the map/game-mode values sent by Nakama and consult the SDK's allocation-handler documentation.

**Players receive connection details but cannot connect.**
Check the allocated instance state, returned public endpoint, configured game transport, firewall rules, and player-authentication flow.

**The update RPC does not change the player count.**
Use `player_count`, supply the application-instance ID, and authenticate as a server with the runtime HTTP key. Check both the RPC response and server logs.

**Players are stuck waiting after allocation.**
Check the allocation result, storage errors, and notification-delivery errors. The client must handle timeout, failure, lost connectivity and missed nonpersistent notifications. In the reviewed open-source Nakama runtime, pending callbacks are local to the process and do not survive restart.

**Can a health-check RPC confirm that allocation works?**
A simple handler health check confirms liveness. Use an actual allocation smoke test and allocation outcome/duration metrics to validate the provider integration.

## Observe allocation and recovery

Nakama exposes i3d_allocation_total / i3d_allocation_duration and i3d_reconciliation_total / i3d_reconciliation_duration, with result tags success, error, timeout and canceled. Alert on persistent reconciliation failures and changes in allocation failure/latency. Labels exclude user and instance IDs. The HTTP client records route templates and status without logging credentials, headers or payload bodies.

## Publication checks for i3D and Heroic Labs

Remove this section and the opening draft note when the release evidence is complete.

Candidate checks: the public root module installs without cloning or replace; the guide's complete Go files and filter example are compiled in a clean consumer; the exact runtime loads the plugin; the two-client smoke validates metadata, notifications, native storage concurrency and cursor sorting, HTTP-key lifecycle payloads, missed-update recovery, invalid readiness and timeout. Unit/race tests cover OAuth, request capture, callback lifetime and concurrent session state.

Publication gates remain:
- Merge the reviewed PR stack and select the release tag (no tags existed at review; v0.1.0 is proposed for the first public root module).
- Repeat clean-consumer installation against the approved tag.
- Identify the currently deployed image/commit, Nakama version, configuration source and game-server authentication; prepare migration and rollback.
- Run an authorized staging allocation/restart test on a designated i3D fleet, including Arcus metadata/readiness and propagation delay.
- Replace the candidate pin with the approved release version, remove the draft/publication notes, and publish the guide.

The user will send this draft to Heroic Labs. No partner message, release tag or deployment has been sent or published by this work.
