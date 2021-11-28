package Repository
 
import (
	"fmt"
	"../ApiClient"
	//"../Domain"
)

func TestApiClient(SX string, SY string, EX string, EY string, apikey string) {
	URL := fmt.Sprintf("https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=%s&SY=%s&EX=%s&EY=%s",SX,SY,EX,EY)
	
	paths := ApiClient.CallRoute(URL,apikey)
	
	for _,i := range paths{
		fmt.Println(i)
	}
}
