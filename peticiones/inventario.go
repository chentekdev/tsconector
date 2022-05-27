package peticiones

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chentekdev/tsconector/models"
	"github.com/chentekdev/tsconector/utils"
)

func Inventario() (string, bool, error) {
	//var respuesta map[string]interface{}
	respuesta := models.RespuestaOrigen{}
	inventario := []models.Producto{}
	cliente := &http.Client{}
	url := "https://api.tecnosinergia.info/v3/item/list"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-token", utils.GetToken())

	if err != nil {
		return "Error: al crear request", false, err
	}

	resp, err := cliente.Do(req)

	if err != nil {
		return "No se pudo acceder a la url", false, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error al decodificar los datos", false, err
	}

	json.Unmarshal(bodyBytes, &respuesta)

	fmt.Println("Productos encontrados: ", len(respuesta.Data))

	for i := 0; i < len(respuesta.Data); i++ {
		p := models.Producto{}
		p.Codigo = respuesta.Data[i].Codigo
		p.CodigoSat = respuesta.Data[i].CodigoSat
		p.Costo = respuesta.Data[i].Costo
		p.Descripcion = respuesta.Data[i].Descripcion
		p.DescripcionLarga = respuesta.Data[i].DescripcionLarga
		p.SKU = respuesta.Data[i].SKU
		p.ExistenciaP = respuesta.Data[i].StockTotal
		p.HojaDatos = respuesta.Data[i].HojaDatos
		p.Imagen = respuesta.Data[i].Imagen
		p.Marca = respuesta.Data[i].Marca
		p.Oversize = respuesta.Data[i].SobreDimenciones
		p.PrecioSugerido = respuesta.Data[i].PrecioSugerido
		p.Status = respuesta.Data[i].Status
		p.Proveedor = "TECNOSINERGIA"
		p.CategoriaOrigen = respuesta.Data[i].Categoria
		p.IDP = fmt.Sprint(respuesta.Data[i].ID)
		p.Alto = respuesta.Data[i].Alto
		p.Ancho = respuesta.Data[i].Ancho
		p.Largo = respuesta.Data[i].Largo
		p.Peso = respuesta.Data[i].Peso
		p.Volumen = respuesta.Data[i].Volumen
		inventario = append(inventario, p)
	}
	return fmt.Sprint("Se agregraron ", len(inventario), " productos"), true, nil
}
