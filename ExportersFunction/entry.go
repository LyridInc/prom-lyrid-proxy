package ExportersFunction

import (
	"PerseusNocApp.ExporterModuleNocPromGo/cmd"
	"PerseusNocApp.ExporterModuleNocPromGo/db"
	"PerseusNocApp.ExporterModuleNocPromGo/model"
	"log"
	"net/http"
	"os"

	persapi "PerseusNocApp.ExporterModuleNocPromGo/api"
	"github.com/gin-gonic/gin"
)

// LyFnInputParams user fills up these parameters
// The struct name need to be static, but the internal composition of the struct can be changed to fit your usage
type LyFnInputParams struct {
	Command string

	Exporter    model.ExporterEndpoint
	Gateway     model.Gateway
	ScapeResult model.ScrapesEndpointResult
	Payload     model.RequestParam
}

// LyFnOutputParams a struct that will be returned
// The struct name need to be static, but the internal composition of the struct can be changed to fit your usage
type LyFnOutputParams struct {
	ReturnPayload interface{}
}

// PreRun (optional) will be executed prior to Run()
func PreRun() {
	if db.Init() {
		log.Println("Mongo URL: ", os.Getenv("MONGO_URL"))
		log.Println("Connected to Mongo URL: ", os.Getenv("MONGO_URL"))
	}
}

// Run will be the
// The function name and definition need to be static, but the in
func Run(input LyFnInputParams) *LyFnOutputParams {

	response := LyFnOutputParams{ReturnPayload: nil}
	switch input.Command {
	case "ListTenants":
		log.Println("Listing Tenant")

	case "AddTenant":
		log.Println("Add Tenant")

	case "DeleteTenant":
		log.Println("Deleting Tenant")

	case "UpdateTenant":
		log.Println("Update Tenant")

	case "ListGateways":
		log.Println("Listing Registered Gateway")
		response.ReturnPayload = cmd.ListGateways()

	case "AddGateway":
		log.Println("Add Gateway")
		cmd.AddGateway(input.Gateway)

	case "DeleteGateway":
		log.Println("Delete Gateway")
		cmd.DeleteGateway(input.Gateway.ID)

	case "UpdateGateway":
		log.Println("Update Gateway")

	case "ListExporter":
		log.Println("Listing Exporter")
		response.ReturnPayload = cmd.ListExporters()

	case "AddExporter":
		log.Println("Adding Exporter")
		log.Println("Parameters: ", input.Exporter)
		cmd.AddExporters(input.Exporter)

	case "UpdateExporter":
		log.Println("Updating Exporter")
		cmd.UpdateExporter(input.Exporter)

	case "DeleteExporter":
		log.Println("Deleting Exporter")
		cmd.DeleteExporter(input.Exporter.ID)

	case "UpdateScrapeResult":
		log.Println("Updating Scrape Result")
		//log.Println("Parameters: ", input.ScapeResult )
		cmd.UpdateScrapes(input.ScapeResult)

	case "GetScrapeResult":
		log.Println("Get Scrape Result")
		//log.Println("Parameters: ", input.ScapeResult )
		response.ReturnPayload = cmd.GetLatestScrapeResult(input.Exporter.ID)

	case "ListPromDiscoService":
	case "AddPromDiscoService":
	case "UpdatePromDiscoService":
	case "DeletePromDiscoService":

	default:
		log.Println("Command not found")
	}

	return &response
}

// PostRun (optional) will be executed after Run() is executed
func PostRun() {

}

func Initialize() *gin.Engine {

	if db.Init() {
		log.Println("Mongo URL: ", os.Getenv("MONGO_URL"))
		log.Println("Connected to Mongo URL: ", os.Getenv("MONGO_URL"))
	}

	r := gin.Default()
	r.LoadHTMLGlob("dist/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	api := r.Group("/api")
	{
		// gateways
		api.GET("/gateways", persapi.ListGateways)
		api.POST("/gateways", persapi.AddGateway)
		api.DELETE("/gateways/:id", persapi.DeleteGateway)

		// exporters
		api.GET("/exporters", persapi.ListExporters)
		api.POST("/exporters", persapi.AddUpdateExporters)
		api.DELETE("/exporters/:id", persapi.DeleteExporters)

		// scrapes
		api.GET("/scrapes/:id", persapi.GetScrapes)
		api.POST("/scrapes", persapi.AddUpdateScrapes)
	}

	return r
}
