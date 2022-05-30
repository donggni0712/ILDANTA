package ApiClient

import (
	"ILDANTA/Domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=127.08186574229312&SY=37.23993898645113&EX=127.05981200975921&EY=37.28556112210226&apiKey=Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI

//Require (x,y) of start position, end position and apikey. Return reponse of Calling ODsay API.
func CallAPI(SX string, SY string, EX string, EY string, apikey string) Domain.SearchPubTransPathT {

	URL := fmt.Sprintf("https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=%s&SY=%s&EX=%s&EY=%s&apiKey=%s", SX, SY, EX, EY, apikey)
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	var searchPubTransPathT Domain.SearchPubTransPathT
	err = json.Unmarshal(data, &searchPubTransPathT)

	return searchPubTransPathT
}
