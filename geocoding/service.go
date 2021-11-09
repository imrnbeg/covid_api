package geocoding

import (
	"encoding/json"
	"log"
	"net/http"
)

// Get Location from lattitude and longitude
func GetLocationStateCode(lat, long string) string {
	url := "https://revgeocode.search.hereapi.com/v1/revgeocode?at=" + lat + "%2C" + long + "&lang=en-US&apiKey=rV416XQlryCpoZO4T-r-cX0iGszP1gIMar1rk0fKhHw"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	var geoCodingResponse Response
	err = json.NewDecoder(resp.Body).Decode(&geoCodingResponse)
	if err != nil {
		log.Println(err)
	}
	items := geoCodingResponse.Items
	if len(items) > 0 {
		item := items[0]
		stateCode := item.Address.StateCode
		log.Println("GeoCodingResponse Item =", item)
		return stateCode
	}

	return DefaultStateCode
}
