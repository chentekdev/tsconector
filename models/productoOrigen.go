package models

type ProductoOrigen struct {
	Active           bool    `json:"active" bson:"activo"`
	Marca            string  `json:"brand" bson:"marca"`
	Categoria        string  `json:"category" bson:"categoria"`
	Moneda           string  `json:"currency" bson:"moneda"`
	Codigo           string  `json:"code" bson:"codigo"`
	HojaDatos        string  `json:"data_sheet" bson:"dataSheet"`
	DescripcionLarga string  `json:"description" bson:"descripcionLarga"`
	SobreDimenciones int32   `json:"oversize" bson:"oversize"`
	CategoriaPadre   string  `json:"parent_subcategory" bson:"padre"`
	Costo            float32 `json:"regular_price" bson:"costo"`
	CostoDollar      float32 `json:"regular_price_USD" bson:"costoUSD"`
	PrecioSugerido   float32 `json:"sale_price" bson:"precioSugerido"`
	CodigoSat        string  `json:"sat_code" bson:"codigoSat"`
	SKU              string  `json:"sku" bson:"sku"`
	Status           string  `json:"status" bson:"status"`
	StockCDMX        int32   `json:"stock_CDMX" bson:"stockCDMX"`
	StockGDL         int32   `json:"stock_GDL" bson:"stockGDL"`
	StockMER         int32   `json:"stock_MER" bson:"stockMER"`
	StockMTY         int32   `json:"stock_MTY" bson:"stockMTY"`
	StockPUE         int32   `json:"stock_PUE" bson:"stockPUE"`
	StockQRO         int32   `json:"stock_QUE" bson:"stockQRO"`
	StockVER         int32   `json:"stock_VER" bson:"stockVER"`
	StockTotal       int32   `json:"stock_total" bson:"stockTotal"`
	Tipo             int32   `json:"type" bson:"tipo"`
	Volumen          float32 `json:"volume" bson:"volumen"`
	Peso             float32 `json:"weight" bson:"peso"`
	Ancho            float32 `json:"width" bson:"width"`
	Alto             float32 `json:"height" bson:"alto"`
	Imagen           string  `json:"image" bson:"image"`
	ID               int32   `json:"item_id" bson:"id"`
	Largo            float32 `json:"length" bson:"largo"`
	Linea            string  `json:"line" bson:"linea"`
	Descripcion      string  `json:"name" bson:"descripcion"`
}
