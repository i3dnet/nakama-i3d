module gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/example

go 1.23.5

require (
	github.com/heroiclabs/nakama-common v1.36.0
	gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager v0.0.0-00010101000000-000000000000
)

require google.golang.org/protobuf v1.36.4 // indirect

replace gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager => ../../fleetmanager/src
