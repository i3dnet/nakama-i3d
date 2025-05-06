package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/mock-api-server/api/models"
	"io"
	"net/http"
)

type ApplicationInstanceController struct {
	allocated bool
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type MetaData struct {
	MetaData KeyValue `json:"metadata"`
}

func NewApplicationInstanceController() *ApplicationInstanceController {
	return &ApplicationInstanceController{allocated: false}
}

func (gm *ApplicationInstanceController) Create(context *gin.Context) {

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read body"})
		return
	}
	fmt.Println("Received JSON Body:")
	fmt.Println(string(body))
	var appInstances = make([]*models.ApplicationInstance, 0, 1)
	appInstance := models.GetApplicationInstance()
	appInstance.Status = 5
	appInstances = append(appInstances, appInstance)

	gm.allocated = true
	context.JSON(200, appInstances)
}

func (gm *ApplicationInstanceController) Delete(context *gin.Context) {
	gm.allocated = false
	context.JSON(200, models.NewCommand())
}

func (gm *ApplicationInstanceController) Get(context *gin.Context) {
	var appInstances = make([]*models.ApplicationInstance, 0, 1)
	appInstance := models.GetApplicationInstance()
	appInstances = append(appInstances, appInstance)
	context.JSON(200, appInstances)
}

func (gm *ApplicationInstanceController) List(context *gin.Context) {
	var appInstances = make([]*models.ApplicationInstance, 0, 1)
	appInstance := models.GetApplicationInstance()
	appInstances = append(appInstances, appInstance)
	context.JSON(200, appInstances)
}
