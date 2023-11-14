package fetchserver

import "github.com/cdvelop/model"

func AddFetchAdapter(h *model.Handlers) error {
	const e = "error fetchserver nil"

	if h.DataConverter == nil {
		return model.Error(e, "DataConverter")
	}

	n := fetchServer{
		DataConverter: h,
	}

	h.FetchAdapter = n

	return nil
}
