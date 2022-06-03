package Domain

type FirstPage struct {
	WhereOns []WhereOnComponent `json:"whereOns"`
}

type WhereOnComponent struct {
	WhereOn string            `json:"whereOn"`
	WhatOns []WhatOnComponent `json:"whatOns"`
}

type WhatOnComponent struct {
	WhatOn      string `json:"whatOn"`
	TransferNum string `json:"transferNum"`
	TotalTime   string `json:"totalTime"`
}

type SubPage struct {
	WhatTookOn  string              `json:"whatTookOn"`
	WhereTookOn string              `json:"whereTookOn"`
	WhereOffs   []WhereOffComponent `json:"whereOffs"`
}

type WhereOffComponent struct {
	WhereOff string             `json:"whereOff"`
	WhereOns []WhereOnComponent `json:"whereOns"`
}

//response

type Search struct {
	Sx string `json:"sx"`
	Sy string `json:"sy"`
	Ex string `json:"ex"`
	Ey string `json:"ey"`
}

type SearchSubPath struct {
	Coordinate  Search `json:"coordinate"`
	FirstChoice struct {
		WhereOn string `json:"whereOn"`
		WhatOn  string `json:"whatOn"`
	} `json:"firstChoice"`
	Choices []struct {
		WhereOff string `json:"whereOff"`
		WhereOn  string `json:"whereOn"`
		WhatOn   string `json:"whatOn"`
	} `json:"choices"`
}
