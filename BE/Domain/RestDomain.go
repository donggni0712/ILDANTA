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
	IsFinal  bool               `json:"isFinal"`
	WhereOns []WhereOnComponent `json:"whereOns"`
}
