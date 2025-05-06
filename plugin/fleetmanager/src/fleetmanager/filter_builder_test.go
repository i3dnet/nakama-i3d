package fleetmanager

import (
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/clients"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FilterBuilderTestSuite struct {
	suite.Suite
	builder *FilterBuilder
}

func (suite *FilterBuilderTestSuite) SetupTest() {
	suite.builder = NewFilterBuilder()
}

func (suite *FilterBuilderTestSuite) TestAddAndCreateFilter() {
	// act
	filter := suite.builder.
		Add(FleetId, "fleet-123").
		Add(RegionName, "eu-west").
		Add(DcLocationName, "berlin-dc").
		Query()

	// assert
	expected := "fleetId%3Dfleet-123%20and%20regionName%3Deu-west%20and%20dcLocationName%3Dberlin-dc"
	suite.Equal(expected, filter)
}

func (suite *FilterBuilderTestSuite) TestClear_shouldResultInEmptyQuery() {
	// arrange
	suite.builder.Add(FleetId, "fleet-123")
	suite.builder.Clear()

	// act
	filter := suite.builder.Query()
	expected := ""

	// assert
	suite.Equal(expected, filter)
}

func (suite *FilterBuilderTestSuite) TestAddWithEmptyValue() {
	// arrange
	suite.builder.Add(FleetId, "")

	// act
	filter := suite.builder.Query()

	// assert
	expected := ""
	suite.Equal(expected, filter)
}

func (suite *FilterBuilderTestSuite) TestAddFilterToMetaData_WhenFiltersSet_shouldAddFiltersToMetaDataMap() {
	// arrange
	suite.builder.Add(FleetId, "fleet-123").Add(RegionName, "eu-west")
	metaData := map[string]any{}

	// act
	result := suite.builder.AddFiltersToMetaData(metaData)

	// assert
	_, ok := result[clients.I3dFilters]
	suite.NotEmpty(result)
	suite.True(ok)
	suite.Equal(result[clients.I3dFilters], suite.builder.Query())

}

func (suite *FilterBuilderTestSuite) TestAddFilterToMetaData_WhenNoFilters_shouldReturnMetaDataMap() {
	// arrange
	metaData := map[string]any{}

	// act
	result := suite.builder.AddFiltersToMetaData(metaData)

	// assert
	suite.Equal(metaData, result)
}

func (suite *FilterBuilderTestSuite) TestGetFilters_WhenFilterPropertyIsSet_ShouldReturnFilterString() {
	// arrange
	suite.builder.Add(FleetId, "fleet-123").Add(RegionName, "eu-west")
	metaData := suite.builder.AddFiltersToMetaData(map[string]any{})

	// act
	result := GetFilters(metaData)

	// assert
	suite.Equal(result, suite.builder.Query())
}

func (suite *FilterBuilderTestSuite) TestGetFilters_WhenFilterPropertyIsNotSet_ShouldReturnEmptyString() {
	// arrange
	suite.builder.Add(FleetId, "fleet-123").Add(RegionName, "eu-west")
	metaData := map[string]any{}

	// act
	result := GetFilters(metaData)

	// assert
	suite.Equal("", result)
}

func (suite *FilterBuilderTestSuite) TestFilterKeyConstants() {
	suite.Equal("deploymentEnvironmentId", string(DeploymentEnvironmentId))
	suite.Equal("deploymentEnvironmentName", string(DeploymentEnvironmentName))
	suite.Equal("fleetName", string(FleetName))
	suite.Equal("regionId", string(RegionId))
	suite.Equal("regionName", string(RegionName))
	suite.Equal("dcLocationId", string(DcLocationId))
	suite.Equal("dcLocationName", string(DcLocationName))
	suite.Equal("hostId", string(HostId))
	suite.Equal("applicationBuildId", string(ApplicationBuildId))
	suite.Equal("applicationBuildName", string(ApplicationBuildName))
}

// This function runs the test suite
func TestFilterBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(FilterBuilderTestSuite))
}
