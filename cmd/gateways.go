package cmd

import (
	"PerseusNocApp.ExporterModuleNocPromGo/db"
	"PerseusNocApp.ExporterModuleNocPromGo/model"
	"context"
)

func AddGateway(gateway model.Gateway) {
	db.InsertGateway(gateway)
}

func ListGateways() []model.Gateway {
	response := make([]model.Gateway, 0)

	cursor := db.GetGatewayCursor()

	if cursor != nil {
		for cursor.Next(context.Background()) {
			var gateway model.Gateway
			if cursor.Decode(&gateway) == nil {
				response = append(response, gateway)
			}
		}
	}

	return response
}

func DeleteGateway(id string) {
	db.DeleteGateway(id)
}
