package Utils

import (
	"fmt"
	"log"
)

func GetFromMinMax(min int, max int, unit string) string {
	if min == max {
		return fmt.Sprintf("%d%s", min, unit)
	}
	return fmt.Sprintf("%d%s~%d%s", min, unit, max, unit)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
