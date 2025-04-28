# Nakama i3D Plugin

This plugin integrates [Nakama's](https://heroiclabs.com/nakama/) matchmaking with [i3D.net's](https://www.i3d.net/) fleet management services, allowing Nakama to create matches for players and allocate them to servers hosted or managed by i3D.net. It is designed to help game developers using Nakama to quickly get started hosting multiplayer games on the i3D.net platform. 

**N.B.**  Prior to working with the plugin, it's highly recommened you familiarize yourself with Nakama's [Matchmaker documentation](https://heroiclabs.com/docs/nakama/concepts/multiplayer/matchmaker/) and i3D's [Game Integration Overview](https://www.i3d.net/docs/one/odp/Game-Integration/).


## Getting Started

The plugin contains a Nakama `FleetManager` implementation which integrates with i3D. After registering the `FleetManager` with Nakama it hooks Nakama's `MatchmakerMatched` event to allocate and manage game servers at i3D. 

Depending on your game's requirements, the plugin module might need customising with bespoke code, but a default Nakama module is implemented in `main.go` which configures and registers the `FleetManager` with Nakama.

Before getting started with the plugin, a number of pre-requisites need to be met:

### Pre-requisites

- Nakama running locally, in a docker container, or on a remote server.
- An active i3D.net account with one or more hosts registered.
- A game server binary deployed to one or more of the i3D hosts.
- If you're building the plugin locally you'll also need `go 1.23.2` installed.

Once the pre-requisites have been met, there are several steps to get the plugin working.

### 1. Install Plugin Module

Download the plugin source code, using the following command:

```
git clone https://github.com/codewizards/nakama-i3d.git
```

If you don't have Git installed, you can download a zip file and unzip it into a suitable path on your development computer.

### 2. Configure Plugin

The plugin's configuration is stored in the `config.json` file, although configuration values can also be provided using environment variables if preferred.

To configure the plugin, the `fleet_manager_api_token` value needs updating with an i3D API token for secure access to i3D. 

Additionally, the `application_id` value can optionally be updated with the i3D `Application ID` hosting your game servers, otherwise it will need providing through Nakama's matchmaking metadata, as described in the next section.

### 3. Building the plugin

Once you have the source code and configuration values set, navigate to the root folder in a terminal window and run the command:

```
docker compose up --build
```

This will launch a new docker container that builds and runs the plugin. Alternatively, if you want to deploy the plugin to an existing Nakama instance, run the command:

```
go build --trimpath --buildmode=plugin -o ./i3dplugin.so
```

Then copy the shared object library `i3dplugin.so` into the `/data/modules` folder of your Nakama instance.

**N.B.** Ensure you run this command in a Linux shell, either on a Linux machine or using Windows Subsystem for Linux (WSL). It will not work in a Windows command prompt.

### 4. Call Matchmaking Client

When creating Nakama matchmaking tickets from your game client you can provide key-value pairs as properties, which can be used for controlling the i3D allocation and for transmitting metadata to the game server.

A key use for these properties is to control i3D allocation, allowing i3D to choose a host based on matchmaking properites. For example, setting an `applicationId` property allows the i3D `Application ID` to be specified, rather than using the value from the plugin configuration described earlier.

The default module works with the following properties:

* `applicationId` - Allows the i3D `Application ID` to be specified and override the configured value.
* `region` - Allows the i3D `Region` to be specified. 
* `filters` - Allows i3D allocation filters to be specified, as described later.

Although the actual properties you'll need for your game will depend on the specific i3D allocation rules you want to implement, the default module provides a flexible set for getting started.


### 5. Plugin Testing

After configuring the plugin and ensuring the matchmaking tickets provide the requisite properties, you can test your deployment by running enough game clients to trigger a successful matchmaking. 

If things don't run smoothly any errors should be reported in the terminal window. Often a lot of text is produced but if you read it carefully the answer is normally shown in the output.

## Customising the plugin

While the default module makes it easy getting started with Nakama and i3D, you can customise the module if it doesn't meet your game's specific requirements. 

Customising the module can be done in a variety of ways, but most commonly involves: changing module configuration; adding logic to the `MatchmakerMatched` event handler; or modifying the metadata or filters.

All of these steps are demonstrated in the supplied `main.go` file. This is a sample implementation showing how the i3D Fleet Manager could potentially be used. This file should only be considered a reference and your own code should reflect your intended usage and the parameters that are applicable to your particular system.

### Module initialisation

The `InitModule` is responsible for initialising the module, and configuring the module by loading the `config.json` file and reading any matching environment variables.

The configuration can be adjusted to suit the needs of the game. You can add or remove properties to the configuration object as required, of change verification logic to the `VerifyFleetManagerConfig` function in the `fleet_manager_config.go` file.


### MatchMakerMatched

The `MatchmakerMatched` event handler is registered with the Nakama Initializer by calling the `RegisterMatchmakerMatched` function, passing in the `MatchmakerMatched` function. 

Triggered once Nakama successfully matches two or more players to create a new match. In this function we retrieve the `FleetManager` registered with Nakama and call the `Create` method to create a new server instance to host the newly created match. This `Create` method also accepts a callback function so individual players can be notified once the game server is started and available.

The `MatchMakerMatched` event handler function takes various parameters including the Nakama module itself and a list of entries containing the information about the players and the matchmaking ticket that was used to create this Match.

```
func MatchmakerMatched(
    ctx context.Context, 
    logger runtime.Logger, 
    db *sql.DB, 
    nk runtime.NakamaModule, 
    entries []runtime.MatchmakerEntry
) (string, error)
```

A callback is then registered on the `Create` function to do any processing once a server has been successfully allocated and started. For example, using the Nakama module to send a notification to each client that they can now connect to the server.

```
var callback runtime.FmCreateCallbackFn = func(
    status runtime.FmCreateStatus,
    instanceInfo *runtime.InstanceInfo,
    sessionInfo []*runtime.SessionInfo,
    metadata map[string]any, err error
)
```

The final part of this function creates metadata required by the game server, which is very likely to be specific to the actual game. This can be hard-coded, loaded from config or read from the matchmaker data passed in by Nakama. Once everything is set, the `Create` method can be called on the `FleetManager` which will then call i3D to find and allocate a free server instance for the match to run on for the appropriate number of players.

```
metadata := make(map[string]any)
metadata["players"] = "5"
metadata["applicationId"] = properties["applicationId"]

err := fm.Create(ctx, maxplayers, userIds, nil, metadata, callback)
```

### Filters
One of the most common requirements when using the i3D Fleet Manager is to be able to filter the servers that are returned by the Fleet Manager. This is achieved by setting the `filters` property in the metadata passed to the `Create` method.

The `filters` property is a string that is passed to the Fleet Manager API and is used to filter the servers that are returned. It takes the form 
    
    "property=value AND property=value AND ...."

and can be set in the GameServer when creating the matchmaking ticket.	

```
metadata["filters"] = "hostId=3762469 AND applicationId=6268349608002583795"
```

The i3D api allows for a number of different filters to be set, these are defined in the `fleet_manager_api_client.go` file and are as follows:

```
 - applicationBuildId
 - applicationBuildName
 - applicationId
 - applicationName
 - applicationType
 - arcusAvailable
 - createdAt
 - dcLocationId
 - dcLocationName
 - deploymentEnvironmentId
 - deploymentEnvironmentName
 - executable
 - fleetId
 - fleetName
 - hostId
 - isVirtual
 - regionId
 - regionName
 - startupParams
```

### HealthCheck.go

The healthCheck function is declared in a separate file and is an example of creating your own RPC which can be registered with Nakama. This function will be available in your game server and can be called in the Nakama UI.

To register the RPC, add the following lines to the `InitModule` function in `main.go`:

```
	//Register the healthcheck RPC
	err = initializer.RegisterRpc("healthcheck", RpcHealthCheck)
	if err != nil {
		logger.Error("Error registering healthcheck RPC: %v", err)
		return err
	}
```

To verify the function is working, log in to the Nakama Console at the address http://localhost:7351 using the default credentials

```
Username: admin
Password: password
```

Once you have logged in, navigate to the `Accounts` section and copy the `UserId` from the default `User Account`, then in the `Api Explorer` section you can select the `HealthCheck` function from the drop down and paste the copied `UserId` into the `UserId` box. Clicking the `Send Request` button should result in a response message from the HealthCheck RPC indicating that the call was a success.

You can also call the function via the CLI using the following command:

```
curl -X POST http://localhost:7351/v2/rpc/healthcheck -d '{"user_id": "default_user"}'
```

================================================================================================

