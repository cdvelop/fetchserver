package fetchserver

import (
	"bytes"
	"io"
	"net/http"
)

func (h fetchServer) SendOneRequest(method, endpoint, object string, body_rq any, response func([]map[string]string, error)) {

	var back string
	if object != "" {
		back = "/"
	}

	endpoint = endpoint + back + object

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		response(nil, err)
		return
	}

	var content_type = "application/json"

	var body []byte

	if body_form, ok := body_rq.([]byte); ok {
		body = body_form
		content_type = "multipart/form-data"
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

	// Content-Type ej: multipart/form-data, application/json
	req.Header.Set("Content-Type", content_type)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		response(nil, err)
		return
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		response(nil, err)
		return
	}

	out, err := h.DecodeMaps(resp, object)
	if err != nil {
		response(nil, err)
		return
	}

	response(out, nil)

}
