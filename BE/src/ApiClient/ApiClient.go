package ApiClient
 
import (
	"encoding/json"
    "io/ioutil"
    "net/http"
	"../Domain"
)



func CallRoute(URL string, apikey string) []*Domain.Path{

 	resp, err := http.Get(URL+apikey)
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
	
	for _,streamPath := range searchPubTransPathT.Result.Path{
		var path Domain.Path
		ptr = &path
		IsExist := 0
		var subptr *Domain.SubPath
		path.VehiclesType = streamPath.PathType
		path.GetIn = streamPath.Info.FirstStartStation
		for _,tempSubPath := range streamPath.SubPath{
			if tempSubPath.TrafficType == 3{
				continue
			}
			if tempSubPath.TrafficType == 2{
				path.Name = tempSubPath.Lane[0].BusNo
			}
			if tempSubPath.TrafficType == 1{
				path.Name = tempSubPath.Lane[0].Name
			}
			path.VehicleType = tempSubPath.TrafficType
			path.Getoff = tempSubPath.EndName
			break
		}
		path.TransferNum = streamPath.Info.BusTransitCount + streamPath.Info.SubwayTransitCount
		path.TotalTime = streamPath.Info.TotalTime
		
		for _,SearchSame := range paths{
			if SearchSame.Name == path.Name && SearchSame.GetIn == path.GetIn && SearchSame.Getoff == path.Getoff{
				IsExist=1
				ptr = SearchSame
			}
		}
		
		IsNotFirstSubPath := 0
		i:=0
		
		for _,streamSubPath := range streamPath.SubPath{
			var subpath Domain.SubPath
			if streamSubPath.TrafficType == 3{
				continue;
			}
			//
			if IsNotFirstSubPath==1{
				if streamSubPath.TrafficType == 2{
					subpath.Name = streamSubPath.Lane[0].BusNo
				}
				if streamSubPath.TrafficType == 1{
					subpath.Name = streamSubPath.Lane[0].Name
				}
				subpath.Gotoff = ptr.Getoff
				subpath.GetIn = streamSubPath.StartName
				subpath.Getoff = streamSubPath.EndName
				subpath.VehicleType = streamSubPath.TrafficType
				IsNotFirstSubPath++
					
				ptr.Next = append(ptr.Next,subpath)
				subptr = &ptr.Next[0]
				continue
			}
			if IsNotFirstSubPath==0{
					IsNotFirstSubPath ++
					continue
				}
			/*
			if streamSubPath.TrafficType == 2{
				if IsNotFirstSubPath==1{
					subpath.Name = streamSubPath.Lane[0].BusNo
					subpath.Gotoff = ptr.Getoff
					subpath.GetIn = streamSubPath.StartName
					subpath.Getoff = streamSubPath.EndName
					subpath.VehicleType = streamSubPath.TrafficType
					IsNotFirstSubPath++
					
					ptr.Next = append(ptr.Next,subpath)
					subptr = &ptr.Next[0]
					continue
				}
				if IsNotFirstSubPath==0{
					IsNotFirstSubPath ++
					continue
				}
			subpath.Name = streamSubPath.Lane[0].BusNo
			}
			if streamSubPath.TrafficType == 1{
				if IsNotFirstSubPath==1{
					subpath.Name = streamSubPath.Lane[0].Name
					subpath.Gotoff = ptr.Getoff
					subpath.GetIn = streamSubPath.StartName
					subpath.Getoff = streamSubPath.EndName
					subpath.VehicleType = streamSubPath.TrafficType
					IsNotFirstSubPath++
					
					ptr.Next = append(ptr.Next,subpath)
					subptr = &ptr.Next[0]
					continue
				}
				if IsNotFirstSubPath==0{
					IsNotFirstSubPath ++
					continue
				}
			subpath.Name = streamSubPath.Lane[0].Name
			}
*/
			if  streamSubPath.TrafficType == 1{
				subpath.Name = streamSubPath.Lane[0].Name
			}
			if streamSubPath.TrafficType == 2{
				subpath.Name = streamSubPath.Lane[0].BusNo
			}
		subpath.Gotoff = ptr.Next[i].Getoff
		i++
		subpath.GetIn = streamSubPath.StartName
		subpath.Getoff = streamSubPath.EndName
		subpath.VehicleType = streamSubPath.TrafficType
			
		subptr.Next = append(subptr.Next,&subpath)
		subptr = subptr.Next[len(subptr.Next)-1]
		}
		if IsExist==0{
			paths = append(paths,&path)
		}
	}
	return paths
}
