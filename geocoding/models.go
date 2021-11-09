package geocoding

const DefaultStateCode = "TT"

type Response struct {
	Items []*Item `json:"items"`
}

type Item struct {
	Title   string   `json:title`
	Address *Address `json:address`
}

type Address struct {
	State     string `json:state`
	StateCode string `json:stateCode`
}
