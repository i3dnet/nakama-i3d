package models

type ApplicationInstance struct {
	Id                        string      `json:"id"`
	DeploymentEnvironmentID   string      `json:"deploymentEnvironmentId"`
	DeploymentEnvironmentName string      `json:"deploymentEnvironmentName"`
	FleetID                   string      `json:"fleetId"`
	FleetName                 string      `json:"fleetName"`
	HostID                    int         `json:"hostId"`
	IsVirtual                 int         `json:"isVirtual"`
	ApplicationID             string      `json:"applicationId"`
	ApplicationName           string      `json:"applicationName"`
	ApplicationType           int         `json:"applicationType"`
	ApplicationBuildID        string      `json:"applicationBuildId"`
	ApplicationBuildName      string      `json:"applicationBuildName"`
	DCLocationID              int         `json:"dcLocationId"`
	DCLocationName            string      `json:"dcLocationName"`
	RegionID                  string      `json:"regionId"`
	RegionName                string      `json:"regionName"`
	Status                    int         `json:"status"`
	CreatedAt                 int64       `json:"createdAt"`
	StartedAt                 int64       `json:"startedAt"`
	StoppedAt                 int64       `json:"stoppedAt"`
	PID                       int         `json:"pid"`
	PIDChangedAt              int64       `json:"pidChangedAt"`
	StartupParams             string      `json:"startupParams"`
	Executable                string      `json:"executable"`
	ManuallyDeployed          int         `json:"manuallyDeployed"`
	Properties                []Property  `json:"properties"`
	IPAddresses               []IPAddress `json:"ipAddress"`
	LabelReadOnly             []KeyValue  `json:"labelReadOnly"`
	LabelPublic               []KeyValue  `json:"labelPublic"`
	Metadata                  []KeyValue  `json:"metadata"`
	NumPlayersMax             int         `json:"numPlayersMax"`
	NumPlayers                int         `json:"numPlayers"`
	LiveHostName              string      `json:"liveHostName"`
	LiveMap                   string      `json:"liveMap"`
	LiveGameVersion           string      `json:"liveGameVersion"`
	UpdatedAt                 int64       `json:"updatedAt"`
	LiveRules                 string      `json:"liveRules"`
	AutoRestart               int         `json:"autoRestart"`
	MarkedForDeletion         int         `json:"markedForDeletion"`
	ArcusAvailable            int         `json:"arcusAvailable"`
}

type OriginalMetadata struct {
	Metadata []map[string]any `json:"metadata"`
}

type KeyValueMetadata struct {
	Metadata []KeyValue `json:"metadata"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Property struct {
	ID            string `json:"id"`
	PropertyType  int    `json:"propertyType"`
	PropertyKey   string `json:"propertyKey"`
	PropertyValue string `json:"propertyValue"`
}

type IPAddress struct {
	IPAddress string `json:"ipAddress"`
	IPVersion int    `json:"ipVersion"`
	Private   int    `json:"private"`
}
