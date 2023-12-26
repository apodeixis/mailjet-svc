package service

import (
	"context"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/apodeixis/mailjet-svc/internal/config"
	"github.com/apodeixis/mailjet-svc/internal/service/api"
)

type Service struct {
	cfg config.Config
	api api.API
}

func New(cfg config.Config) *Service {
	return &Service{
		cfg: cfg,
		api: api.NewAPI(cfg),
	}
}

func (s *Service) Run(ctx context.Context) error {
	if err := s.cfg.NotificationsRouter().Register(ctx); err != nil {
		return errors.Wrap(err, "failed to register service in notifications router")
	}
	return s.api.Start()
}
