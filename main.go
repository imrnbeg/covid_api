package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	_ "covid_api/docs"

	"covid_api/covidstats"
)

func main() {
	// create cache connection
	cache := redis.NewClient(&redis.Options{
		Addr:     "redis-19966.c291.ap-southeast-2-1.ec2.cloud.redislabs.com:19966",
		Password: "1hLDbJKRW5gpdDOTN5EhNKoTSCxSXaLZ", // no password set
		DB:       0,                                  // use default DB
	})

	// create mongo client connection
	db, err := mongo.NewClient(options.
		Client().
		ApplyURI("mongodb+srv://imran:imran@cluster0.okx0m.mongodb.net/myFirstDatabase"))

	err = db.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(context.Background())
	err = db.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	godotenv.Load()
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	e.GET("/covid/stats/", covidstats.HandleStats(cache, db))
	e.GET("/covid/location/stats", covidstats.HandleLocationStats(cache, db))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(address))
}
