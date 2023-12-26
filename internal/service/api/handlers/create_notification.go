package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/mailjet-svc/internal/mailjet"
	"github.com/apodeixis/mailjet-svc/internal/service/api/ctx"
	"github.com/apodeixis/mailjet-svc/internal/service/api/requests"
)

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	request, err := requests.NewCreateNotification(r)
	if err != nil {
		log.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	err = ctx.MailjetClient(r).Send(&mailjet.Message{
		Destination: request.Data.Relationships.Destination.Data.Id,
		Subject:     request.Data.Attributes.Message.Attributes.Subject,
		Text:        request.Data.Attributes.Message.Attributes.Text,
		Html:        request.Data.Attributes.Message.Attributes.Html,
		From:        request.Data.Attributes.Message.Attributes.From,
	})
	if err != nil {
		log.WithError(err).Error("failed to send notification")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
