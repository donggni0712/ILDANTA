package ApiClient

import (
	"ILDANTA/Domain"
	"ILDANTA/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=127.08186574229312&SY=37.23993898645113&EX=127.05981200975921&EY=37.28556112210226&apiKey=Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI

func CallAPI(SX string, SY string, EX string, EY string, apikey string) []*Domain.Path {

	URL := fmt.Sprintf("https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=%s&SY=%s&EX=%s&EY=%s&apiKey=%s", SX, SY, EX, EY, apikey)
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	var searchPubTransPathT Domain.SearchPubTransPathT
	err = json.Unmarshal(data, &searchPubTransPathT)

	var paths []*Domain.Path
	var ptr *Domain.Path

	for _, streamPath := range searchPubTransPathT.Result.Path {
		var path Domain.Path
		ptr = &path
		IsExist := 0
		var subptr *Domain.SubPath
		path.VehiclesType = streamPath.PathType
		path.GetIn = streamPath.Info.FirstStartStation
		for _, tempSubPath := range streamPath.SubPath {
			if tempSubPath.TrafficType == 3 {
				continue
			}
			if tempSubPath.TrafficType == 2 {
				path.Name = tempSubPath.Lane[0].BusNo
			}
			if tempSubPath.TrafficType == 1 {
				path.Name = tempSubPath.Lane[0].Name
			}
			path.VehicleType = tempSubPath.TrafficType
			path.Getoff = tempSubPath.EndName
			break
		}
		path.MaxTransferNum = streamPath.Info.BusTransitCount + streamPath.Info.SubwayTransitCount
		path.MinTransferNum = path.MaxTransferNum
		path.MaxTotalTime = streamPath.Info.TotalTime
		path.MinTotalTime = streamPath.Info.TotalTime

		for _, SearchSame := range paths {
			if SearchSame.Name == path.Name && SearchSame.GetIn == path.GetIn && SearchSame.Getoff == path.Getoff {
				IsExist = 1
				ptr = SearchSame
				if ptr.MaxTransferNum < path.MaxTransferNum {
					ptr.MaxTransferNum = path.MaxTransferNum
				}
				if ptr.MinTransferNum > path.MinTransferNum {
					ptr.MinTransferNum = path.MinTransferNum
				}
				if ptr.MaxTotalTime < path.MaxTotalTime {
					ptr.MaxTotalTime = path.MaxTotalTime
				}
				if ptr.MinTotalTime > path.MinTotalTime {
					ptr.MinTotalTime = path.MinTotalTime
				}
			}
		}

		IsNotFirstSubPath := 0
		i := 0

		for _, streamSubPath := range streamPath.SubPath {
			var subpath Domain.SubPath
			if streamSubPath.TrafficType == 3 {
				continue
			}
			//
			if IsNotFirstSubPath == 1 {
				if streamSubPath.TrafficType == 2 {
					subpath.Name = streamSubPath.Lane[0].BusNo
				}
				if streamSubPath.TrafficType == 1 {
					subpath.Name = streamSubPath.Lane[0].Name
				}
				subpath.Gotoff = ptr.Getoff
				subpath.GetIn = streamSubPath.StartName
				subpath.Getoff = streamSubPath.EndName
				subpath.VehicleType = streamSubPath.TrafficType
				IsNotFirstSubPath++

				ptr.Next = append(ptr.Next, subpath)
				subptr = &ptr.Next[0]
				continue
			}
			if IsNotFirstSubPath == 0 {
				IsNotFirstSubPath++
				continue
			}
			if streamSubPath.TrafficType == 1 {
				subpath.Name = streamSubPath.Lane[0].Name
			}
			if streamSubPath.TrafficType == 2 {
				subpath.Name = streamSubPath.Lane[0].BusNo
			}
			subpath.Gotoff = ptr.Next[i].Getoff
			i++
			subpath.GetIn = streamSubPath.StartName
			subpath.Getoff = streamSubPath.EndName
			subpath.VehicleType = streamSubPath.TrafficType

			subptr.Next = append(subptr.Next, &subpath)
			subptr = subptr.Next[len(subptr.Next)-1]
		}
		if IsExist == 0 {
			paths = append(paths, &path)
		}
	}
	return paths
}

func CallRoute(SX string, SY string, EX string, EY string, apikey string) []*Domain.Result {
	res := CallAPI(SX, SY, EX, EY, apikey)

	var ResForPrints []*Domain.Result

	for _, path := range res {
		var rfp *Domain.Result
		rfp = &Domain.Result{}
		rfp.Where = path.GetIn
		isExistrfp := 0
		for _, streamResForPrint := range ResForPrints {
			if streamResForPrint.Where == rfp.Where {
				rfp = streamResForPrint
				isExistrfp = 1
				break
			}
		}

		var firstPath *Domain.FirstPath
		firstPath = &Domain.FirstPath{}
		firstPath.Name = path.Name
		firstPath.TransferNum = Utils.GetFromMinMax(path.MinTransferNum, path.MaxTransferNum, "번")
		firstPath.TotalTime = Utils.GetFromMinMax(path.MinTotalTime, path.MaxTotalTime, "분")

		rfp.FirstPaths = append(rfp.FirstPaths, firstPath)

		if isExistrfp == 0 {
			ResForPrints = append(ResForPrints, rfp)
		}

		for _, subpath := range path.Next {
			var afterpathChild *Domain.AfterPathChild
			afterpathChild = &Domain.AfterPathChild{}
			afterpathChild = AppendAfterPath(firstPath, subpath)
			ReculsiveAppend(afterpathChild, subpath)
		}
	}
	return ResForPrints

}

func ReculsiveAppend(TopPath *Domain.AfterPathChild, subpath Domain.SubPath) {
	if len(subpath.Next) == 0 {
		TopPath.IsFinal = 1
		return
	}
	for _, temppath := range subpath.Next {
		var afterpathChild *Domain.AfterPathChild
		afterpathChild = &Domain.AfterPathChild{}
		afterpathChild = AppendAfterPathFromTop(TopPath, *temppath)
		ReculsiveAppend(afterpathChild, *temppath)
	}
	return
}

func AppendAfterPath(firstPath *Domain.FirstPath, subpath Domain.SubPath) *Domain.AfterPathChild {
	var afterpathTheme *Domain.AfterPathTheme
	afterpathTheme = &Domain.AfterPathTheme{}
	isExistAfterPathTheme := 0
	afterpathTheme.Getoff = subpath.Gotoff
	for _, streamAfterPathTheme := range firstPath.AfterPathThemes {
		if streamAfterPathTheme.Getoff == afterpathTheme.Getoff {
			afterpathTheme = streamAfterPathTheme
			isExistAfterPathTheme = 1
			break
		}
	}

	var afterpathParent *Domain.AfterPathParent
	afterpathParent = &Domain.AfterPathParent{}
	afterpathParent.Getin = subpath.GetIn

	var afterpathChild *Domain.AfterPathChild
	afterpathChild = &Domain.AfterPathChild{}
	afterpathChild.NextName = subpath.Name
	afterpathChild.Getoff = subpath.Getoff

	isExistAfterPathParent := 0

	for _, streamAfterPathParent := range afterpathTheme.AfterPathParents {
		if afterpathParent.Getin == streamAfterPathParent.Getin {
			afterpathParent = streamAfterPathParent
			isExistAfterPathParent = 1
			break
		}
	}

	afterpathParent.AfterPathChilds = append(afterpathParent.AfterPathChilds, afterpathChild)

	if isExistAfterPathParent == 0 {
		afterpathTheme.AfterPathParents = append(afterpathTheme.AfterPathParents, afterpathParent)

		if isExistAfterPathTheme == 0 {
			firstPath.AfterPathThemes = append(firstPath.AfterPathThemes, afterpathTheme)
		}
	}
	return afterpathChild
}

func AppendAfterPathFromTop(TopPath *Domain.AfterPathChild, subpath Domain.SubPath) *Domain.AfterPathChild {

	var afterpathTheme *Domain.AfterPathTheme
	afterpathTheme = &Domain.AfterPathTheme{}
	isExistAfterPathTheme := 0
	afterpathTheme.Getoff = subpath.Gotoff
	for _, streamAfterPathTheme := range TopPath.AfterPathThemes {
		if streamAfterPathTheme.Getoff == afterpathTheme.Getoff {
			afterpathTheme = streamAfterPathTheme
			isExistAfterPathTheme = 1
			break
		}
	}

	var afterpathParent *Domain.AfterPathParent
	afterpathParent = &Domain.AfterPathParent{}
	afterpathParent.Getin = subpath.GetIn

	var afterpathChild *Domain.AfterPathChild
	afterpathChild = &Domain.AfterPathChild{}
	afterpathChild.NextName = subpath.Name
	afterpathChild.Getoff = subpath.Getoff

	isExistAfterPathParent := 0

	for _, streamAfterPathParent := range afterpathTheme.AfterPathParents {
		if afterpathParent.Getin == streamAfterPathParent.Getin {
			afterpathParent = streamAfterPathParent
			isExistAfterPathParent = 1
			break
		}
	}

	afterpathParent.AfterPathChilds = append(afterpathParent.AfterPathChilds, afterpathChild)

	if isExistAfterPathParent == 0 {
		afterpathTheme.AfterPathParents = append(afterpathTheme.AfterPathParents, afterpathParent)

		if isExistAfterPathTheme == 0 {
			TopPath.AfterPathThemes = append(TopPath.AfterPathThemes, afterpathTheme)
		}
	}
	return afterpathChild
}
