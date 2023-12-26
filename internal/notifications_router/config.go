package notifications_router

import (
	"net/url"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const notificationsRouterConfigKey = "notifications_router"

type (
	NotificationsRouterer interface {
		NotificationsRouter() Connector
	}
	notificationsRouterer struct {
		getter kv.Getter
		once   comfig.Once
	}
	config struct {
		RegisterEndpoint *url.URL `fig:"register_endpoint,required"`
		Upstream         *url.URL `fig:"upstream,required"`
		Channel          string   `fig:"channel,required"`
	}
)

func NewNotificationsRouterer(getter kv.Getter) NotificationsRouterer {
	return &notificationsRouterer{
		getter: getter,
	}
}

func (r *notificationsRouterer) NotificationsRouter() Connector {
	return r.once.Do(func() interface{} {
		cfg := new(config)
		err := figure.
			Out(cfg).
			From(kv.MustGetStringMap(r.getter, notificationsRouterConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out notifications router config"))
		}
		return NewConnector(cfg)
	}).(Connector)
}
