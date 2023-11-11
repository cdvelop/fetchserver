package fetchserver

import (
	"bytes"
	"io"
	"net/http"
)

// action ej: "create","update","delete","upload" = "POST",  "file" and "read" == GET
// object ej: patientcare.printdoc
func (h *httpServer) SendOneRequest(action, object string, body_rq any, response func([]map[string]string, error)) {
	var method = "GET"
	switch action {
	case "create", "update", "delete", "upload":
		method = "POST"
	}

	endpoint := action + "/" + object

	req, err := http.NewRequest(method, h.server_url+endpoint, nil)
	if err != nil {
		response(nil, err)
		return
	}

	var content_type = "application/json"

	var body []byte

	if body_form, ok := body_rq.([]byte); !ok {
		body = body_form
		content_type = "multipart/form-data"
	} else {

		body, err = h.EncodeMaps(body_rq)
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
