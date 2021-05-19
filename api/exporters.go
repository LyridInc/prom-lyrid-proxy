package api

import (
	"PerseusNocApp.ExporterModuleNocPromGo/cmd"
	"PerseusNocApp.ExporterModuleNocPromGo/db"
	"PerseusNocApp.ExporterModuleNocPromGo/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListExporters(c *gin.Context) {
	c.JSON(200, cmd.ListExporters())
}

func AddUpdateExporters(c *gin.Context) {
	var request model.ExporterEndpoint
	err := c.ShouldBindJSON(&request)
	if err == nil {
		found := db.FindExporter(request.ID)
		if found == nil {
			cmd.AddExporters(request)
		} else {
			cmd.UpdateExporter(request)
		}
		c.JSON(200, request)
	} else {
		c.JSON(http.StatusBadRequest, errors.New("input error"))
	}
}

func DeleteExporters(c *gin.Context) {
	id := c.Param("id")
	cmd.DeleteExporter(id)

	c.String(200, "OK")
}
