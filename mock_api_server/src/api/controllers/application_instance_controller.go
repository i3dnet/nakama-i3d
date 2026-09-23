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
	pages                                 map[string]mockPage
	nextPage                              uint64
}

func NewApplicationInstanceController() *ApplicationInstanceController {
	first := models.GetApplicationInstance()
	second := models.GetApplicationInstance()
	second.Id = "723709572904"
	size, _ := strconv.Atoi(os.Getenv("MOCK_PAGE_SIZE"))
	return &ApplicationInstanceController{instances: map[string]*models.ApplicationInstance{first.Id: first, second.Id: second}, pageSize: size, pages: make(map[string]mockPage)}
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
	terms, err := filterTerms(filter)
	if err != nil {
		return false, err
	}
	for _, term := range terms {
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
	changes, err := metadataChanges(body.Metadata)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
			instance.Metadata = mergeMetadata(instance.Metadata, changes)
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
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.lists++
	now := time.Now()
	for token, page := range gm.pages {
		if !page.expires.After(now) {
			delete(gm.pages, token)
		}
	}
	if gm.pageSize > 0 && limit > gm.pageSize {
		limit = gm.pageSize
	}
	var page mockPage
	if token := c.GetHeader("PAGE-TOKEN"); token != "" {
		var ok bool
		page, ok = gm.pages[token]
		if !ok {
			c.JSON(400, gin.H{"error": "invalid or expired cursor"})
			return
		}
	} else {
		page = mockPage{results: make([]json.RawMessage, 0), expires: now.Add(time.Minute)}
		for _, instance := range gm.sorted() {
			match, err := matches(instance, c.Query("filters"))
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			if match {
				data, err := json.Marshal(instance)
				if err != nil {
					c.JSON(500, gin.H{"error": "cannot snapshot instance"})
					return
				}
				page.results = append(page.results, data)
			}
		}
	}
	end := min(limit, len(page.results))
	if end < len(page.results) {
		gm.nextPage++
		token := fmt.Sprintf("page-%d", gm.nextPage)
		gm.pages[token] = mockPage{results: page.results[end:], expires: page.expires}
		c.Header("PAGE-TOKEN", token)
	}
	c.JSON(200, page.results[:end])
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

type mockPage struct {
	results []json.RawMessage
	expires time.Time
}

func filterTerms(filter string) ([]string, error) {
	var terms []string
	quoted, escaped, start := false, false, 0
	for i := 0; i < len(filter); i++ {
		if escaped {
			escaped = false
			continue
		}
		if quoted && filter[i] == '\\' {
			escaped = true
			continue
		}
		if filter[i] == '"' {
			quoted = !quoted
			continue
		}
		if !quoted && strings.HasPrefix(filter[i:], " and ") {
			terms = append(terms, filter[start:i])
			i += 4
			start = i + 1
		}
	}
	if quoted || escaped {
		return nil, fmt.Errorf("unterminated mock filter value")
	}
	return append(terms, filter[start:]), nil
}

func metadataChanges(changes []models.MetadataChange) (map[string]*string, error) {
	if changes == nil {
		return nil, fmt.Errorf("metadata array is required")
	}
	result := make(map[string]*string, len(changes))
	for _, change := range changes {
		if _, exists := result[change.Key]; change.Key == "" || exists {
			return nil, fmt.Errorf("invalid metadata key")
		}
		var value *string
		if err := json.Unmarshal(change.Value, &value); err != nil {
			return nil, fmt.Errorf("metadata value must be a string or null")
		}
		result[change.Key] = value
	}
	return result, nil
}

func mergeMetadata(existing []models.KeyValue, changes map[string]*string) []models.KeyValue {
	values := make(map[string]string, len(existing)+len(changes))
	for _, pair := range existing {
		values[pair.Key] = pair.Value
	}
	for key, value := range changes {
		if value == nil {
			delete(values, key)
		} else {
			values[key] = *value
		}
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]models.KeyValue, 0, len(keys))
	for _, key := range keys {
		result = append(result, models.KeyValue{Key: key, Value: values[key]})
	}
	return result
}
