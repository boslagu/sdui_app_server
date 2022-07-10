package models

type FieldProperty struct {
	Fieldname  string      `json:"fieldName"`
	Fieldtype  string      `json:"fieldType"`
	Widgettype string      `json:"widgetType"`
	Isshown    bool        `json:"isShown"`
	Isenabled  bool        `json:"isEnabled"`
	Options    []Option    `json:"options"`
	Design     FieldDesign `json:"design"`
}

type FieldDesign struct {
	Color      string `json:"hexColor"`
	Fontsize   uint   `json:"fontsize"`
	Padding    uint   `json:"padding"`
	Visibility string `json:"visibility"`
}

type Option struct {
	Optiontitle string `json:"optiontitle"`
	Optionvalue uint   `json:"optionlist"`
}
