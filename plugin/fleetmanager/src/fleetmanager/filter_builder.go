package fleetmanager

import (
	"fmt"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/clients"
	"net/url"
	"strings"
)

// FilterKey represents supported filter keys used in the allocation/filtering API.
type FilterKey string

const (
	DeploymentEnvironmentId   FilterKey = "deploymentEnvironmentId"
	DeploymentEnvironmentName FilterKey = "deploymentEnvironmentName"
	FleetId                   FilterKey = "fleetId"
	FleetName                 FilterKey = "fleetName"
	HostId                    FilterKey = "hostId"
	ApplicationBuildId        FilterKey = "applicationBuildId"
	ApplicationBuildName      FilterKey = "applicationBuildName"
	DcLocationId              FilterKey = "dcLocationId"
	DcLocationName            FilterKey = "dcLocationName"
	RegionId                  FilterKey = "regionId"
	RegionName                FilterKey = "regionName"
	status                    FilterKey = "status"
	applicationId             FilterKey = "applicationId"
)

// FilterBuilder is used to construct filter maps for querying or allocating game servers.
type FilterBuilder struct {
	filters []string
}

// NewFilterBuilder creates a new FilterBuilder instance.
func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{}
}

// Add adds a key-value pair to the filter builder.
func (fb *FilterBuilder) Add(key FilterKey, value string) *FilterBuilder {

	if value == "" {
		return fb
	}

	fb.filters = append(fb.filters, fmt.Sprintf("%s=%s", string(key), value))
	return fb
}

func (fb *FilterBuilder) Query() string {
	if len(fb.filters) == 0 {
		return ""
	}

	raw := strings.Join(fb.filters, " and ")
	return strings.ReplaceAll(url.QueryEscape(raw), "+", "%20")
}

func (fb *FilterBuilder) Clear() *FilterBuilder {
	fb.filters = []string{}
	return fb
}

func (fb *FilterBuilder) AddFiltersToMetaData(metaData map[string]any) map[string]any {
	if len(fb.filters) == 0 {
		return metaData
	}

	metaData[clients.I3dFilters] = fb.Query()
	return metaData
}

func GetFilters(metaData map[string]any) string {
	if filters, ok := metaData[clients.I3dFilters]; ok {
		if filterString, ok := filters.(string); ok {
			return filterString
		}
	}
	return ""
}
