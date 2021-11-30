package main
 
import (
	"./Repository"
	"./Domain"
	"fmt"
)

func main() {
	SX := "127.08186574229312"
	SY := "37.23993898645113"
	EX := "127.05981200975921"
	EY := "37.28556112210226"
	apikey := "&apiKey=Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI"
	
	var res []*Domain.Result
	
	  res = Repository.ShowFirstRoute(SX,SY,EX,EY,apikey)
	
	var where string
	var what string
	fmt.Println()
	fmt.Scanf("%s",&where)
	fmt.Scanf("%s",&what)
	
	Repository.ClickRoute(where,what,res)
}