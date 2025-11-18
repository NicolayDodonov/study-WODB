package server

// HttpServer - структура описывающая работу сервера
type HttpServer struct {
	Port string
	Addr string
}

func New() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Start() {

}
