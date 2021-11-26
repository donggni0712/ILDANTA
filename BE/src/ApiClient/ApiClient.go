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
	
	
	var paths []*Domain.Path
	var ptr *Domain.Path
	
	for _,screamPath := range searchPubTransPathT.Result.Path{
		var path Domain.Path

		var subptr *Domain.SubPath
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
		
		for _,SearchSame := range paths{
			if SearchSame.Name == path.Name && SearchSame.GetIn == path.GetIn && SearchSame.Getoff == path.Getoff{
				ptr = &SearchSame
			}
		}
		
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
					
					path.Next = append(path.Next,subpath)
					subptr = &path.Next[0]
					continue
				}
				if IsNotFirstSubPath==0{
					IsNotFirstSubPath ++
					continue
				}
			subpath.Name = screamSubPath.Lane[0].BusNo
			subpath.Gotoff = path.Next[i].Getoff
			i++
			subpath.GetIn = screamSubPath.StartName
			subpath.Getoff = screamSubPath.EndName
			subpath.VehicleType = screamSubPath.TrafficType
				
			subptr.Next = append(subptr.Next,&subpath)
			subptr = subptr.Next[len(subptr.Next)-1]

			}
			if screamSubPath.TrafficType == 1{
				if IsNotFirstSubPath==1{
					subpath.Name = screamSubPath.Lane[0].Name
					subpath.Gotoff = path.Getoff
					subpath.GetIn = screamSubPath.StartName
					subpath.Getoff = screamSubPath.EndName
					subpath.VehicleType = screamSubPath.TrafficType
					IsNotFirstSubPath++
					
					path.Next = append(path.Next,subpath)
					subptr = &path.Next[0]
					continue
				}
				if IsNotFirstSubPath==0{
					IsNotFirstSubPath ++
					continue
				}
			subpath.Name = screamSubPath.Lane[0].Name
			subpath.Gotoff = path.Next[i].Getoff
			i++
			subpath.GetIn = screamSubPath.StartName
			subpath.Getoff = screamSubPath.EndName
			subpath.VehicleType = screamSubPath.TrafficType
				
			subptr.Next = append(subptr.Next,&subpath)
			subptr = subptr.Next[len(subptr.Next)-1]
			}
		}
		paths = append(paths,path)
	}
	for _,i := range paths{
		fmt.Println(i)
	}
}
