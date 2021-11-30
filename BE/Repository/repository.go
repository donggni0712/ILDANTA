package Repository
 
import (
	"fmt"
	"../ApiClient"
	"../Domain"
)

func TestApiClient(SX string, SY string, EX string, EY string, apikey string) {

	paths := ApiClient.CallRoute(SX,SY,EX,EY,apikey)
	
	for _,i := range paths{
		fmt.Println(i)
	}
}

func ShowFirstRoute(SX string, SY string, EX string, EY string, apikey string) {
	var results []*Domain.Result
	
	results = ApiClient.CallRoute(SX,SY,EX,EY,apikey)
	
	fmt.Println(results)
}