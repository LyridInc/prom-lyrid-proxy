package model

import "time"

type Gateway struct {
	ID       string
	Status   string
	Hostname string
}

// exporter type list:
// windows
// node
// mongo
// unknown

type ExporterEndpoint struct {
	ID               string
	Gateway          string
	URL              string
	ExporterType     string
	AdditionalLabels map[string]string `json:"additional_labels"`
}

type ScrapesEndpointResult struct {
	ExporterID     string
	ScrapeResult   string
	IsCompress     bool
	ScrapeTime     time.Time
	LastUpdateTime time.Time
}
