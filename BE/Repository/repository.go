package Repository
 
import (
	"fmt"
	"../ApiClient"
	"../Domain"
)
func ShowFirstRoute(SX string, SY string, EX string, EY string, apikey string) {
	var results []*Domain.Result
	
	results = ApiClient.CallRoute(SX,SY,EX,EY,apikey)
	
	fmt.Printf("\t\t여기서\t\t\t이거타면\t환승은\t\t시간은")
	for _,result := range results{
		fmt.Printf("\n====================================================================\n%15s\t",result.Where)
		for _,firstPath := range result.FirstPaths{
			fmt.Printf("\t%10s %10s %10s\n\t\t\t",firstPath.Name,firstPath.TransferNum,firstPath.TotalTime)
		}
	}
}