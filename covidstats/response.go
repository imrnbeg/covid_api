package covidstats

type CovidDataResponse struct {
	State   *StateData `json:"state"`
	Country *StateData `json:"country"`
}

func NewCovidDataResponse(stateData, countryData *StateData) *CovidDataResponse {
	return &CovidDataResponse{
		State:   stateData,
		Country: countryData,
	}
}

type CovidData struct {
	StateWise []*StateData `json:statewise`
}

type StateData struct {
	Active          string `json:active`
	Confirmed       string `json:confirmed`
	Deaths          string `json:deaths`
	DeltaConfirmed  string `json:deltaconfirmed`
	DeltaDeaths     string `json:deltadeaths`
	DeltaRecovered  string `json:deltarecovered`
	LastUpdatedTime string `json:lastupdatedtime`
	Recovered       string `json:recovered`
	MigratedOther   string `json:migratedother`
	State           string `json:state`
	StateCode       string `json:statecode`
	StateNotes      string `json:statenotes`
}
