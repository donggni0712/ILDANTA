package Domain

type Result struct {
	Where      string       `json:"whereOn"`
	FirstPaths []*FirstPath `json:"firstPath"`
}

type FirstPath struct {
	Name            string            `json:"whatOn"`
	TransferNum     string            `json:"transferNum"`
	TotalTime       string            `json:"totalTime"`
	AfterPathThemes []*AfterPathTheme `json:"subPath"`
}

type AfterPathTheme struct {
	IsFinal          int                `json:"isFinal"`
	Getoff           string             `json:"whereOff"`
	AfterPathParents []*AfterPathParent `json:"subPath"`
}

type AfterPathParent struct {
	Getin           string            `json:"whereOn"`
	AfterPathChilds []*AfterPathChild `json:"subPath"`
}

type AfterPathChild struct {
	IsFinal         int               `json:"isFinal"`
	Getoff          string            `json:"whereOff"`
	NextName        string            `json:"whatOn"`
	AfterPathThemes []*AfterPathTheme `json:"subPath"`
}
