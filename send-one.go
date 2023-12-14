package fetchserver

import (
	"bytes"
	"io"
	"net/http"

	"github.com/cdvelop/strings"
)

func (h fetchServer) SendOneRequest(method, endpoint, object string, body_rq any, response func(result []map[string]string, err string)) {

	switch method {
	case "GET", "POST":
	default:
		response(nil, "Método "+method+" no soportado")
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

	req, e := http.NewRequest(method, endpoint, nil)
	if e != nil {
		response(nil, e.Error())
		return
	}

	if body_rq != nil {
		var content_type = "application/json"
		var err string
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
			body, err = h.EncodeMaps(body_rq)
			if err != "" {
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
	res, e := client.Do(req)
	if e != nil {
		response(nil, e.Error())
		return
	}
	defer res.Body.Close()

	// fmt.Println("CODE:", res.StatusCode)
	// fmt.Println("ESTATUS 1:", res.Status)

	// OBTENER EL ERROR
	status := res.Header.Get("Status")

	if res.StatusCode != 200 {
		response(nil, status)
		return
	}

	resp, e := io.ReadAll(res.Body)
	if e != nil {
		response(nil, e.Error())
		return
	}

	// fmt.Println("BODY NIL:", resp)
	if file_content {
		// retornamos el fichero
		response([]map[string]string{{"file": string(resp)}}, "")
	} else {

		out, err := h.DecodeMaps(resp)
		if err != "" {
			response(nil, err)
			return
		}

		response(out, "")

	}

}
