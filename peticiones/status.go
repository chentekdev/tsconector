package peticiones

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/chentekdev/tsconector/models"
	"github.com/chentekdev/tsconector/utils"
)

func Status() (string, bool, error) {
	respuesta := models.Status{}

	cliente := &http.Client{}
	url := "https://api.tecnosinergia.info/v3/status"
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

	if !respuesta.Status {
		return "Token inactivo", false, nil
	}

	return respuesta.Mensaje, true, nil
}
