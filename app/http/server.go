package http

import (
	"embed"
	"html/template"
	"net/http"
	"ptdog/app/config"

	"github.com/rs/zerolog/log"
)

//go:embed view
var view embed.FS

type Server struct {
	addr string
}

func NewServer() *Server {
	return &Server{
		addr: config.Conf.Http.Addr,
	}
}

func (s *Server) Run() error {
	s.info()
	http.HandleFunc("/", s.view)
	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) view(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFS(view, "view/index.html")
	t.Execute(w, nil)
}

func (s *Server) info() {
	log.Info().Msgf("控制面板: %s (未实现，自行配置config.json)", s.addr)
}
