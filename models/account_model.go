package models

import "time"

type LoginModel struct {
	Fields  []FieldProperty `json:"fields"`
	Version string          `json:"version"`
}

type RegistrationModel struct {
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Birtdate    time.Time `json:"birtdate"`
	Mobile      string    `json:"mobile"`
	EmailAdd    string    `json:"emailAdd"`
	Cid         string    `json:"cid"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Imei        string    `json:"imei"`
	DeviceModel string    `json:"deviceModel"`
}
