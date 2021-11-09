package geocoding

const DefaultStateCode = "TT"

type Response struct {
	Data []*Data `json:"data"`
}

type Data struct {
	State     string `json:"region"`
	StateCode string `json:"region_code"`
}
