# Nakama And i3D Integration

## Table of Contents

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
	* [Arcus protocol](#arcus-protocol)
3. [Installation](#installation)
4. [Quick Start](#quick-start)
	* [Extra Information](#extra-information)
5. [Configuration](#configuration)

   * [Configuration Sources](#configuration-sources)
   * [Environment Variables](#environment-variables)
6. [Building the Plugin](#building-the-plugin)
7. [Usage](#usage)

   * [Integrating in Nakama](#integrating-in-nakama)
   * [Matchmaker Events](#matchmaker-events)
   * [Using Filters](#using-filters)
   * [Listing Sessions](#listing-sessions)
   * [Health Check RPC](#health-check-rpc)
   * [Fleetmanager RPCs](#fleetmanager-rpcs)
8. [Limitations](#limitations)
9. [Advice](#advice)
10. [Contributing](#contributing)
11. [License](#license)


## 1. Introduction

This plugin implements Nakama’s [fleet manager interface](https://heroiclabs.com/docs/nakama/concepts/multiplayer/fleet-manager/) using i3D.net’s [One API](https://docs.i3d.net/game-hosting/game-integration/overview). It enables:

* **Matchmaking**: Use Nakama’s social and matchmaking features.
* **Server allocation**: Automatically spin up game servers on i3D hosts.
* **Notifications**: Inform clients when sessions are ready.

> **Note:** Familiarize yourself with Nakama’s [Matchmaker docs](https://heroiclabs.com/docs/nakama/concepts/multiplayer/matchmaker/) and i3D’s [Game Integration Overview](https://docs.i3d.net/game-hosting/game-integration/overview) before proceeding.

## 2. Prerequisites

* **i3D.net account** with registered hosts and deployed game server binaries.
* **Go** ≥ 1.23.5
* **Nakama Server** ≥ 1.36.0
* **Arcus Protocol SDK** on your game server (Unity, Unreal, or C++).

## 2.1 Arcus protocol
To make the game server work with our platform it is recommended to use the Arcus protocol. This available for:

- [Unity](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/sdk-unity-plugin)
- [Unreal](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/sdk-unreal-plugin)
- [c++](https://docs.i3d.net/game-hosting/game-integration/sdk-overview/integration-guide)

The protocol makes sure before players are send to the game server is running and also metadata needed by the game
server will be send by the i3D Api to the game server. This ensures you that the players will start the game in the
right map. After the game has finished you need to update the status inside of the game server to 4 (online) so that the
game server can be allocated again.

## 3. Installation

```bash
go get github.com/i3d/nakama-fleetmanager
```

## 4. Quick start

1. Clone the repo:

   ```bash
   git clone https://github.com/i3d/nakama-fleetmanager.git
   cd nakama-fleetmanager
   ```
2. Start Nakama, PostgreSQL, and a mock i3D server:

   ```bash
   cd _infra
   docker compose up --build -d
   ```
3. Observe logs for matchmaking and session notifications:

   ```bash
   docker logs -f nakama
   ```

### 4.1 Extra information
For a quick running example we have added a test client, the setting is that 2 clients need to run before you have a
match. In the docker-compose file we start already one client. When you run on you local machine an client as well than
automatically a match will be created. and you will see in you console a message with a notification that there is a
match started:

```logs
2025-05-06 06:37:35 DBG Join: added matchmaker ticket "d3807ef7-e2fc-4286-9643-fa8948f01d48"
2025-05-06T06:39:18.487227482Z 2025-05-06 06:39:18 DBG MatchmakerMatched: ticket:"d3807ef7-e2fc-4286-9643-fa8948f01d48"  token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDY1MTM1ODgsIm1pZCI6IjEyMzAwOGFmLWM5ZTctNDczZS1hNTA0LTYzZGFiZDY5ZTJmOS4ifQ.fNOdx5cEKNrXRTIpSy9mcTcekre1iqKj-yBsYq7-6aM"  users:{presence:{user_id:"3a9a787e-5b4b-4e63-9e11-30c562f3c415"  session_id:"97890f43-2a44-11f0-a578-006100a0eb06"  username:"d0cqsb9rb554lb4p6umg"}}  users:{presence:{user_id:"1c5b15bc-ec15-4c26-b287-e28a48b9a600"  session_id:"d08333f9-2a44-11f0-a578-006100a0eb06"  username:"d0cqt3ca1a8r6n6uicv0"}}  self:{presence:{user_id:"3a9a787e-5b4b-4e63-9e11-30c562f3c415"  session_id:"97890f43-2a44-11f0-a578-006100a0eb06"  username:"d0cqsb9rb554lb4p6umg"}}
2025-05-06T06:39:18.492070259Z 2025-05-06 06:39:18 DBG Notifications: notifications:{id:"88430fc1-24c1-4f41-ac4d-f406390e3478"  subject:"Game Session Created"  content:"{\"IpAddress\":\"127.0.0.1\",\"Port\":7777,\"SessionId\":\"\"}"  code:9000  sender_id:"00000000-0000-0000-0000-000000000000"  create_time:{seconds:1746513558}}
```

## 5. Configuration

### 5.1 Configuration Sources
You can figure the plugin in multiple ways:

- ***local.yml*** - This file contains the configuration values for the plugin. You can set the values in this file.
- ***.env*** - This contains the environment variables for the plugin. You can set them in docker compose or on you
  server directly as environment variables.

### 5.2 Environment Variables
| Variable                 | Required | Description                                 | Default               |
| ------------------------ | -------- | ------------------------------------------- | --------------------- |
| `I3D_APPLICATION_ID`     | Yes      | i3D application (game) ID.                  | —                     |
| `I3D_ACCESS_TOKEN`       | Yes      | REST API token.                             | —                     |
| `I3D_USE_BEARER_AUTH`    | No       | Use M2M OAuth2 authentication (true/false). | `false`               |
| `I3D_CLIENT_ID`          | Cond.    | M2M client ID (when using bearer auth).     | —                     |
| `I3D_CLIENT_SECRET`      | Cond.    | M2M client secret (when using bearer auth). | —                     |
| `I3D_AUDIENCE`           | Cond.    | OAuth2 audience (when using bearer auth).   | —                     |
| `I3D_AUTHENTICATION_URL` | Cond.    | OAuth2 token URL (when using bearer auth).  | —                     |
| `I3D_API_URL`            | No       | One API base URL.                           | `https://api.i3d.net` |
| `I3D_RETRY_ATTEMPTS`     | No       | Retry count for allocation calls.           | `3`                   |
| `I3D_RETRY_DELAY`        | No       | Initial retry delay (e.g. `1500ms`).        | `1500ms`              |
| `I3D_RETRY_MAX_DELAY`    | No       | Max backoff delay (e.g. `7500ms`).          | `7500ms`              |

> Backoff doubles the delay on each retry up to `I3D_RETRY_MAX_DELAY`.

## 6. Building the Plugin

* **Docker** (build & run):

  ```bash
  docker compose up --build
  ```

* **Native** (build shared object):

  ```bash
  go build --trimpath --buildmode=plugin -o i3dplugin.so
  # Copy i3dplugin.so into /data/modules on your Nakama server
  ```

> **Note:** Run builds in a Linux shell (Linux or WSL).

## 7. Usage

### 7.1 Integrating in Nakama

In your `InitModule` in `main.go`:

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

### 7.2 Matchmaker Events

Handle `MatchmakerMatched` to allocate a game server:

```go
func MatchmakerMatched(
    ctx context.Context,
    logger runtime.Logger,
    db *sql.DB,
    nk runtime.NakamaModule,
    entries []runtime.MatchmakerEntry,
) (string, error) {
    userIDs := extractUserIDs(entries)

    // Callback: notify each user when ready
    callback := func(
        status runtime.FmCreateStatus,
        inst *runtime.InstanceInfo,
        sessions []*runtime.SessionInfo,
        metadata map[string]any,
        err error,
    ) {
        if err != nil {
            logger.Error("creation error: ", err)
            return
        }
        notifyPlayers(ctx, nk, inst, sessions)
    }

    // Create session
    return "", fm.Create(ctx, maxPlayers, userIDs, nil, metadata, callback)
}
```

### 7.3 Using Filters

Build allocation filters:

```go
fb := fleetmanager.NewFilterBuilder()
fb.Add(fleetmanager.FleetId, "<fleet-id>")
fb.Add(fleetmanager.RegionId, "<region-id>")
metadata := fb.AddFiltersToMetaData(map[string]any{"players": "5"})
fm.Create(ctx, 5, userIDs, nil, metadata, callback)
```

available filters:
These are the available filters:
 - deploymentEnvironmentId
 - deploymentEnvironmentName
 - fleetId
 - fleetName
 - hostId
 - applicationBuildId
 - applicationBuildName
 - dcLocationId
 - dcLocationName
 - regionId
 - regionName

### 7.4 Listing Sessions

```go
query, limit, cursor := "+value.playerCount:2", 10, ""
for {
    instances, nextCursor, err := fm.List(ctx, query, limit, cursor)
    if err != nil {
        logger.Error(err)
        break
    }
    process(instances)
    if nextCursor == "" {
        break
    }
    cursor = nextCursor
}
```

### 7.5 Health Check RPC

Register a health check endpoint:

```go
if err := initializer.RegisterRpc("healthcheck", fleetmanager.RpcHealthCheck); err != nil {
    return err
}
```

### 7.6 Fleetmanager RPCs

For Nakama to be up-to-date on the number of Player Sessions currently connected to a Game Session and for it to be aware of when a Game Session is terminated, the Game Session SDK should invoke two RPCs that are automatically exposed as part of the `fleetmanager` registration. These RPCs can be invoked as [Server-to-Server calls](https://heroiclabs.com/docs/nakama/server-framework/runtime-examples/#server-to-server).

#### Update RPC
* RpcId: `update_instance_info`

This RPC should be invoked any time a player connects or disconnects from the Game Session to update the playerCount by passing the current number of connected players.

It can also be used to update a Game Session's `GameSessionName`, `GameSessionData` or `GameProperties`. Updating these metadata fields allow the filtering of Game Sessions in the `List` function based on the properties values.
Each of metadata's expected keys are optional, and only if any is present with a value will it overwrite the current value in the Nakama Storage Index.
The expected RPC payload is the following:
```json
{
  "id": "<Game Session ARN>",
  "playerCount": 10,
  "metadata": {
    "GameSessionData": "<game_session_data>",
    "GameSessionName": "<game_session_name>",
    "GameProperties": {
      "key1": "value1",
      "key2": "value2"
    }
  }
}
```

#### Delete RPC
* RpcId: `delete_instance_info`

This RPC should be invoked any time a Game Session terminates.
The expected RPC payload is the following:
```json
{
  "id": "<Game Session ARN>"
}
```

## 8. Limitations

There are some limitations to the Nakama Fleet Manager plugin. These are:

* The `list` functionality is implemented but not used. Because the i3D API only manage getting the game server and not
  cares about player sessions. The list will only give back game servers that are occupied (have players).
* `Latency` is not implemented. The i3D API does the allocation based on filters, you should self taking care of how you
  want to get the right game server for the user. Because we don't know what your settings will be and how you will
  implement the latency inside of your client.
* `Join` because we don't care about players joining the game server, the join is not needed to let players join the
  game server. But we have build in a game server cache, so when you have a long running session where players can join
  and leave the game server. You can use the join to add new players to the game server. The join doing nothing more
  than checking the **max amount of players** against the players that will join, if the amount is bigger than the max
  amount only the sessions that can join will be given back from the join.

## 9. Advice

To make sure that the game server is always in the right state we recommend that after your players have left the game,
to exit the game. Our system will automatically restart the game server with new ports. This ensures that your game
server is always in the right state when players join your game. May your game suffer from unforeseen memory leaks or
other issues, you ensure in this way that your game server is always in the right state.

## 10. Contributing

Contributions are welcome! Please open issues or pull requests on GitHub.

## 11. License

This project is licensed under the [MIT License](./LICENSE).
