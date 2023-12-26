package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"

	"github.com/apodeixis/mailjet-svc/internal/mailjet"
	"github.com/apodeixis/mailjet-svc/internal/notifications_router"
)

type Config interface {
	comfig.Logger
	comfig.Listenerer

	notifications_router.NotificationsRouterer
	mailjet.Mailjeter
}

type config struct {
	comfig.Logger
	comfig.Listenerer
	getter kv.Getter

	notifications_router.NotificationsRouterer
	mailjet.Mailjeter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),

		NotificationsRouterer: notifications_router.NewNotificationsRouterer(getter),
		Mailjeter:             mailjet.NewMailjeter(getter),
	}
}
