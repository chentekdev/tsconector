package peticiones

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/chentekdev/tsconector/models"
	"github.com/chentekdev/tsconector/utils"
)

func GetProductobyID(id string) (models.ProductoOrigen, string, bool, error) {
	type RespuestaOrigenMod struct {
		Status  bool                  `json:"status"`
		Mensaje string                `json:"message"`
		Codigo  int32                 `json:"code"`
		Data    models.ProductoOrigen `json:"data"`
	}
	respuesta := RespuestaOrigenMod{}
	cliente := &http.Client{}
	url := "https://api.tecnosinergia.info/v3/item/list?item_id=" + id
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-token", utils.GetToken())

	if err != nil {
		return respuesta.Data, "Error: al crear request", false, err
	}

	resp, err := cliente.Do(req)

	if err != nil {
		return respuesta.Data, "No se pudo acceder a la url", false, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respuesta.Data, "Error al decodificar los datos", false, err
	}

	json.Unmarshal(bodyBytes, &respuesta)

	if respuesta.Codigo != 0 {
		return respuesta.Data, respuesta.Mensaje, false, nil
	}

	return respuesta.Data, respuesta.Mensaje, true, nil
}
