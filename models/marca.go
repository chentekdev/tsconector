package models

type Marca struct {
	Nombre string `json:"nombre" bson:"nombre"`
	Imagen string `json:"img" bson:"img"`
}
