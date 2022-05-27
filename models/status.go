package models

type Status struct {
	Status  bool   `json:"status"`
	Mensaje string `json:"message"`
	Codigo  int32  `json:"code"`
}
