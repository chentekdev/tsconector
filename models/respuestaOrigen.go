package models

type RespuestaOrigen struct {
	Mensaje string           `json:"message"`
	Status  bool             `json:"status"`
	Codigo  int              `json:"code"`
	Data    []ProductoOrigen `json:"data"`
}
