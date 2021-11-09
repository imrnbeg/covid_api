package geocoding

import (
	"encoding/json"
	"log"
	"net/http"
)

// Get Location from lattitude and longitude
func GetLocationStateCode(lat, long string) string {
	url := "http://api.positionstack.com/v1/reverse?access_key=9cbece0bf036ffe410052b831fb1d835&query=" + lat + "%2C" + long
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	var geoCodingResponse Response
	err = json.NewDecoder(resp.Body).Decode(&geoCodingResponse)
	if err != nil {
		log.Println(err)
	}
	data := geoCodingResponse.Data
	if len(data) > 0 && data[0] != nil {
		item1 := data[0]
		stateCode := item1.StateCode
		log.Println("GeoCodingResponse Item =", item1)
		return stateCode
	}

	return DefaultStateCode
}
