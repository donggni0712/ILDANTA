package Service

import (
	"ILDANTA/ApiClient"
	"ILDANTA/Domain"
	"ILDANTA/Utils"
)

//https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=127.08186574229312&SY=37.23993898645113&EX=127.05981200975921&EY=37.28556112210226&apiKey=Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI

func IsSameExist(paths []*Domain.Path, name, getIn, getOff string) (IsExist int, result *Domain.Path) {
	IsExist = 0
	for _, SearchSame := range paths {
		if SearchSame.Name == name && SearchSame.GetIn == getIn && SearchSame.Getoff == getOff {
			IsExist = 1
			result = SearchSame
		}
	}
	return IsExist, result
}

func IsSameSubPathExist(paths []*Domain.SubPath, name, gotOff, getIn, getOff string) (IsExist int, result *Domain.SubPath) {
	IsExist = 0
	for _, SearchSame := range paths {
		if SearchSame.Name == name && SearchSame.Gotoff == gotOff && SearchSame.GetIn == getIn && SearchSame.Getoff == getOff {
			IsExist = 1
			result = SearchSame
		}
	}
	return IsExist, result
}

func addFirstsubpath(streamSubPath Domain.SubPath_response, ptr *Domain.Path, subptr *Domain.SubPath) (Domain.SubPath, *Domain.SubPath) {
	var subpath Domain.SubPath
	if streamSubPath.TrafficType == 2 {
		for _, subOfsubPath := range streamSubPath.Lane {
			subpath.SetSubpath(subOfsubPath.BusNo, ptr.Getoff, streamSubPath.StartName, streamSubPath.EndName, streamSubPath.TrafficType)
			if IsExist, result := IsSameSubPathExist(ptr.Next, subpath.Name, subpath.Gotoff, subpath.GetIn, subpath.Getoff); IsExist == 1 {
				subptr = result
			}
		}
	}
	if streamSubPath.TrafficType == 1 {
		for _, subOfsubPath := range streamSubPath.Lane {
			subpath.SetSubpath(subOfsubPath.Name, ptr.Getoff, streamSubPath.StartName, streamSubPath.EndName, streamSubPath.TrafficType)
			if IsExist, result := IsSameSubPathExist(ptr.Next, subpath.Name, subpath.Gotoff, subpath.GetIn, subpath.Getoff); IsExist == 1 {
				subptr = result
			}
		}
	}
	return subpath, subptr
}

func addsubpath(streamSubPath Domain.SubPath_response, ptr *Domain.SubPath, subptr *Domain.SubPath) (Domain.SubPath, *Domain.SubPath) {
	var subpath Domain.SubPath
	if streamSubPath.TrafficType == 2 {
		for _, subOfsubPath := range streamSubPath.Lane {
			subpath.SetSubpath(subOfsubPath.BusNo, ptr.Getoff, streamSubPath.StartName, streamSubPath.EndName, streamSubPath.TrafficType)
			if IsExist, result := IsSameSubPathExist(ptr.Next, subpath.Name, subpath.Gotoff, subpath.GetIn, subpath.Getoff); IsExist == 1 {
				subptr = result
			}
		}
	}
	if streamSubPath.TrafficType == 1 {
		for _, subOfsubPath := range streamSubPath.Lane {
			subpath.SetSubpath(subOfsubPath.Name, ptr.Getoff, streamSubPath.StartName, streamSubPath.EndName, streamSubPath.TrafficType)
			if IsExist, result := IsSameSubPathExist(ptr.Next, subpath.Name, subpath.Gotoff, subpath.GetIn, subpath.Getoff); IsExist == 1 {
				subptr = result
			}
		}
	}
	return subpath, subptr
}

func Match(paths []*Domain.Path, path Domain.Path, streamPaths []Domain.SubPath_response) []*Domain.Path {
	var ptr *Domain.Path
	ptr = &path

	var subptr *Domain.SubPath

	IsNotFirstSubPath := 0
	i := 0
	IsExist := 0
	if i, res := IsSameExist(paths, path.Name, path.GetIn, path.Getoff); i == 1 {
		IsExist = 1
		ptr = res
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

	for _, streamSubPath := range streamPaths {
		var subpath Domain.SubPath
		if streamSubPath.TrafficType == 3 {
			continue
		}

		if IsNotFirstSubPath == 1 {
			subpath, subptr = addFirstsubpath(streamSubPath, ptr, subptr)
			IsNotFirstSubPath++

			ptr.Next = append(ptr.Next, &subpath)
			subptr = ptr.Next[0]
			continue
		}
		if IsNotFirstSubPath == 0 {
			IsNotFirstSubPath++
			continue
		}

		subpath, subptr = addsubpath(streamSubPath, subptr, subptr)
		i++

		subptr.Next = append(subptr.Next, &subpath)
		subptr = subptr.Next[len(subptr.Next)-1]
	}
	if IsExist == 0 {
		paths = append(paths, &path)
	}
	return paths

}

func MatchFirstPath(response Domain.SearchPubTransPathT) []*Domain.Path {
	var paths []*Domain.Path

	for _, streamPath := range response.Result.Path {
		var path Domain.Path

		for _, tempSubPath := range streamPath.SubPath {
			if tempSubPath.TrafficType == 3 {
				continue
			}
			if tempSubPath.TrafficType == 2 {
				for _, busLists := range tempSubPath.Lane {
					path.SetPath(busLists.BusNo, streamPath.Info.FirstStartStation, tempSubPath.EndName,
						streamPath.PathType, streamPath.Info.BusTransitCount+streamPath.Info.SubwayTransitCount, path.MaxTransferNum,
						streamPath.Info.TotalTime, streamPath.Info.TotalTime)
					paths = Match(paths, path, streamPath.SubPath)
				}
			}
			if tempSubPath.TrafficType == 1 {
				path.SetPath(tempSubPath.Lane[0].Name, streamPath.Info.FirstStartStation, tempSubPath.EndName,
					streamPath.PathType, streamPath.Info.BusTransitCount+streamPath.Info.SubwayTransitCount, path.MaxTransferNum,
					streamPath.Info.TotalTime, streamPath.Info.TotalTime)

				paths = Match(paths, path, streamPath.SubPath)
			}
			break
		}
	}
	return paths
}

func CallRoute(SX string, SY string, EX string, EY string, apikey string) []*Domain.Result {
	response := ApiClient.CallAPI(SX, SY, EX, EY, apikey)
	res := MatchFirstPath(response)
	var ResForPrints []*Domain.Result

	for _, path := range res {
		var rfp *Domain.Result
		rfp = &Domain.Result{}
		rfp.Where = path.GetIn
		isExistrfp := 0
		isExistrfpnum := 0

		var firstPath *Domain.FirstPath
		firstPath = &Domain.FirstPath{}

		for _, streamResForPrint := range ResForPrints {
			if streamResForPrint.Where == rfp.Where {
				rfp = streamResForPrint
				isExistrfp = 1
				for _, streamFirstPaths := range rfp.FirstPaths {
					if streamFirstPaths.Name == path.Name {
						isExistrfpnum = 1
						firstPath = streamFirstPaths
					}
				}
				break
			}
		}
		if isExistrfpnum == 0 {
			firstPath.Name = path.Name
			firstPath.TransferNum = Utils.GetFromMinMax(path.MinTransferNum, path.MaxTransferNum, "번")
			firstPath.TotalTime = Utils.GetFromMinMax(path.MinTotalTime, path.MaxTotalTime, "분")

			rfp.FirstPaths = append(rfp.FirstPaths, firstPath)
		}
		if isExistrfp == 0 {
			ResForPrints = append(ResForPrints, rfp)
		}

		for _, subpath := range path.Next {
			var afterpathChild *Domain.AfterPathChild
			afterpathChild = &Domain.AfterPathChild{}
			afterpathChild = AppendAfterPath(firstPath, *subpath)
			ReculsiveAppend(afterpathChild, *subpath)
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
