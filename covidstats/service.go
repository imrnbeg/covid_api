package covidstats

import (
	"context"
	"covid_api/geocoding"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get location covid stats
func getLocationCovidStats(cache *redis.Client, db *mongo.Client, lat, long string) (*CovidDataResponse, error) {
	// get state code from lat, long
	stateCode := geocoding.GetLocationStateCode(lat, long)

	var stateCodeData *StateData
	// first try to get from cache - data for statecode
	foundState, stateCodeData, stateErr := getFromCache(cache, stateCode)
	if stateErr == nil && foundState {
		return NewCovidDataResponse(stateCodeData), nil
	}

	// if not found in cache, the hit api -> extract data -> save to db -> save to cache
	covidData := fetchData(stateCode)
	stateCodeData = extractData(stateCode, covidData.StateWise)

	saveCovidData(db, covidData)

	cacheErr := setInCache(cache, stateCode, stateCodeData)
	if cacheErr != nil {
		log.Println("Error setting in cache, key", stateCode, ",value", stateCodeData, "err", cacheErr)
	}

	return NewCovidDataResponse(stateCodeData), nil
}

func fetchData(stateCode string) *CovidData {
	log.Println("covidstats - fetchData args ", stateCode)
	resp, err := http.Get("https://api.covid19india.org/data.json")
	if err != nil {
		log.Fatalln(err)
	}

	var covidData CovidData
	err = json.NewDecoder(resp.Body).Decode(&covidData)
	if err != nil {
		log.Println("Error decoding response body err =", err)
	}
	return &covidData
}

func extractData(stateCode string, stateWiseData []*StateData) (stateCodeData *StateData) {
	// log.Println("Response parsed", covidData)
	for _, stateData := range stateWiseData {
		if stateData.StateCode == stateCode {
			stateCodeData = stateData
		}
	}
	return
}

func saveCovidData(db *mongo.Client, covidData *CovidData) {
	// saves complete statewise covid data list to mongodb
	quickstartDatabase := db.Database("covidData")
	covidCasesCollection := quickstartDatabase.Collection("testCollection")
	currentTime := time.Now()
	data := CovidStats{
		Data:      covidData,
		Timestamp: currentTime.Format("2006-01-02 15:04:05"),
	}
	testData, _ := covidCasesCollection.InsertOne(context.Background(), data)
	fmt.Println("Inserted document into testCollection!\n", testData)

}

func setInCache(cache *redis.Client, key string, value *StateData) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	log.Println("setInCache, key=", key, ", value=", value)
	// using simple cache get and set instead of Hset and Hget for simplicity
	// since get, set supports string values , the struct is first marshalled into bytes and a byte string
	// is stored in cache
	_, err = cache.Set(context.Background(), key, string(bytes), time.Minute*30).Result()
	return err
}

func getFromCache(cache *redis.Client, key string) (bool, *StateData, error) {
	strVal, err := cache.Get(context.Background(), key).Result()
	if err != nil {
		return false, nil, err
	}
	// since byte string is stored in cache, after get, first it is unmarshalled back to struct
	var result StateData
	err = json.Unmarshal([]byte(strVal), &result)
	if err != nil {
		return false, nil, err
	}
	log.Println("getInCache, key=", key, ", value=", result)
	return true, &result, nil
}
