package models

type Reponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
}

type NumberModel struct {
	NumberOne int `json:"number_one"`
	NumberTwo int `json:"number_two"`
}
