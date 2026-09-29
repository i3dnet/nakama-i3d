# i3D.net Integration using the Nakama i3D Plugin

This guide details the integration process between Nakama's powerful social matchmaking system and i3D.net's session management, enabling developers to efficiently scale multiplayer games using i3D.net's global infrastructure.

# Prerequisites

Nakama Version

The Nakama-i3D plugin has been tested with Nakama 3.36 and above.

Before proceeding ensure that you have:

* Created an i3D.net Account with registered hosts and deployed game server binaries
* Installed Nakama server
* Installed the appropriate Arcus Protocol SDK (Unity, Unreal, or C++)
* Go version 1.23.5 or higher

# Architecture Overview

An overview of how Nakama orchestrates the i3D.net fleet:

The Nakama-i3D integration leverages i3D.net's One API to manage game server allocations. When players are matched through Nakama's matchmaking system, the plugin automatically:

1. Requests a game server allocation from i3D.net
2. Waits for the server to become ready (Arcus protocol status)
3. Notifies players with connection details
4. Manages the server lifecycle through status updates

# Installing the Nakama-i3D Plugin

This guide assumes you:

* Have a Nakama server project up and running using the Go runtime. Follow the Introduction to Nakama Go Runtime documentation to get started, if needed.
* Are familiar with hosting a Nakama instance. See our Heroic Cloud documentation for details on how to launch an instance of Nakama in the cloud.

This guide focuses on implementing the necessary functionality to communicate with i3D.net in order to allocate game servers and manage game sessions.

The Nakama-i3D plugin provides an i3D.net implementation of Nakama's Fleet Manager interface and is available as an Open-Source plugin via GitHub.

To install the plugin in your Nakama project, run the following command in the project folder:

| 1 | go get github.com/i3d/nakama-fleetmanager |
| - | ----------------------------------------- |

Once installed, a new `FleetManager` instance can be created within the `InitModule` function:

```go
import (
    "github.com/i3d/nakama-fleetmanager"
    "github.com/heroiclabs/nakama/runtime"
)

func InitModule(ctx context.Context, nk runtime.NakamaModule, initializer runtime.Initializer, logger runtime.Logger) error {
    // Load config
    cfg, err := fleetmanager_config.NewConfigFromRuntime(ctx)
    if err != nil {
        cfg, err = fleetmanager_config.NewConfig()
        if err != nil {
            return fmt.Errorf("config error: %w", err)
        }
    }

    // Create and register fleet manager
    fm, err := fleetmanager.NewI3dFleetManager(ctx, logger, initializer, nk, cfg)
    if err != nil {
        return err
    }
    if err := initializer.RegisterFleetManager(fm); err != nil {
        return err
    }

    // Register matchmaker handler
    return initializer.RegisterMatchmakerMatched(MatchmakerMatched)
}
```

# Maintaining State Synchronization

The i3D.net integration uses the Arcus protocol to maintain state synchronization between your game servers and Nakama. The Arcus protocol ensures that:

1. Game servers report their readiness status before players can connect
2. Metadata from Nakama (such as map selection) is passed to the game server
3. Server lifecycle is properly managed through status updates

Important!

Your game server MUST implement the Arcus protocol to work with this integration. The protocol handles critical lifecycle events and ensures players only connect to ready servers.

## Server Status Management

The Arcus protocol defines several status codes that your game server must implement:

* **Status 2 (Starting)**: Server is initializing
* **Status 3 (Ready)**: Server is ready to accept players
* **Status 4 (Online)**: Game has finished, server is available for reallocation
* **Status 5 (Error)**: Server encountered an error

# Creating the Headless Game Server

This guide covers the required steps to get your headless game server configured to work with Nakama and i3D.net. It does not cover synchronizing game state between the headless server and the client, for this please see the documentation for your chosen networking framework (e.g. Nakama, Unity Netcode for GameObjects, Mirror, etc).

## Installing the Arcus Protocol SDK

The Arcus protocol is available for multiple platforms:

### Unity

1. Download the [Unity Arcus SDK](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/sdk-unity-plugin)
2. Import the package into your Unity project
3. Configure the Arcus component on your server initialization object

### Unreal Engine

1. Download the [Unreal Arcus Plugin](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/sdk-unreal-plugin)
2. Install the plugin in your Unreal project
3. Initialize the Arcus subsystem in your game mode

### C++

1. Follow the [C++ Integration Guide](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/integration-guide)
2. Link the Arcus library to your project
3. Initialize the Arcus client in your server startup code

## Implementing Server Lifecycle

Your game server must handle the following lifecycle events:

### Server Initialization

When your server starts, it should:

1. Initialize the Arcus protocol
2. Read metadata passed from Nakama (map, game mode, etc.)
3. Set status to Ready (3) when initialization is complete

```csharp
// Unity Example
void Start() {
    // Initialize Arcus
    arcusClient.Initialize();
    
    // Read metadata
    var metadata = arcusClient.GetMetadata();
    var mapName = metadata["map"];
    
    // Load the specified map
    LoadMap(mapName);
    
    // Set status to Ready
    arcusClient.SetStatus(ArcusStatus.Ready);
}
```

### Game Completion

After a game session ends:

1. Clean up game state
2. Set status to Online (4) to allow reallocation
3. The i3D system will automatically restart your server with new ports

```csharp
// Unity Example
void OnGameEnd() {
    // Clean up game state
    CleanupGameSession();
    
    // Set status to Online for reallocation
    arcusClient.SetStatus(ArcusStatus.Online);
}
```

# Configuring i3D Services

## Configuration Methods

The plugin can be configured through two methods:

### 1. Environment Variables (Recommended for Production)

Set these environment variables in your Nakama deployment:

| Variable                 | Required | Description                                 | Default               |
| ------------------------ | -------- | ------------------------------------------- | --------------------- |
| `I3D_APPLICATION_ID`     | Yes      | i3D application (game) ID                   | —                     |
| `I3D_ACCESS_TOKEN`       | Yes      | REST API token                              | —                     |
| `I3D_USE_BEARER_AUTH`    | No       | Use M2M OAuth2 authentication (true/false)  | `false`               |
| `I3D_CLIENT_ID`          | Cond.    | M2M client ID (when using bearer auth)      | —                     |
| `I3D_CLIENT_SECRET`      | Cond.    | M2M client secret (when using bearer auth)  | —                     |
| `I3D_AUDIENCE`           | Cond.    | OAuth2 audience (when using bearer auth)    | —                     |
| `I3D_AUTHENTICATION_URL` | Cond.    | OAuth2 token URL (when using bearer auth)   | —                     |
| `I3D_API_URL`            | No       | One API base URL                            | `https://api.i3d.net` |
| `I3D_RETRY_ATTEMPTS`     | No       | Retry count for allocation calls            | `3`                   |
| `I3D_RETRY_DELAY`        | No       | Initial retry delay (e.g. `1500ms`)         | `1500ms`              |
| `I3D_RETRY_MAX_DELAY`    | No       | Max backoff delay (e.g. `7500ms`)          | `7500ms`              |

### 2. Configuration File (local.yml)

Alternatively, create a `local.yml` file in your Nakama data directory with the configuration values.

## Authentication Setup

### Basic Authentication

For basic authentication, you only need:
- `I3D_APPLICATION_ID`: Your application ID from i3D.net
- `I3D_ACCESS_TOKEN`: Your API access token

### M2M OAuth2 Authentication

For enhanced security using OAuth2:
1. Set `I3D_USE_BEARER_AUTH=true`
2. Provide your OAuth2 credentials:
   - `I3D_CLIENT_ID`
   - `I3D_CLIENT_SECRET`
   - `I3D_AUDIENCE`
   - `I3D_AUTHENTICATION_URL`

# Implementing Matchmaking

## Handling Matchmaker Events

When players are matched, allocate a game server using the fleet manager:

```go
func MatchmakerMatched(
    ctx context.Context,
    logger runtime.Logger,
    db *sql.DB,
    nk runtime.NakamaModule,
    entries []runtime.MatchmakerEntry,
) (string, error) {
    // Extract user IDs from matched players
    userIDs := make([]string, len(entries))
    for i, entry := range entries {
        userIDs[i] = entry.Presence.UserId
    }

    // Define callback for when server is ready
    callback := func(
        status runtime.FmCreateStatus,
        inst *runtime.InstanceInfo,
        sessions []*runtime.SessionInfo,
        metadata map[string]any,
        err error,
    ) {
        if err != nil {
            logger.Error("Failed to create game session", err)
            return
        }

        // Notify players with connection details
        for _, session := range sessions {
            notification := map[string]any{
                "IpAddress": inst.IpAddress,
                "Port":      inst.Port,
                "SessionId": session.SessionId,
            }
            
            content, _ := json.Marshal(notification)
            nk.NotificationSend(ctx, session.UserId, "Game Session Created", content, 9000, "", false)
        }
    }

    // Metadata for the game server
    metadata := map[string]any{
        "map": "desert_arena",
        "gameMode": "team_deathmatch",
        "maxPlayers": len(entries),
    }

    // Create the game session
    fm := nk.GetFleetManager()
    return "", fm.Create(ctx, len(entries), userIDs, nil, metadata, callback)
}
```

## Using Allocation Filters

The plugin supports filtering server allocations by various criteria:

```go
// Create a filter builder
fb := fleetmanager.NewFilterBuilder()

// Add filters as needed
fb.Add(fleetmanager.FleetId, "your-fleet-id")
fb.Add(fleetmanager.RegionId, "eu-west")
fb.Add(fleetmanager.ApplicationBuildId, "v1.2.3")

// Apply filters to metadata
metadata := fb.AddFiltersToMetaData(map[string]any{
    "map": "desert_arena",
    "gameMode": "ranked",
})

// Create session with filters
fm.Create(ctx, maxPlayers, userIDs, nil, metadata, callback)
```

### Available Filters

- `deploymentEnvironmentId` / `deploymentEnvironmentName`
- `fleetId` / `fleetName`
- `hostId`
- `applicationBuildId` / `applicationBuildName`
- `dcLocationId` / `dcLocationName`
- `regionId` / `regionName`

# Advanced Features

## Health Check RPC

Register a health check endpoint to monitor the fleet manager status:

```go
func InitModule(ctx context.Context, ...) error {
    // ... other initialization code ...
    
    // Register health check RPC
    if err := initializer.RegisterRpc("healthcheck", fleetmanager.RpcHealthCheck); err != nil {
        return err
    }
    
    return nil
}
```

## Session Listing

Query active game sessions (limited to occupied servers):

```go
func ListActiveSessions(ctx context.Context, nk runtime.NakamaModule) {
    fm := nk.GetFleetManager()
    
    query := "+value.playerCount:2" // Sessions with 2+ players
    limit := 10
    cursor := ""
    
    for {
        instances, nextCursor, err := fm.List(ctx, query, limit, cursor)
        if err != nil {
            logger.Error("Failed to list sessions", err)
            break
        }
        
        // Process instances
        for _, instance := range instances {
            logger.Info("Active session", 
                "id", instance.Id,
                "players", instance.PlayerCount,
                "ip", instance.IpAddress,
                "port", instance.Port)
        }
        
        if nextCursor == "" {
            break
        }
        cursor = nextCursor
    }
}
```

# Best Practices

## Server Lifecycle Management

1. **Always exit after game completion**: After players leave, exit the game server process. i3D.net will automatically restart it with new ports, ensuring a clean state for the next session.

2. **Handle errors gracefully**: Set status to Error (5) if initialization fails, allowing i3D.net to handle the faulty instance.

3. **Implement connection validation**: Use the Arcus protocol to ensure players only connect to ready servers.

## Performance Optimization

1. **Use appropriate retry settings**: Configure retry attempts and delays based on your latency requirements.

2. **Implement regional allocation**: Use region filters to allocate servers close to your players.

3. **Monitor allocation times**: Use the health check RPC to monitor fleet manager performance.

# Limitations

Be aware of the following limitations when using the Nakama-i3D plugin:

1. **List functionality**: The `List` method only returns occupied game servers, as i3D.net focuses on allocation rather than session management.

2. **Latency-based allocation**: The plugin does not implement automatic latency-based server selection. Implement your own logic using region filters.

3. **Join operations**: The `Join` method only validates player capacity against max players. It does not communicate with i3D.net as player management is handled directly between clients and game servers.

# FAQ

**My game server isn't receiving metadata from Nakama**

Ensure your game server properly initializes the Arcus protocol and reads metadata during startup. Check the Arcus SDK documentation for your platform.

**Players can't connect to the allocated server**

Verify that:
1. Your game server sets status to Ready (3) before the callback is triggered
2. The correct ports are configured in your i3D.net deployment
3. Network security rules allow player connections

**How do I handle server crashes?**

Implement proper error handling in your game server and set status to Error (5) when critical failures occur. i3D.net will handle the cleanup and replacement.

**Can I reuse game servers for multiple sessions?**

Yes, set status to Online (4) after a game completes. However, we recommend exiting the process to ensure a clean state and prevent memory leaks.

**How do I implement region-based matchmaking?**

Use the region filters when creating sessions to ensure players are allocated servers in their preferred regions. Combine this with Nakama's matchmaking properties for optimal results.

### Table of Contents

1. [Prerequisites](#prerequisites)
2. [Architecture Overview](#architecture-overview)
3. [Installing the Nakama-i3D Plugin](#installing-the-nakama-i3d-plugin)
4. [Maintaining State Synchronization](#maintaining-state-synchronization)
5. [Creating the Headless Game Server](#creating-the-headless-game-server)
6. [Configuring i3D Services](#configuring-i3d-services)
7. [Implementing Matchmaking](#implementing-matchmaking)
8. [Advanced Features](#advanced-features)
9. [Best Practices](#best-practices)
10. [Limitations](#limitations)
11. [FAQ](#faq)
