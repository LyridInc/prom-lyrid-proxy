package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"perseus.proxy/cmd"
	"perseus.proxy/model"
)

func GetScrapes(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, cmd.GetLatestScrapeResult(id))
}

func AddUpdateScrapes(c *gin.Context) {
	var request model.ScrapesEndpointResult
	err := c.ShouldBindJSON(&request)
	if err == nil {
		cmd.UpdateScrapes(request)
		c.JSON(200, "OK")
	} else {
		c.JSON(http.StatusBadRequest, errors.New("input error"))
	}
}
