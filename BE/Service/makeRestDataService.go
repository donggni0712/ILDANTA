package Service

import (
	"ILDANTA/Domain"
)

func GetFirstRout(results []*Domain.Result) Domain.FirstPage {
	var res Domain.FirstPage
	for _, result := range results {
		var resWhereOn Domain.WhereOnComponent
		resWhereOn.WhereOn = result.Where
		for _, firstPath := range result.FirstPaths {
			var resWhatOn Domain.WhatOnComponent
			resWhatOn.WhatOn = firstPath.Name
			resWhatOn.TotalTime = firstPath.TotalTime
			resWhatOn.TransferNum = firstPath.TransferNum
			resWhereOn.WhatOns = append(resWhereOn.WhatOns, resWhatOn)
		}
		res.WhereOns = append(res.WhereOns, resWhereOn)
	}

	return res
}

func GetFirstSubPath(firstPath Domain.FirstPath, where string) Domain.SubPage {
	var res Domain.SubPage
	res.WhatTookOn = firstPath.Name
	res.WhereTookOn = where
	for _, afterPathTheme := range firstPath.AfterPathThemes {
		var resWhereOff Domain.WhereOffComponent
		resWhereOff.WhereOff = afterPathTheme.Getoff
		resWhereOff.IsFinal = false
		for _, afterPathParent := range afterPathTheme.AfterPathParents {
			var resWhereOn Domain.WhereOnComponent
			resWhereOn.WhereOn = afterPathParent.Getin
			for _, afterPathChild := range afterPathParent.AfterPathChilds {
				var resWhatOn Domain.WhatOnComponent
				resWhatOn.WhatOn = afterPathChild.NextName
				resWhereOn.WhatOns = append(resWhereOn.WhatOns, resWhatOn)
			}
			resWhereOff.WhereOns = append(resWhereOff.WhereOns, resWhereOn)
		}
		res.WhereOffs = append(res.WhereOffs, resWhereOff)
	}
	return res
}

func GetSubPath(subpath Domain.AfterPathChild, where string) Domain.SubPage {
	var res Domain.SubPage
	res.WhereTookOn = where
	res.WhatTookOn = subpath.NextName
	if subpath.IsFinal == 1 {
		var resWhereOff Domain.WhereOffComponent
		resWhereOff.IsFinal = true
		resWhereOff.WhereOff = subpath.Getoff
		res.WhereOffs = append(res.WhereOffs, resWhereOff)
		return res
	}
	for _, afterPathTheme := range subpath.AfterPathThemes {
		var resWhereOff Domain.WhereOffComponent
		resWhereOff.IsFinal = false
		resWhereOff.WhereOff = afterPathTheme.Getoff
		for _, afterPathParent := range afterPathTheme.AfterPathParents {
			var resWhereOn Domain.WhereOnComponent
			resWhereOn.WhereOn = afterPathParent.Getin
			for _, afterPathChild := range afterPathParent.AfterPathChilds {
				var resWhatOn Domain.WhatOnComponent
				resWhatOn.WhatOn = afterPathChild.NextName
				resWhereOn.WhatOns = append(resWhereOn.WhatOns, resWhatOn)
			}
			resWhereOff.WhereOns = append(resWhereOff.WhereOns, resWhereOn)
		}
		res.WhereOffs = append(res.WhereOffs, resWhereOff)
	}
	return res
}
