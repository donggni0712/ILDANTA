package main
 
import (
	"fmt"
	"encoding/json"
    "io/ioutil"
    "net/http"
	"./Domain"
)

func main() {

	var URL string = "https://api.odsay.com/v1/api/searchPubTransPathT?lang=0&SX=127.08186574229312&SY=37.23993898645113&EX=127.05981200975921&EY=37.28556112210226"
	var apikey string = "&apiKey=Mi%2B95EDTMwWb2pbwhatNbhwx4tE4XkBsZ1GiAS2HoGI"
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
	
	
	var paths []Domain.Path
	
	for _,screamPath := range searchPubTransPathT.Result.Path{
		var path Domain.Path
		path.VehiclesType = screamPath.PathType
		path.GetIn = screamPath.Info.FirstStartStation
		for _,tempSubPath := range screamPath.SubPath{
			if tempSubPath.TrafficType == 3{
				continue
			}
			if tempSubPath.TrafficType == 2{
				path.VehicleType = tempSubPath.TrafficType
				path.Name = tempSubPath.Lane[0].BusNo
				path.Getoff = tempSubPath.EndName
				break
			}
			if tempSubPath.TrafficType == 1{
				path.VehicleType = tempSubPath.TrafficType
				path.Name = tempSubPath.Lane[0].Name
				path.Getoff = tempSubPath.EndName
				break
			}
		}
		path.TransferNum = screamPath.Info.BusTransitCount + screamPath.Info.SubwayTransitCount
		path.TotalTime = screamPath.Info.TotalTime
		
		IsNotFirstSubPath := 0
		i:=0
		
		for _,screamSubPath := range screamPath.SubPath{
			var subpath Domain.SubPath
			if screamSubPath.TrafficType == 3{
				continue;
			}
			if screamSubPath.TrafficType == 2{
				if IsNotFirstSubPath==1{
					subpath.Name = screamSubPath.Lane[0].BusNo
					subpath.Gotoff = path.Getoff
					subpath.GetIn = screamSubPath.StartName
					subpath.Getoff = screamSubPath.EndName
					subpath.VehicleType = screamSubPath.TrafficType
					IsNotFirstSubPath++
					path.Subpath = append(path.Subpath,subpath)
					continue
				}
				if IsNotFirstSubPath==0{
					IsNotFirstSubPath ++
					continue
				}
			}
			if screamSubPath.TrafficType == 1{
				if IsNotFirstSubPath==1{
					subpath.Name = screamSubPath.Lane[0].Name
					subpath.Gotoff = path.Getoff
					subpath.GetIn = screamSubPath.StartName
					subpath.Getoff = screamSubPath.EndName
					subpath.VehicleType = screamSubPath.TrafficType
					IsNotFirstSubPath++
					path.Subpath = append(path.Subpath,subpath)
					continue
				}
				if IsNotFirstSubPath==0{
					IsNotFirstSubPath ++
					continue
				}
			subpath.Name = screamSubPath.Lane[0].Name
			subpath.Gotoff = path.Subpath[i].Getoff
			i++
			subpath.GetIn = screamSubPath.StartName
			subpath.Getoff = screamSubPath.EndName
			subpath.VehicleType = screamSubPath.TrafficType
			path.Subpath = append(path.Subpath,subpath)
			}
		}
		paths = append(paths,path)
	}
	//fmt.Println(paths)
}
