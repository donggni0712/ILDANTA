package Domain

type SearchPubTransPathT struct {
	Result struct {
		SearchType      int `json:"searchType"`
		OutTrafficCheck int `json:"outTrafficCheck"`
		BusCount        int `json:"busCount"`
		SubwayCount     int `json:"subwayCount"`
		SubwayBusCount  int `json:"subwayBusCount"`
		PointDistance   int `json:"pointDistance"`
		StartRadius     int `json:"startRadius"`
		EndRadius       int `json:"endRadius"`
		Path            []struct {
			PathType int `json:"pathType"`
			Info     struct {
				TrafficDistance    int    `json:"trafficDistance"`
				TotalWalk          int    `json:"totalWalk"`
				TotalTime          int    `json:"totalTime"`
				Payment            int    `json:"payment"`
				BusTransitCount    int    `json:"busTransitCount"`
				SubwayTransitCount int    `json:"subwayTransitCount"`
				MapObj             string `json:"mapObj"`
				FirstStartStation  string `json:"firstStartStation"`
				LastEndStation     string `json:"lastEndStation"`
				TotalStationCount  int    `json:"totalStationCount"`
				BusStationCount    int    `json:"busStationCount"`
				SubwayStationCount int    `json:"subwayStationCount"`
				TotalDistance      int    `json:"totalDistance"`
				TotalWalkTime      int    `json:"totalWalkTime"`
			} `json:"info"`
			SubPath []struct {
				TrafficType  int `json:"trafficType"`
				Distance     int `json:"distance"`
				SectionTime  int `json:"sectionTime"`
				StationCount int `json:"stationCount,omitempty"`
				Lane         []struct {
					BusNo string `json:"busNo"`
					Type  int    `json:"type"`
					BusID int    `json:"busID"`
					Name string  `json:"name"`
					SubwayCode int `json:"subwayCode"`
					SubwayCityCode int `json:"subwayCityCode"`
				} `json:"lane,omitempty"`
				StartName    string  `json:"startName,omitempty"`
				StartX       float64 `json:"startX,omitempty"`
				StartY       float64 `json:"startY,omitempty"`
				EndName      string  `json:"endName,omitempty"`
				EndX         float64 `json:"endX,omitempty"`
				EndY         float64 `json:"endY,omitempty"`
				StartID      int     `json:"startID,omitempty"`
				EndID        int     `json:"endID,omitempty"`
				PassStopList struct {
					Stations []struct {
						Index       int    `json:"index"`
						StationID   int    `json:"stationID"`
						StationName string `json:"stationName"`
						X           string `json:"x"`
						Y           string `json:"y"`
						IsNonStop   string `json:"isNonStop"`
					} `json:"stations"`
				} `json:"passStopList,omitempty"`
				Way         string  `json:"way,omitempty"`
				WayCode     int     `json:"wayCode,omitempty"`
				Door        string  `json:"door,omitempty"`
				StartExitNo string  `json:"startExitNo,omitempty"`
				StartExitX  float64 `json:"startExitX,omitempty"`
				StartExitY  float64 `json:"startExitY,omitempty"`
				EndExitNo   string  `json:"endExitNo,omitempty"`
				EndExitX    float64 `json:"endExitX,omitempty"`
				EndExitY    float64 `json:"endExitY,omitempty"`
			} `json:"subPath"`
		} `json:"path"`
	} `json:"result"`
}

type Path struct{
    Name string
    GetIn string
	Getoff string
    VehicleType int
    VehiclesType int
    MaxTransferNum int
	MinTransferNum int
    Payment int
    MaxTotalTime int
	MinTotalTime int
    Next []SubPath
}

type SubPath struct{
    Name string
    Gotoff string
    GetIn string
    Getoff string
    VehicleType int
	Next []*SubPath
}

type Result struct{
	Where string
	FirstPaths []*FirstPath
}

type FirstPath struct{
	Where string //Not Showed
	Name string
	TransferNum string
	TotalTime string
	AfterPathThemes []*AfterPathTheme
}

type AfterPathTheme struct{
	Where string //Not Showed
	Name string //Not Showed
	Getoff string
	AfterPathParents []*AfterPathParent
}

type AfterPathParent struct{
	Where string //Not Showed
	Name string // Not Showed
	Getoff string //Not Showed
	Getin string
	AfterPathChilds []*AfterPathChild
}

type AfterPathChild struct{
	Getin string //NotShowed
	IsFinal int
	NextName string
	AfterPathThemes []*AfterPathTheme
}