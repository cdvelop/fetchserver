package fetchserver

import "github.com/cdvelop/model"

func AddFetchAdapter(h *model.MainHandler) (err string) {
	const e = "error fetchserver nil "

	if h.DataConverter == nil {
		return e + "DataConverter"
	}

	n := fetchServer{
		DataConverter: h,
	}

	h.FetchAdapter = n

	return ""
}
