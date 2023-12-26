package mailjet

import (
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/pkg/errors"
)

type (
	Message struct {
		Destination string
		Subject     string
		Text        *string
		Html        *string
		From        *string
	}
	Connector interface {
		Send(msg *Message) error
	}
	connector struct {
		cfg    *config
		client *mailjet.Client
	}
)

func NewConnector(cfg *config) Connector {
	return &connector{
		cfg:    cfg,
		client: mailjet.NewMailjetClient(cfg.PublicAPIKey, cfg.PrivateAPIKey),
	}
}

func (c *connector) Send(msg *Message) error {
	data := c.composeMessagesV31(msg)
	_, err := c.client.SendMailV31(data)
	return errors.Wrap(err, "failed to send message")
}

func (c *connector) composeMessagesV31(msg *Message) *mailjet.MessagesV31 {
	var (
		textPart string
		htmlPart string
		from     = c.cfg.FromEmail
	)

	switch {
	case msg.Text != nil:
		textPart = *msg.Text
	case msg.Html != nil:
		htmlPart = *msg.Html
	case msg.From != nil:
		from = *msg.From
	}

	info := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: from,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: msg.Destination,
				},
			},
			Subject:  msg.Subject,
			TextPart: textPart,
			HTMLPart: htmlPart,
		},
	}

	return &mailjet.MessagesV31{Info: info}
}
