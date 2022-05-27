package models

type Producto struct {
	Marca            string  `json:"marca" bson:"marca"`
	Codigo           string  `json:"codigo" bson:"codigo"`
	HojaDatos        string  `json:"dataSheet" bson:"dataSheet"`
	DescripcionLarga string  `json:"descripcionLarga" bson:"descripcionLarga"`
	Oversize         int32   `json:"oversize" bson:"oversize"`
	Costo            float32 `json:"costo" bson:"costo"`
	PrecioSugerido   float32 `json:"precioSugerido" bson:"precioSugerido"`
	CodigoSat        string  `json:"codigoSat" bson:"codigoSat"`
	SKU              string  `json:"sku" bson:"sku"`
	Status           string  `json:"status" bson:"status"`
	ExistenciaP      int32   `json:"existenciaP" bson:"existenciaP"`
	Existencia       float32 `json:"existencia" bson:"existencia"`
	Imagen           string  `json:"img" bson:"img"`
	Descripcion      string  `json:"descripcion" bson:"descripcion"`
	Precio           float32 `json:"precio" bson:"precio"`
	Proveedor        string  `json:"proveedor" bson:"proveedor"`
	CategoriaOrigen  string  `json:"categoriaOrigen" bson:"categoriaOrigen"`
	IDP              string  `json:"idProveedor" bson:"idProveedor"`
	Categoria        string  `json:"categoria" bson:"categoria"`
	Ancho            float32 `json:"ancho" bson:"ancho"`
	Alto             float32 `json:"alto" bson:"alto"`
	Largo            float32 `json:"largo" bson:"largo"`
	Peso             float32 `json:"peso" bson:"peso"`
	Volumen          float32 `json:"volumen" bson:"volumen"`
}
