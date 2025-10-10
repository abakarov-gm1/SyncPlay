package drt

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Register struct {
	Name     string  `json:"name"`
	Photo    *string `json:"photo"`
	Password string  `json:"password"`
	Status   string  `json:"status"`
}

type ResponseLogin struct {
	Text   string `json:"resp"`
	Status bool   `json:"status"`
	Token  string `json:"token"`
}

type ResponseRegister struct {
	Resp string `json:"resp"`
}
