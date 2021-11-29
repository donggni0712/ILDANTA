package Repository
 
import (
	"fmt"
	"../ApiClient"
	"../Domain"
	"../Utils"
)

func TestApiClient(SX string, SY string, EX string, EY string, apikey string) {

	paths := ApiClient.CallRoute(SX,SY,EX,EY,apikey)
	
	for _,i := range paths{
		fmt.Println(i)
	}
}

func ShowFirstRoute(SX string, SY string, EX string, EY string, apikey string) {
	var paths []*Domain.Path
	paths = ApiClient.CallRoute(SX,SY,EX,EY,apikey)
	
	fmt.Println("\t여기서 \t\t 이거타면 \t 환승은 \t 총 시간은")
	fmt.Println("\t===== \t\t======== \t====== \t\t=========")
	for _,path := range paths{
		transferNum := Utils.GetFromMinMax(path.MinTransferNum,path.MaxTransferNum)
		totalTime := Utils.GetFromMinMax(path.MinTotalTime,path.MaxTotalTime)
		fmt.Printf("%10s \t %6s \t %6s \t %6s\n",path.GetIn,path.Name,transferNum,totalTime)
	}
}