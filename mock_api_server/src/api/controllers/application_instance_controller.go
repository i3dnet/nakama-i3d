package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/mock-api-server/api/models"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ApplicationInstanceController struct {
	allocationDelay                       time.Duration
	responseStatus                        int
	mu                                    sync.Mutex
	instances                             map[string]*models.ApplicationInstance
	allocations, updates, restarts, lists int
	pageSize                              int
}

func NewApplicationInstanceController() *ApplicationInstanceController {
	first := models.GetApplicationInstance()
	second := models.GetApplicationInstance()
	second.Id = "723709572904"
	size, _ := strconv.Atoi(os.Getenv("MOCK_PAGE_SIZE"))
	return &ApplicationInstanceController{instances: map[string]*models.ApplicationInstance{first.Id: first, second.Id: second}, pageSize: size}
}
func decode(c *gin.Context, value any) bool {
	decoder := json.NewDecoder(io.LimitReader(c.Request.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(value); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON body"})
		return false
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		c.JSON(400, gin.H{"error": "trailing JSON"})
		return false
	}
	return true
}
func validMetadata(metadata []models.KeyValue) bool {
	if metadata == nil {
		return false
	}
	seen := map[string]bool{}
	for _, pair := range metadata {
		if pair.Key == "" || seen[pair.Key] {
			return false
		}
		seen[pair.Key] = true
	}
	return true
}

// This mock implements equality joined by "and", the subset used by local tests.
// Reject unsupported syntax instead of silently ignoring a filter.
func matches(instance *models.ApplicationInstance, filter string) (bool, error) {
	fields := map[string]string{"applicationId": instance.ApplicationID, "fleetId": instance.FleetID, "dcLocationId": strconv.Itoa(instance.DCLocationID), "regionId": instance.RegionID, "hostId": strconv.Itoa(instance.HostID), "applicationBuildId": instance.ApplicationBuildID, "deploymentEnvironmentId": instance.DeploymentEnvironmentID, "fleetName": instance.FleetName, "dcLocationName": instance.DCLocationName, "regionName": instance.RegionName, "applicationBuildName": instance.ApplicationBuildName, "deploymentEnvironmentName": instance.DeploymentEnvironmentName}
	matched := true
	if filter == "" {
		return true, nil
	}
	for _, term := range strings.Split(filter, " and ") {
		term = strings.Trim(strings.TrimSpace(term), "()")
		key, value, ok := strings.Cut(term, "=")
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		actual, supported := fields[key]
		if !ok || !supported || value == "" {
			return false, fmt.Errorf("unsupported mock filter")
		}
		if strings.HasPrefix(value, "\"") {
			var err error
			value, err = strconv.Unquote(value)
			if err != nil {
				return false, err
			}
		}
		if actual != value {
			matched = false
		}
	}
	return matched, nil
}
func (gm *ApplicationInstanceController) sorted() []*models.ApplicationInstance {
	keys := make([]string, 0, len(gm.instances))
	for key := range gm.instances {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	instances := make([]*models.ApplicationInstance, 0, len(keys))
	for _, key := range keys {
		instances = append(instances, gm.instances[key])
	}
	return instances
}
func (gm *ApplicationInstanceController) Create(c *gin.Context) {
	var body models.KeyValueMetadata
	if !decode(c, &body) {
		return
	}
	if !validMetadata(body.Metadata) {
		c.JSON(400, gin.H{"error": "metadata array is required"})
		return
	}
	gm.mu.Lock()
	locked := true
	defer func() {
		if locked {
			gm.mu.Unlock()
		}
	}()
	for _, instance := range gm.sorted() {
		match, err := matches(instance, c.Query("filters"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if match && instance.ApplicationID == c.Param("applicationId") && instance.Status == 4 {
			instance.Status = 5
			instance.Metadata = body.Metadata
			gm.allocations++
			response := *instance
			if gm.responseStatus != 0 {
				response.Status = gm.responseStatus
			}
			delay := gm.allocationDelay
			gm.allocationDelay = 0
			gm.responseStatus = 0
			data, _ := json.Marshal([]*models.ApplicationInstance{&response})
			gm.mu.Unlock()
			locked = false
			// Deliberately finish the provider mutation even if the caller times out.
			if delay > 0 {
				time.Sleep(delay)
			}
			c.Data(200, "application/json", data)
			return
		}
	}
	c.JSON(http.StatusConflict, gin.H{"error": "no matching online instance"})
}
func (gm *ApplicationInstanceController) Restart(c *gin.Context) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	instance := gm.instances[c.Param("instanceId")]
	if instance == nil {
		c.JSON(404, gin.H{"error": "unknown instance"})
		return
	}
	instance.Status = 4
	instance.NumPlayers = 0
	instance.Metadata = []models.KeyValue{}
	gm.restarts++
	c.JSON(200, []*models.CommandResult{models.NewCommand()})
}
func (gm *ApplicationInstanceController) Update(c *gin.Context) {
	var body models.ApplicationInstance
	if !decode(c, &body) {
		return
	}
	gm.mu.Lock()
	defer gm.mu.Unlock()
	instance := gm.instances[c.Param("instanceId")]
	if instance == nil {
		c.JSON(404, gin.H{"error": "unknown instance"})
		return
	}
	if body.Id != instance.Id || body.ApplicationID != instance.ApplicationID || body.FleetID != instance.FleetID || !validMetadata(body.Metadata) {
		c.JSON(400, gin.H{"error": "invalid instance update"})
		return
	}
	instance.Metadata = body.Metadata
	gm.updates++
	c.JSON(200, []*models.ApplicationInstance{instance})
}
func (gm *ApplicationInstanceController) Get(c *gin.Context) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	instance := gm.instances[c.Param("instanceId")]
	if instance == nil {
		c.JSON(404, gin.H{"error": "unknown instance"})
		return
	}
	c.JSON(200, []*models.ApplicationInstance{instance})
}
func (gm *ApplicationInstanceController) List(c *gin.Context) {
	limit := 100
	if value := c.GetHeader("RANGED-DATA"); value != "" {
		if !strings.HasPrefix(value, "results=") {
			c.JSON(400, gin.H{"error": "invalid range"})
			return
		}
		var err error
		limit, err = strconv.Atoi(strings.TrimPrefix(value, "results="))
		if err != nil || limit < 1 || limit > 100 {
			c.JSON(400, gin.H{"error": "invalid range"})
			return
		}
	}
	offset := 0
	if value := c.GetHeader("PAGE-TOKEN"); value != "" {
		var err error
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			c.JSON(400, gin.H{"error": "invalid cursor"})
			return
		}
	}
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.lists++
	if gm.pageSize > 0 && limit > gm.pageSize {
		limit = gm.pageSize
	}
	result := make([]*models.ApplicationInstance, 0)
	for _, instance := range gm.sorted() {
		match, err := matches(instance, c.Query("filters"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if match {
			result = append(result, instance)
		}
	}
	if offset > len(result) {
		c.JSON(400, gin.H{"error": "cursor out of range"})
		return
	}
	end := min(offset+limit, len(result))
	if end < len(result) {
		c.Header("PAGE-TOKEN", strconv.Itoa(end))
	}
	c.JSON(200, result[offset:end])
}

// Test controls exist only in this mock service, never in the provider adapter.
func (gm *ApplicationInstanceController) State(c *gin.Context) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	c.JSON(200, gin.H{"allocations": gm.allocations, "updates": gm.updates, "restarts": gm.restarts, "list_pages": gm.lists, "instances": gm.sorted()})
}
func (gm *ApplicationInstanceController) SetStatus(c *gin.Context) {
	var body struct {
		Status *int `json:"status"`
	}
	if !decode(c, &body) {
		return
	}
	if body.Status == nil || *body.Status < 0 || *body.Status > 6 {
		c.JSON(400, gin.H{"error": "invalid status"})
		return
	}
	gm.mu.Lock()
	defer gm.mu.Unlock()
	instance := gm.instances[c.Param("instanceId")]
	if instance == nil {
		c.JSON(404, gin.H{"error": "unknown instance"})
		return
	}
	instance.Status = *body.Status
	c.JSON(200, gin.H{"ok": true})
}

func (gm *ApplicationInstanceController) SetAllocationBehavior(c *gin.Context) {
	var body struct {
		DelayMs        int `json:"delay_ms"`
		ResponseStatus int `json:"response_status"`
	}
	if !decode(c, &body) {
		return
	}
	if body.DelayMs < 0 || body.DelayMs > 2000 || (body.ResponseStatus != 0 && body.ResponseStatus != 4 && body.ResponseStatus != 5) {
		c.JSON(400, gin.H{"error": "invalid test behavior"})
		return
	}
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.allocationDelay = time.Duration(body.DelayMs) * time.Millisecond
	gm.responseStatus = body.ResponseStatus
	c.JSON(200, gin.H{"ok": true})
}
