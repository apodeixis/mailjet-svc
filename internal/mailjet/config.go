package mailjet

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const mailjetConfigKey = "mailjet"

type (
	Mailjeter interface {
		Mailjet() Connector
	}
	mailjeter struct {
		getter kv.Getter
		once   comfig.Once
	}
	config struct {
		PublicAPIKey  string `fig:"public_api_key,required"`
		PrivateAPIKey string `fig:"private_api_key,required"`
		FromEmail     string `fig:"from_email,required"`
	}
)

func NewMailjeter(getter kv.Getter) Mailjeter {
	return &mailjeter{
		getter: getter,
	}
}

func (m *mailjeter) Mailjet() Connector {
	return m.once.Do(func() interface{} {
		cfg := new(config)
		err := figure.
			Out(cfg).
			From(kv.MustGetStringMap(m.getter, mailjetConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out mailjet config"))
		}
		return NewConnector(cfg)
	}).(Connector)
}
