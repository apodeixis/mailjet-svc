package api

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/apodeixis/mailjet-svc/internal/config"
	"github.com/apodeixis/mailjet-svc/internal/service/api/ctx"
	"github.com/apodeixis/mailjet-svc/internal/service/api/handlers"
)

type API interface {
	Start() error
}

type api struct {
	router   chi.Router
	listener net.Listener
	log      *logan.Entry
}

func (a *api) Start() error {
	a.log.Info("Api started on ", a.listener.Addr().String())
	return http.Serve(a.listener, a.router)
}

func NewAPI(cfg config.Config) API {
	return &api{
		router:   newRouter(cfg),
		listener: cfg.Listener(),
		log:      cfg.Log(),
	}
}

func newRouter(cfg config.Config) chi.Router {
	r := chi.NewRouter()
	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleware(
			ctx.SetLog(cfg.Log()),
			ctx.SetMailjetClient(cfg.Mailjet()),
		),
	)
	r.Post("/notifications", handlers.CreateNotification)
	return r
}
