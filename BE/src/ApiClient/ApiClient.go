package ApiClient
 
import (
	"fmt"
	"encoding/json"
    "io/ioutil"
    "net/http"
	"./Domain"
)

func main() {

	var URL string = "https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=127.08186574229312&SY=37.23993898645113&EX=127.05981200975921&EY=37.28556112210226"
	var apikey string = "&apiKey=Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI"
 	resp, err := http.Get(URL+apikey)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
	
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	var m Domain.SearchPubTransPathT
	err = json.Unmarshal(data, &m)
	fmt.Println(m.Result.Path)
}
