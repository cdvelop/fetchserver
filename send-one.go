package fetchserver

import (
	"bytes"
	"io"
	"net/http"

	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func (h fetchServer) SendOneRequest(method, endpoint, object string, body_rq any, response func([]map[string]string, error)) {

	switch method {
	case "GET", "POST":
	default:
		response(nil, model.Error("Método", method, "no soportado"))
		return
	}

	var file_content bool
	if strings.Contains(endpoint, "file?") != 0 || strings.Contains(endpoint, "static/") != 0 {
		file_content = true
	}

	var back string
	if object != "" && !file_content {
		back = "/" + object
	}

	endpoint = endpoint + back

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		response(nil, err)
		return
	}

	if body_rq != nil {
		var content_type = "application/json"

		var body []byte

		// Content-Type multipart/form-data
		if body_form, ok := body_rq.(map[string][]byte); ok {
			if len(body_form) == 1 {
				for boundary, form := range body_form {
					content_type = boundary
					body = form
				}
			}

		} else {
			body, err = h.EncodeMaps(body_rq, object)
			if err != nil {
				response(nil, err)
				return
			}
		}

		if body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		req.Header.Set("Content-Type", content_type)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		response(nil, err)
		return
	}
	defer res.Body.Close()

	// fmt.Println("CODE:", res.StatusCode)
	// fmt.Println("ESTATUS 1:", res.Status)

	// OBTENER EL ERROR
	status := res.Header.Get("Status")

	if res.StatusCode != 200 {
		response(nil, model.Error(status))
		return
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		response(nil, err)
		return
	}

	// fmt.Println("BODY NIL:", resp)
	if file_content {
		// retornamos el fichero
		response([]map[string]string{{"file": string(resp)}}, nil)
	} else {

		out, err := h.DecodeMaps(resp, object)
		if err != nil {
			response(nil, err)
			return
		}

		response(out, nil)

	}

}
