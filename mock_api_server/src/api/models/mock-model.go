package models

import "time"

func GetApplicationInstance() *ApplicationInstance {
	timestamp := time.Now().Unix()
	properties := make([]Property, 0, 1)
	properties = append(properties, Property{
		PropertyKey:   "port",
		PropertyValue: "7777",
		PropertyType:  1,
		ID:            "235235235",
	})

	IPAddresses := make([]IPAddress, 0, 1)
	IPAddresses = append(IPAddresses, IPAddress{
		Private:   0,
		IPAddress: "127.0.0.1",
		IPVersion: 4,
	})

	return &ApplicationInstance{
		Id:                        "723709572903",
		DeploymentEnvironmentID:   "235235",
		DeploymentEnvironmentName: "Mock Environment",
		FleetID:                   "34325423",
		FleetName:                 "Mock Fleet",
		HostID:                    23344,
		IsVirtual:                 0,
		ApplicationID:             "725092305",
		ApplicationName:           "Mock Application",
		ApplicationType:           2,
		ApplicationBuildID:        "23523524352",
		ApplicationBuildName:      "Mock Build",
		DCLocationID:              6,
		DCLocationName:            "Rotterdam",
		RegionID:                  "2352352352",
		RegionName:                "Mock Region",
		Status:                    4,
		CreatedAt:                 timestamp,
		StartedAt:                 timestamp,
		StoppedAt:                 timestamp,
		PID:                       1234,
		PIDChangedAt:              timestamp,
		StartupParams:             "",
		Executable:                "game.exe",
		ManuallyDeployed:          0,
		Properties:                properties,
		IPAddresses:               IPAddresses,
		LabelReadOnly:             []KeyValue{},
		LabelPublic:               []KeyValue{},
		Metadata:                  []KeyValue{},
		NumPlayersMax:             2,
		NumPlayers:                0,
		LiveHostName:              "test",
		LiveMap:                   "test",
		LiveGameVersion:           "1",
		UpdatedAt:                 timestamp,
		LiveRules:                 "test",
		AutoRestart:               1,
		MarkedForDeletion:         0,
		ArcusAvailable:            0,
	}
}
