package api

import (
	"PerseusNocApp.ExporterModuleNocPromGo/cmd"
	"PerseusNocApp.ExporterModuleNocPromGo/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListGateways(c *gin.Context) {
	c.JSON(200, cmd.ListGateways())
}

func AddGateway(c *gin.Context) {
	var request model.Gateway
	err := c.ShouldBindJSON(&request)
	if err == nil {
		cmd.AddGateway(request)
		c.JSON(200, request)
	} else {
		c.JSON(http.StatusBadRequest, errors.New("input error"))
	}
}

func DeleteGateway(c *gin.Context) {
	id := c.Param("id")
	cmd.DeleteGateway(id)

	c.String(200, "OK")
}
