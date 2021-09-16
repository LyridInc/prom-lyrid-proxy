package cmd

import (
	"perseus.proxy/db"
	"perseus.proxy/model"
	"time"
)

func UpdateScrapes(scape model.ScrapesEndpointResult) {
	if db.FindExporter(scape.ExporterID) != nil {
		temp := model.ScrapesEndpointResult{
			ExporterID:     scape.ExporterID,
			ScrapeResult:   scape.ScrapeResult,
			ScrapeTime:     scape.ScrapeTime,
			LastUpdateTime: time.Now().UTC(),
			IsCompress:     scape.IsCompress,
		}

		db.UpdateScrape(temp)
	}
}

func GetLatestScrapeResult(id string) model.ScrapesEndpointResult {
	scrape := db.GetScrape(id)

	if scrape != nil {
		return *scrape
	}

	return model.ScrapesEndpointResult{}
}
