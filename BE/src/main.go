package main
 
import (
	"fmt"
	"./ApiClient"
)

func main() {

	paths := ApiClient.CallRoute()
	
	for _,i := range paths{
		fmt.Println(i)
	}
}
