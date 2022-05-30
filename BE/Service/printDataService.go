package Service

import (
	"ILDANTA/Domain"
	"fmt"
)

func ShowFirstRoute(SX string, SY string, EX string, EY string, apikey string) []*Domain.Result {
	var results []*Domain.Result
	results = CallRoute(SX, SY, EX, EY, apikey)

	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\t\t여기서\t\t\t이거타면\t환승은\t\t시간은")
	for _, result := range results {
		fmt.Printf("\n====================================================================\n%15s\t", result.Where)
		for _, firstPath := range result.FirstPaths {
			fmt.Printf("\t%10s %10s %10s\n\t\t\t", firstPath.Name, firstPath.TransferNum, firstPath.TotalTime)
		}
	}

	return results
}

func ShowFirstPath(firstPath *Domain.FirstPath, where string) {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n일단 %15s에서 %10s 탐\n\n", where, firstPath.Name)
	fmt.Printf("\t내려서\t\t\t여기서\t\t\t다시 타")
	for _, afterPathTheme := range firstPath.AfterPathThemes {
		fmt.Printf("\n================================================================\n%10s\t", afterPathTheme.Getoff)
		for _, afterPathParent := range afterPathTheme.AfterPathParents {
			fmt.Printf("\t%10s\t", afterPathParent.Getin)
			for _, afterPathChild := range afterPathParent.AfterPathChilds {
				fmt.Printf("%10s\n\t\t\t\t\t\t\t", afterPathChild.NextName)
			}
		}
	}
	fmt.Println()
}

func ClickRoute(where string, what string, results []*Domain.Result) Domain.FirstPath {

	var res Domain.FirstPath

	for _, result := range results {
		if result.Where == where {
			for _, firstPath := range result.FirstPaths {
				if firstPath.Name == what {
					ShowFirstPath(firstPath, where)
					return *firstPath
				}
			}
		}
	}
	fmt.Println("검색값 없음")
	return res
}

func ShowSubPath(subpath *Domain.AfterPathChild, where string) {
	fmt.Printf("\n\n\n\n\n\n\n\n일단 %15s에서 %10s 탐\n\n", where, subpath.NextName)
	if subpath.IsFinal == 1 {
		fmt.Printf("%10s에서 내려서 걸어가\n", subpath.Getoff)
		return
	}
	fmt.Printf("\t내려서\t\t\t여기서\t\t\t다시 타")
	for _, afterPathTheme := range subpath.AfterPathThemes {
		fmt.Printf("\n==========================================================\n%10s\t", afterPathTheme.Getoff)
		for _, afterPathParent := range afterPathTheme.AfterPathParents {
			fmt.Printf("\t%10s\t", afterPathParent.Getin)
			for _, afterPathChild := range afterPathParent.AfterPathChilds {
				fmt.Printf("%10s\n\t\t\t\t\t\t\t", afterPathChild.NextName)
			}
		}
	}
	fmt.Println()
}

func ClickSubPath(getoff string, getin string, what string, paths []*Domain.AfterPathTheme) Domain.AfterPathChild {

	var res Domain.AfterPathChild

	for _, path := range paths {
		if path.Getoff == getoff {
			for _, afterPathParent := range path.AfterPathParents {
				if afterPathParent.Getin == getin {
					for _, afterPathChild := range afterPathParent.AfterPathChilds {
						if afterPathChild.NextName == what {
							ShowSubPath(afterPathChild, getin)
							return *afterPathChild
						}
					}
				}
			}
		}
	}
	fmt.Println("검색값 없음")
	return res
}
