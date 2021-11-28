package main
 
import (
	"fmt"
	"./ApiClient"
)

func main() {
	URL := "https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=127.08186574229312&SY=37.23993898645113&EX=127.05981200975921&EY=37.28556112210226"
	apikey := "&apiKey=Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI"
	
	paths := ApiClient.CallRoute(URL,apikey)
	
	for _,i := range paths{
		fmt.Println(i)
	}
}
