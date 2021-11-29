package Utils

import(
	"fmt"
)

func GetFromMinMax (min int, max int)string{
	if min == max{
		return fmt.Sprintf("%d",min)
	}
	return fmt.Sprintf("%d~%d",min,max)
}