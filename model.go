package fetchserver

import "github.com/cdvelop/model"

type httpServer struct {
	model.DataConverter
	server_url string
}
