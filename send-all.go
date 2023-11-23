package fetchserver

import "github.com/cdvelop/model"

func (h fetchServer) SendAllRequests(endpoint string, data []model.Response, response func(result []model.Response, err string)) {

	response(nil, "error SendAllRequests no implementado en fetchserver")

}
