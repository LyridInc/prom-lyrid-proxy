package cmd

import (
	"context"
	"perseus.proxy/db"
	"perseus.proxy/model"
)

func ListExporters() []model.ExporterEndpoint {
	response := make([]model.ExporterEndpoint, 0)

	cursor := db.GetExportersCursor()

	if cursor != nil {
		for cursor.Next(context.Background()) {
			var exporter model.ExporterEndpoint
			if cursor.Decode(&exporter) == nil {
				response = append(response, exporter)
			}
		}
	}

	return response
}

func AddExporters(exporter model.ExporterEndpoint) {
	db.InsertExporter(exporter)
}

func UpdateExporter(exporter model.ExporterEndpoint) {
	db.UpdateExporter(exporter)
}

func DeleteExporter(id string) {
	db.DeleteExporter(id)
}
