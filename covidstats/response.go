package covidstats

type CovidDataResponse struct {
	State *StateData `json:"state"`
}

func NewCovidDataResponse(stateData *StateData) *CovidDataResponse {
	return &CovidDataResponse{
		State: stateData,
	}
}

type CovidData struct {
	StateWise []*StateData `json:statewise`
}

type StateData struct {
	Active          string `json:active`
	Confirmed       string `json:confirmed`
	Deaths          string `json:deaths`
	LastUpdatedTime string `json:lastupdatedtime`
	Recovered       string `json:recovered`
	State           string `json:state`
	StateCode       string `json:statecode`
}
