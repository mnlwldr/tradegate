package tradegate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TadegateResponse
type TradegateResponse struct {
	Bid     string      `json:"bid"`
	Ask     string      `json:"ask"`
	Bidsize int         `json:"bidsize"`
	Asksize int         `json:"asksize"`
	Delta   string      `json:"delta"`
	Stueck  int         `json:"stueck"`
	Umsatz  int         `json:"umsatz"`
	Avg     float64     `json:"avg"`
	Last    interface{} `json:"last"`
	High    interface{} `json:"high"`
	Low     interface{} `json:"low"`
}

// Get want an "isin" as parameter to get data from Tradegate
// and return a TradegateResponse
func Get(isin string) (TradegateResponse, error) {

	var tradegateReponse TradegateResponse

	response, err := tradegateApiCall(isin)
	if err != nil {
		return tradegateReponse, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tradegateReponse, err
	}

	json.Unmarshal(data, &tradegateReponse)

	return tradegateReponse, nil
}

func tradegateApiCall(isin string) (*http.Response, error) {
	endpoint := "https://www.tradegate.de/refresh.php?isin="
	response, err := http.Get(endpoint + isin)
	if err != nil {
		return nil, err
	}
	return response, nil
}
