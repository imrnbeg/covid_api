package covidstats

type CovidStats struct {
	Data      *CovidData `bson:"data,omitempty"`
	Timestamp string     `bson:"timestamp,omitempty"`
}
