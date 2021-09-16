package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"perseus.proxy/model"
)

var mongoDb *mongo.Database

func Init() bool {
	if mongoDb != nil {
		return true
	}

	mongo_url := os.Getenv("MONGO_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_url))

	if err != nil {
		log.Fatal(err)
		return false
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Unable to connect to DB: " + err.Error())
		return false
	}

	mongoDb = client.Database(os.Getenv("MONGODB_DBNAME"))

	return true
}

func GetExportersCursor() *mongo.Cursor {
	cur, err := mongoDb.Collection("exporters").Find(context.Background(), bson.D{}, options.Find())
	if err != nil {
		log.Println(err)
		return nil
	}

	return cur
}

func GetGatewayCursor() *mongo.Cursor {
	cur, err := mongoDb.Collection("gateways").Find(context.Background(), bson.D{}, options.Find())
	if err != nil {
		log.Println(err)
		return nil
	}

	return cur
}

func FindExporter(id string) *model.ExporterEndpoint {
	filter := bson.D{{"id", id}}
	res := mongoDb.Collection("exporters").FindOne(context.Background(), filter, options.FindOne())

	if res != nil {
		var exporter model.ExporterEndpoint
		err := res.Decode(&exporter)
		if err == nil {
			return &exporter
		}
	}

	return nil

}

func FindGateway(id string) *model.Gateway {
	filter := bson.D{{"id", id}}
	res := mongoDb.Collection("gateways").FindOne(context.Background(), filter, options.FindOne())

	if res != nil {
		var gateway model.Gateway
		err := res.Decode(&gateway)
		if err == nil {
			return &gateway
		}
	}

	return nil
}

func InsertExporter(exporter model.ExporterEndpoint) {
	if FindExporter(exporter.ID) == nil {
		mongoDb.Collection("exporters").InsertOne(context.Background(), exporter, options.InsertOne())
	}
}

func UpdateExporter(exporter model.ExporterEndpoint) {
	filter := bson.D{{"id", exporter.ID}}
	update := bson.D{{"$set", exporter}}
	upset := true
	_, err := mongoDb.Collection("exporters").UpdateOne(context.Background(), filter, update, &options.UpdateOptions{Upsert: &upset})

	if err != nil {
		log.Println(err)
	}
}

func DeleteExporter(id string) {
	filter := bson.D{{"id", id}}
	_, err := mongoDb.Collection("exporters").DeleteOne(context.Background(), filter, options.Delete())
	if err != nil {
		log.Println(err)
	}
}

func DeleteGateway(id string) {
	filter := bson.D{{"id", id}}
	_, err := mongoDb.Collection("gateways").DeleteOne(context.Background(), filter, options.Delete())
	if err != nil {
		log.Println(err)
	} else {
		filter = bson.D{{"gateway", id}}
		_, err = mongoDb.Collection("exporters").DeleteMany(context.Background(), filter, options.Delete())
		if err != nil {
			log.Println(err)
		}
	}
}

func InsertGateway(gateway model.Gateway) {
	if FindGateway(gateway.ID) == nil {
		mongoDb.Collection("gateways").InsertOne(context.Background(), gateway, options.InsertOne())
	}
}

func UpdateScrape(result model.ScrapesEndpointResult) {
	filter := bson.D{{"exporterid", result.ExporterID}}
	update := bson.D{{"$set", result}}
	upset := true
	_, err := mongoDb.Collection("scrapes").UpdateOne(context.Background(), filter, update, &options.UpdateOptions{Upsert: &upset})

	if err != nil {
		log.Println(err)
	}
}

func GetScrape(id string) *model.ScrapesEndpointResult {
	filter := bson.D{{"exporterid", id}}

	res := mongoDb.Collection("scrapes").FindOne(context.Background(), filter, options.FindOne())

	if res != nil {
		var scrape model.ScrapesEndpointResult
		err := res.Decode(&scrape)
		if err == nil {
			return &scrape
		}
	}
	return nil
}
