package Domain

type FirstPage struct {
	WhereOns []WhatOnComponent `json:"whereOns"`
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
	WhereOffs []WhereOffComponent `json:"whereOffs"`
}

type WhereOffComponent struct {
	WhereOff string             `json:"whereOff"`
	WhereOns []WhereOnComponent `json:"whereOns"`
}
