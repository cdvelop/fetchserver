package fetchserver

import "github.com/cdvelop/model"

func AddFetchAdapter(h *model.Handlers, server_url string) (*httpServer, error) {

	n := httpServer{
		DataConverter: h,
	}

	h.FetchAdapter = &n

	return &n, nil
}
