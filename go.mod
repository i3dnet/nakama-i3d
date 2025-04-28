module github.com/codewizards/nakama-i3d/main

go 1.23.5

require (
	github.com/codewizards/nakama-i3d/fleetmanager v0.0.0-00010101000000-000000000000
	github.com/heroiclabs/nakama-common v1.36.0
)

require google.golang.org/protobuf v1.36.4 // indirect

replace github.com/codewizards/nakama-i3d/fleetmanager => ./fleetmanager

replace github.com/codewizards/nakama-i3d/mockserver => ./mockServer
