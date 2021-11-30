package Repository
 
import (
	"fmt"
	"../ApiClient"
	"../Domain"
)
func ShowFirstRoute(SX string, SY string, EX string, EY string, apikey string) []*Domain.Result {
	var results []*Domain.Result
	
	results = ApiClient.CallRoute(SX,SY,EX,EY,apikey)
	
	fmt.Printf("\t\t여기서\t\t\t이거타면\t환승은\t\t시간은")
	for _,result := range results{
		fmt.Printf("\n====================================================================\n%15s\t",result.Where)
		for _,firstPath := range result.FirstPaths{
			fmt.Printf("\t%10s %10s %10s\n\t\t\t",firstPath.Name,firstPath.TransferNum,firstPath.TotalTime)
		}
	}
		
	return results
}

func ShowFirstPath(firstPath *Domain.FirstPath){
	fmt.Printf("일단 %15s에서 %10s 탐\n\n",firstPath.Where,firstPath.Name)
	fmt.Printf("\t내려서\t\t여기서\t\t다시 타")
	for _,afterPathTheme := range firstPath.AfterPathThemes{
		fmt.Printf("\n==========================================================\n%10s\t",afterPathTheme.Getoff)
		for _,afterPath := range afterPathTheme.AfterPaths{
			fmt.Printf("\t%10s\t%10s\n\t\t\t",afterPath.Getin,afterPath.NextName)
		}
	}

}

func ClickRoute(where string, what string,results []*Domain.Result){
	
	for _,result := range results{
		if result.Where == where{
			for _,firstPath := range result.FirstPaths{
				if firstPath.Name == what{
					ShowFirstPath(firstPath)
					return
				}
			}
		}
	}
	fmt.Println("검색값 없음")
}