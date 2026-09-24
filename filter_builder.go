package fleetmanager

import (
	"fmt"
	"github.com/i3dnet/nakama-i3d/internal/clients"
	"strconv"
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

	// Numeric IDs remain numeric; names are quoted before HTTP URL encoding.
	numeric := value != ""
	for _, c := range value {
		if c < '0' || c > '9' {
			numeric = false
			break
		}
	}
	if !numeric || strings.HasSuffix(string(key), "Name") {
		value = strconv.Quote(value)
	}
	fb.filters = append(fb.filters, fmt.Sprintf("%s=%s", string(key), value))
	return fb
}

// Query returns a raw expression; the HTTP client encodes it exactly once.
func (fb *FilterBuilder) Query() string {
	if len(fb.filters) == 0 {
		return ""
	}

	return strings.Join(fb.filters, " and ")
}

func (fb *FilterBuilder) Clear() *FilterBuilder {
	fb.filters = []string{}
	return fb
}

func (fb *FilterBuilder) AddFiltersToMetaData(metaData map[string]any) map[string]any {
	if len(fb.filters) == 0 {
		return metaData
	}

	if metaData == nil {
		metaData = make(map[string]any)
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
