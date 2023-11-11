package fetchserver

import (
	"net/url"
)

func AddUrlValuesToEndpoint(endpoint string, params ...map[string]string) string {

	url_values := url.Values{}

	for _, data := range params {
		for key, value := range data {
			// Agregar cada clave-valor al url_values
			url_values.Add(key, value)
		}
	}

	values_url := url_values.Encode()

	if values_url != "" {
		endpoint = endpoint + "?" + values_url
	}

	return endpoint
}
