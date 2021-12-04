package Domain

type Result struct{
	Where string
	FirstPaths []*FirstPath
}

type FirstPath struct{
	Name string
	TransferNum string
	TotalTime string
	AfterPathThemes []*AfterPathTheme
}

type AfterPathTheme struct{
	Getoff string
	AfterPathParents []*AfterPathParent
}

type AfterPathParent struct{
	Getin string
	AfterPathChilds []*AfterPathChild
}

type AfterPathChild struct{
	IsFinal int
	Getoff string
	NextName string
	AfterPathThemes []*AfterPathTheme
}